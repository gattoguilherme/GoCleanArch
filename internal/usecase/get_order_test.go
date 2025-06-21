package usecase_test

import (
	"GoCleanArch/internal/domain/entity"
	"GoCleanArch/internal/infra/database"
	"GoCleanArch/internal/usecase"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)


var _ = Describe("GetOrderByIDUseCase", func() {
	var (
		getOrderByIDUseCase *usecase.GetOrderByIDUseCase
		orderRepoMock       *database.OrderRepositoryMock
	)

	BeforeEach(func() {
		orderRepoMock = database.NewOrderRepositoryMock()
		getOrderByIDUseCase = usecase.NewGetOrderByIDUseCase(orderRepoMock)
	})

	Context("when an order exists", func() {
		It("should return the correct order details", func() {
			existingOrder := &entity.Order{
				Data:      "22/06/2025",
				OrderID:   112233,
				Status:    "Delivered",
				Paid:      true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			orderRepoMock.Save(existingOrder)

			input := usecase.GetOrderByIDInputDTO{OrderID: 112233}
			output, err := getOrderByIDUseCase.Execute(input)

			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeNil())
			Expect(output.OrderID).To(Equal(existingOrder.OrderID))
			Expect(output.Status).To(Equal(existingOrder.Status))
			Expect(output.Paid).To(Equal(existingOrder.Paid))
		})
	})

	Context("when an order does not exist", func() {
		It("should return an error", func() {
			input := usecase.GetOrderByIDInputDTO{OrderID: 999999}
			output, err := getOrderByIDUseCase.Execute(input)

			Expect(err).To(HaveOccurred())
			Expect(output).To(BeNil())
			Expect(err.Error()).To(Equal("order not found"))
		})
	})
})
