package main

import "github.com/RedHatInsights/sources-api-go/kafka"

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KafkaMessage struct {
	Topic   string                 `json:"topic"`
	Headers []Header               `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

func toMessage(topic string, m *kafka.Message) *KafkaMessage {
	body := make(map[string]interface{})
	m.ParseTo(&body)

	headers := make([]Header, len(m.Headers))
	for i := range m.Headers {
		headers[i] = Header{
			Key:   m.Headers[i].Key,
			Value: string(m.Headers[i].Value),
		}
	}

	return &KafkaMessage{
		Topic:   topic,
		Headers: headers,
		Body:    body,
	}
}
