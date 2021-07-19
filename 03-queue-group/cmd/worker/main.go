package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	quit    = make(chan struct{})
	sigChan = make(chan os.Signal, 1)
	wg      = &sync.WaitGroup{}
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect")
		os.Exit(1)
	}
	defer nc.Close()

	log.Info().Msg("Connected")

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go worker(nc, "worker-1", "ticket.new")
	go worker(nc, "worker-2", "ticket.new")
	go worker(nc, "worker-3", "ticket.new")

	log.Info().Msg("Launched 3 workers")

	<-sigChan
	log.Info().Msg("Shutting down...")
	close(quit)
	wg.Wait()
	log.Info().Msg("Workers are done")

	nc.Drain()
}

func worker(nc *nats.Conn, name, subject string) {
	logger := log.With().Str("worker", name).Logger()

	ch := make(chan *nats.Msg)
	sub, err := nc.ChanQueueSubscribe(subject, "ticket-workers", ch)
	if err != nil {
		logger.Error().Err(err).Msg("Cannot subscribe to subject")
		return
	}

	wg.Add(1)

	for {
		select {
		case <-quit:
			sub.Unsubscribe()
			wg.Done()
			return
		case msg := <-ch:
			logger.Info().Str("data", string(msg.Data)).Msg("Got msg")
		}
	}
}
