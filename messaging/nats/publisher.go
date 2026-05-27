package messaging

import "github.com/nats-io/nats.go"

func Publish(subject string, payload []byte) {

    nc, _ := nats.Connect(nats.DefaultURL)

    nc.Publish(subject, payload)
}