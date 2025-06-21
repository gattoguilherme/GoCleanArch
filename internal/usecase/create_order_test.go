package usecase_test

import (
	"GoCleanArch/internal/infra/messaging"
	"GoCleanArch/internal/usecase"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)


var _ = Describe("CreateOrderUseCase", func() {
	var (
		createOrderUseCase *usecase.CreateOrderUseCase
		messageQueueMock   *messaging.OrderMessageQueueMock
	)

	BeforeEach(func() {
		messageQueueMock = messaging.NewOrderMessageQueueMock()
		createOrderUseCase = usecase.NewCreateOrderUseCase(messageQueueMock)
	})

	Context("when creating a new order", func() {
		It("should send the order to the message queue and return the correct DTO", func() {
			input := usecase.CreateOrderInputDTO{
				Data:    "21/06/2025",
				OrderID: 78910,
				Status:  "Processing",
			}

			output, err := createOrderUseCase.Execute(input)

			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeNil())
			Expect(output.Data).To(Equal(input.Data))
			Expect(output.OrderID).To(Equal(input.OrderID))
			Expect(output.Status).To(Equal(input.Status))

			// Here you could also assert that the message was 'sent' by checking logs
			// or by modifying the mock to store the sent message.
		})
	})
})
