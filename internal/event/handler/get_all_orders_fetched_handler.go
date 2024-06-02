package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/songomes/desafiocleanarchitecture/pkg/events"
	"github.com/streadway/amqp"
)

type GetAllOrdersFetchedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewGetAllOrdersFetchedHandler(rabbitMQChannel *amqp.Channel) *GetAllOrdersFetchedHandler {
	return &GetAllOrdersFetchedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *GetAllOrdersFetchedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
