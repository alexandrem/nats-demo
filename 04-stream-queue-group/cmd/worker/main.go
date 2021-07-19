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

const (
	stream       = "ticket-new"
	durableQueue = "ticket-new-worker"
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

	js, err := nc.JetStream()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create jetstream conn")
		return
	}

	// define the stream consumer
	_, err = js.AddConsumer(stream, &nats.ConsumerConfig{
		DeliverPolicy: nats.DeliverAllPolicy,
		Durable:       durableQueue,
		AckPolicy:     nats.AckExplicitPolicy,
		AckWait:       30 * time.Second,
		ReplayPolicy:  nats.ReplayInstantPolicy,
		MaxAckPending: 20000,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create stream consumer")
		return
	}

	go worker(js, "ticket-worker-1")
	go worker(js, "ticket-worker-2")
	go worker(js, "ticket-worker-3")

	log.Info().Msg("Launched 3 workers")

	<-sigChan
	log.Info().Msg("Shutting down...")
	close(quit)
	wg.Wait()
	log.Info().Msg("Workers are done")

	nc.Drain()
}

func worker(js nats.JetStreamContext, name string) {
	logger := log.With().Str("worker", name).Logger()

	sub, err := js.PullSubscribe("", durableQueue,
		nats.Bind(stream, durableQueue))
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
		default:
			msgs, err := sub.Fetch(1)
			if err != nil {
				if err == nats.ErrTimeout {
					continue
				}
				log.Error().Err(err).Msg("unexpected error during msg fetch")
			}
			for _, msg := range msgs {
				logger.Info().Str("data", string(msg.Data)).Msg("got msg")
				msg.Ack()
			}
		}
	}
}
