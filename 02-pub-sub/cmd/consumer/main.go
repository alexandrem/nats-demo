package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

	go consumer(nc, "consumer-1", "ticket.new")
	go consumer(nc, "consumer-2", "ticket.new")
	go consumer(nc, "consumer-3", "ticket.*")

	log.Info().Msg("Launched 3 consumers")

	<-sigChan
	log.Info().Msg("Shutting down...")
	close(quit)
	wg.Wait()
	log.Info().Msg("Consumers are done")

	nc.Drain()
}

func consumer(nc *nats.Conn, name, topic string) {
	logger := log.With().Str("consumer", name).Logger()

	sub, err := nc.Subscribe(topic, func(msg *nats.Msg) {
		logger.Info().Str("topic", topic).Str("data",
			string(msg.Data)).Msg("Got msg")
	})
	if err != nil {
		logger.Error().Err(err).Msg("Cannot subscribe to topic")
		return
	}

	wg.Add(1)

	for {
		select {
		case <-quit:
			wg.Done()
			return
		default:
			msg, err := sub.NextMsg(10 * time.Millisecond)
			if err != nil {
				continue
			}
			log.Info().Msgf("got msg: %v", string(msg.Data))
			nc.Publish(msg.Reply, []byte("hey you"))
		}
	}
}
