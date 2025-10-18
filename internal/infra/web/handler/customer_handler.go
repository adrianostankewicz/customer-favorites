package web

import (
	"encoding/json"
	"net/http"

	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/service"
)

type WebCustomerHandler struct {
	CustomerService *customer.CustomerService
}

func NewWebCustomerHandler(customerService *customer.CustomerService) *WebCustomerHandler {
	return &WebCustomerHandler{
		CustomerService: customerService,
	}
}

func (h *WebCustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto customer.CreateCustomerInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.CustomerService.Create(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
