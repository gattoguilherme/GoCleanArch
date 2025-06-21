package handler

import (
	"GoCleanArch/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// OrderHandler handles HTTP requests for orders.
type OrderHandler struct {
	CreateOrderUseCase    *usecase.CreateOrderUseCase
	GetOrderUseCase       *usecase.GetOrderByIDUseCase
	GetAllOrdersUseCase   *usecase.GetAllOrdersUseCase
}

// NewOrderHandler creates a new OrderHandler.
func NewOrderHandler(createOrderUseCase *usecase.CreateOrderUseCase, getOrderUseCase *usecase.GetOrderByIDUseCase, getAllOrdersUseCase *usecase.GetAllOrdersUseCase) *OrderHandler {
	return &OrderHandler{
		CreateOrderUseCase:  createOrderUseCase,
		GetOrderUseCase:     getOrderUseCase,
		GetAllOrdersUseCase: getAllOrdersUseCase,
	}
}

// CreateOrder handles the creation of a new order.
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(input)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
	log.Printf("Order created successfully: %d", input.OrderID)
}

// GetOrder handles the retrieval of an order by its ID.
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request to get order by ID")
	orderIDStr := chi.URLParam(r, "orderId")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		log.Printf("Invalid Order ID: %s", orderIDStr)
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	input := usecase.GetOrderByIDInputDTO{OrderID: orderID}
	output, err := h.GetOrderUseCase.Execute(input)
	if err != nil {
		log.Printf("Error getting order %d: %v", orderID, err)
		// In a real application, you would check for a 'not found' error and return 404
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
	log.Printf("Order %d retrieved successfully", orderID)
}

// GetAllOrders handles the retrieval of all orders.
// @Summary Get all orders
// @Description Retrieve all orders
// @Tags orders
// @Produce json
// @Success 200 {array} entity.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request to get all orders")
	orders, err := h.GetAllOrdersUseCase.Execute()
	if err != nil {
		log.Printf("Error getting all orders: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
	log.Printf("All orders retrieved successfully")
}
