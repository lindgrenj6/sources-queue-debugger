package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RedHatInsights/sources-api-go/kafka"
	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
)

var (
	cfg         = clowder.LoadedConfig
	eventStream = topic("platform.sources.event-stream")
	status      = topic("platform.sources.status")
)

func main() {
	go listen(eventStream)

	http.Handle("/post-status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	http.ListenAndServe(":8000", nil)
}

func listen(topic string) {
	reader := must(kafka.GetReader(&cfg.Kafka.Brokers[0], "debugger", topic))

	kafka.Consume(reader, func(m kafka.Message) {
		msg := toMessage(topic, &m)
		out, _ := json.Marshal(msg)
		fmt.Printf("%s\n", out)
	})

}

func topic(name string) string {
	for _, t := range cfg.Kafka.Topics {
		if t.RequestedName == name {
			return t.Name
		}
	}

	return name
}

func must[T any](thing T, err error) T {
	if err != nil {
		panic(err)
	}
	return thing
}
