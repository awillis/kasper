package kasper

import "time"

// IncomingMessage describes Kafka incoming message
type IncomingMessage struct {
	Topic     string      // Kafka topic name
	Partition int         // Kafka message xpath: /partition
	Offset    int64       // Kafka message xpath: /offset
	Key       []byte      // Kafka key
	Value     []byte      // Kafka value
	Timestamp time.Time   // Kafka message xpath: /utcGeneratedTime, only set if kafka is version 0.10+
}

// OutgoingMessage describes Kafka outgoing message
type OutgoingMessage struct {
	Topic     string      // Kafka topic name
	Partition int         // Kafka partition
	Key       []byte      // Kafka message key
	Value     []byte      // Kafka message value
}
