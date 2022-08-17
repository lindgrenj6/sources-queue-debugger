package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RedHatInsights/sources-api-go/kafka"
	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
)

var (
	cfg            = clowder.LoadedConfig
	requestedTopic = topic(os.Getenv("TOPIC"))
)

func main() {
	if requestedTopic == "" {
		log.Fatalf("topic %v not found, be sure to set TOPIC in ENV", requestedTopic)
	}

	fmt.Printf("{\"queue_name\": \"%v\"}\n", requestedTopic)

	go listen(requestedTopic)

	http.Handle("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out := getAll(requestedTopic)
		bytes := must(json.MarshalIndent(out, "", "  "))
		w.Write(bytes)
	}))
	http.Handle("/clear", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clearAll(requestedTopic)
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
