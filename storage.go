package main

var (
	db = make(map[string]map[string]int)
)

func getAll(topic string) map[string]int {
	return db[topic]
}

func incrementCount(topic, eventType string) {
	if db[topic] == nil {
		db[topic] = make(map[string]int)
	}

	db[topic][eventType]++
}

func clearAll(topic string) {
	db = make(map[string]map[string]int)
}
