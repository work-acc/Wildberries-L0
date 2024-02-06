package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"work-acc/wildberries-L0/internal/models"
	"work-acc/wildberries-L0/internal/service"
)

type Order struct {
	serviceOrder *service.Order
}

func (h *Order) Init(router *http.ServeMux) {
	router.HandleFunc("/findOrder", h.findOrder)
}

func (h *Order) findOrder(w http.ResponseWriter, r *http.Request) {
	orderUID := r.URL.Query().Get("order_uid")
	if orderUID == "" {
		responseError := models.ResponseError{
			Status:  http.StatusBadRequest,
			IsOk:    false,
			Message: "order_uid is null",
		}

		data, err := json.Marshal(responseError)
		if err != nil {
			fmt.Printf("%v\n", err)

			return
		}

		w.WriteHeader(responseError.Status)
		w.Write(data)

		return
	}

	responseDTO, err := h.serviceOrder.Find(orderUID)
	if err != nil {
		responseError := models.ResponseError{
			Status:  http.StatusBadRequest,
			IsOk:    true,
			Message: err.Error(),
		}

		data, err := json.Marshal(responseError)
		if err != nil {
			fmt.Printf("%v\n", err)

			return
		}

		w.WriteHeader(responseError.Status)
		w.Write(data)

		return
	}

	responseOk := models.ResponseOk{
		Data:    responseDTO,
		Status:  http.StatusOK,
		IsOk:    true,
		Message: "Success!",
	}

	data, err := json.Marshal(responseOk)
	if err != nil {
		fmt.Printf("%v\n", err)

		return
	}

	w.WriteHeader(responseOk.Status)
	w.Write(data)
}

func NewHandlerOrder(serviceOrder *service.Order) *Order {
	return &Order{
		serviceOrder: serviceOrder,
	}
}
