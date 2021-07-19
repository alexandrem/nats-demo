package main

import (
	"fmt"
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

	js, err := nc.JetStream()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create jetstream ctx")
		return
	}

	log.Info().Msg("Connected")

	for i := 1; i <= 10; i++ {
		log.Info().Msg("Sending msg...")
		data := fmt.Sprintf("id [%d]", i)
		if _, err := js.Publish("ticket.new",
			[]byte(data)); err != nil {
			log.Error().Err(err).Msg("Couldn't publish")
		}
		time.Sleep(2 * time.Second)
	}

	log.Info().Msg("Ticket batch sent")
}
