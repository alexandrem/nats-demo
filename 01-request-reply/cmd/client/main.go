package main

import (
	"os"
	"time"

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
	log.Info().Msg("Sending request...")

	msg, err := nc.Request("service-a", []byte("hello"),
		10*time.Millisecond)
	if err != nil {
		log.Error().Err(err).Msg("Didn't receive reply")
		os.Exit(1)
	}

	log.Info().Str("data", string(msg.Data)).Msg("Got reply")
}
