package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type eventListener struct {
	Broker string
	Topic string
	CancelContext context.Context
}

func NewEventListener(broker string, topic string, canelContext context.Context) *eventListener  {
	return &eventListener{
		Broker: broker,
		Topic: broker,
		CancelContext: canelContext,
	}
}

func (eventListener *eventListener) StartAndListenAndPushToChannel(groupId string, event chan string)  {
	readerConfig := kafka.ReaderConfig {
		Brokers: []string{eventListener.Broker},
		Topic: eventListener.Topic,
		GroupID: groupId,
	}

	reader := kafka.NewReader(readerConfig)

	for{
		select{
		case <-eventListener.CancelContext.Done():
			reader.Close()
			event <- "Done"
		default:
			msg, err := reader.ReadMessage(eventListener.CancelContext)
			if err != nil {
				log.Println("cannot read message " + err.Error())
			}
			event <- fmt.Sprintf("Partirion: %d, Offset: %d, Message: %s", msg.Partition, msg.Offset, string (msg.Value))
		}
	}
}