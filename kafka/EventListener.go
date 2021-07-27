package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type eventListener struct {
	Broker        string
	Topic         string
	CancelContext context.Context
}

func NewEventListener(broker string, topic string, cancelContext context.Context) *eventListener {
	return &eventListener{
		Broker:        broker,
		Topic:         topic,
		CancelContext: cancelContext,
	}
}

func (eventListener *eventListener) StartAndListenAndPushToChannel(groupId string, event chan string) {
	readerConfig := kafka.ReaderConfig{
		Brokers: []string{eventListener.Broker},
		Topic:   eventListener.Topic,
		GroupID: groupId,
	}

	reader := kafka.NewReader(readerConfig)

	for {
		select {
		case <-eventListener.CancelContext.Done():
			reader.Close()
			event <- "Done"
		default:
			msg, err := reader.ReadMessage(eventListener.CancelContext)
			if err != nil {
				log.Println("cannot read message " + err.Error())
			}
			event <- fmt.Sprintf("Partition: %d, Offset: %d, Message: %s, Topic: %s", msg.Partition, msg.Offset, string(msg.Value), string(msg.Topic))
		}
	}
}
