package messaging

import (
	"GoCleanArch/internal/domain/entity"
	"encoding/json"
	"log"
)

// OrderMessageQueueMock is a mock implementation of the OrderMessageQueue interface.
type OrderMessageQueueMock struct{}

// NewOrderMessageQueueMock creates a new OrderMessageQueueMock.
func NewOrderMessageQueueMock() *OrderMessageQueueMock {
	return &OrderMessageQueueMock{}
}

// Send simulates sending an order message to a queue.
func (m *OrderMessageQueueMock) Send(order *entity.Order) error {
	jsonData, err := json.Marshal(order)
	if err != nil {
		log.Printf("Error marshalling order for message queue: %v", err)
		return err
	}
	log.Printf("Simulating sending message to SQS: %s", string(jsonData))
	return nil
}
