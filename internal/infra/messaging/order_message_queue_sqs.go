package messaging

import (
	"GoCleanArch/internal/domain/entity"
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// OrderMessageQueueSQS implements the OrderMessageQueue interface for AWS SQS.
type OrderMessageQueueSQS struct {
	Client   *sqs.Client
	QueueURL string
}

// NewOrderMessageQueueSQS creates a new SQS message queue.
func NewOrderMessageQueueSQS(client *sqs.Client, queueURL string) *OrderMessageQueueSQS {
	return &OrderMessageQueueSQS{Client: client, QueueURL: queueURL}
}

// Send sends an order message to the SQS queue.
func (q *OrderMessageQueueSQS) Send(order *entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, err = q.Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &q.QueueURL,
		MessageBody: aws.String(string(body)),
	})

	return err
}
