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

	end := time.Since(start)
	handler.prometheus.ObserveRequestDuration("customers", "http", http.StatusOK, float64(end.Milliseconds()))

	data := struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}

	helpers.ResponseSuccess[any](w, http.StatusOK, data)
}
