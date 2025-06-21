package handler_test

import (
	"GoCleanArch/internal/domain/entity"
	"GoCleanArch/internal/infra/database"
	"GoCleanArch/internal/infra/handler"
	"GoCleanArch/internal/infra/messaging"
	"GoCleanArch/internal/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOrderHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OrderHandler Suite")
}

var _ = Describe("OrderHandler", func() {
	var (
		orderHandler *handler.OrderHandler
		router       *chi.Mux
	)

	BeforeEach(func() {
		orderRepo := database.NewOrderRepositoryMock()
		messageQueue := messaging.NewOrderMessageQueueMock()

		createOrderUseCase := usecase.NewCreateOrderUseCase(messageQueue)
		getOrderUseCase := usecase.NewGetOrderByIDUseCase(orderRepo)

		// Pre-populate data for GET tests
		prePopulatedOrder := &entity.Order{OrderID: 123, Status: "Complete", Paid: true}
		orderRepo.Save(prePopulatedOrder)

		getAllOrdersUseCase := usecase.NewGetAllOrdersUseCase(orderRepo)
		orderHandler = handler.NewOrderHandler(createOrderUseCase, getOrderUseCase, getAllOrdersUseCase)

		router = chi.NewRouter()
		router.Post("/orders", orderHandler.CreateOrder)
		router.Get("/orders/{orderId}", orderHandler.GetOrder)
	})

	Describe("GET /orders", func() {
		Context("when there are orders", func() {
			It("should return 200 OK and the list of orders", func() {
				req := httptest.NewRequest("GET", "/orders", nil)
				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				var response []entity.Order
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				Expect(err).To(BeNil())
				Expect(len(response)).To(BeNumerically(">=", 1))
				Expect(response[0].OrderID).To(Equal(123))
			})
		})
	})

	Describe("POST /orders", func() {
		Context("with a valid request body", func() {
			It("should return 201 Created and the created order", func() {
				orderData := map[string]interface{}{"Data": "23/06/2025", "OrderId": 456, "Status": "New"}
				body, _ := json.Marshal(orderData)

				req := httptest.NewRequest("POST", "/orders", bytes.NewBuffer(body))
				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusCreated))

				var response usecase.CreateOrderOutputDTO
				json.Unmarshal(rr.Body.Bytes(), &response)
				Expect(response.OrderID).To(Equal(456))
			})
		})

		Context("with an invalid request body", func() {
			It("should return 400 Bad Request", func() {
				req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString("invalid json"))
				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})

	Describe("GET /orders/{orderId}", func() {
		Context("when the order exists", func() {
			It("should return 200 OK and the order details", func() {
				req := httptest.NewRequest("GET", "/orders/123", nil)
				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))

				var response usecase.GetOrderByIDOutputDTO
				json.Unmarshal(rr.Body.Bytes(), &response)
				Expect(response.OrderID).To(Equal(123))
				Expect(response.Status).To(Equal("Complete"))
				Expect(response.Paid).To(BeTrue())
			})
		})

		Context("when the order does not exist", func() {
			It("should return 500 Internal Server Error", func() {
				req := httptest.NewRequest("GET", "/orders/999", nil)
				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
