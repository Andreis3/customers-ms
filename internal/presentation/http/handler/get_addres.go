package handler

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
)

type GetAddressHandler struct {
}

func NewGetAddressHandler() GetAddressHandler {
	return GetAddressHandler{}
}

func (h *GetAddressHandler) Handle(w http.ResponseWriter, r *http.Request) {

	helpers.ResponseSuccess[any](w, http.StatusOK, nil)
}
