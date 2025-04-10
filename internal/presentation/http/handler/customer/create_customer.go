package customer

import (
	"net/http"
	"time"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
)

type CreateCustomerHandler struct {
	log        interfaces.Logger
	prometheus interfaces.Prometheus
}

func NewCreateCustomerHandler(
	log interfaces.Logger,
	prometheus interfaces.Prometheus,
) CreateCustomerHandler {
	return CreateCustomerHandler{
		log:        log,
		prometheus: prometheus,
	}
}

func (handler *CreateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	data := struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}

	end := time.Since(start)
	handler.prometheus.ObserveRequestDuration("/customers", "http", http.StatusCreated, float64(end.Milliseconds()))
	helpers.ResponseSuccess[any](w, http.StatusCreated, data)
}
