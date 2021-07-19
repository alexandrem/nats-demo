package main

import (
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	quit = make(chan struct{})
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
	service(nc)

	nc.Drain()
}

func service(nc *nats.Conn) {
	sub, err := nc.SubscribeSync("service-a")
	if err != nil {
		log.Error().Err(err).Msg("Cannot subscribe to topic")
		return
	}

	for {
		select {
		case <-quit:
			return
		default:
			msg, err := sub.NextMsg(10 * time.Millisecond)
			if err != nil {
				continue
			}
			log.Info().Str("data", string(msg.Data)).Msg("Got msg")
			// we can also use msg.Respond([]byte("hey you"))
			nc.Publish(msg.Reply, []byte("hey you"))
		}
	}
}
