package main

import (
	"os"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	if err := nc.Publish("ticket.new", []byte("id: 1")); err != nil {
		log.Error().Err(err).Msg("Couldn't publish")
	}

	if err := nc.Publish("ticket.update", []byte("id: 1, "+
		"new value")); err != nil {
		log.Error().Err(err).Msg("Couldn't publish")
	}

	log.Info().Msg("Ticket event sent")
}
