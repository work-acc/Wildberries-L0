package service

import (
	"fmt"
	"reflect"

	"work-acc/wildberries-L0/internal/models"
	"work-acc/wildberries-L0/internal/storage/pdb"
)

type Order struct {
	storageOrder *pdb.Order
}

func (s *Order) Find(order_uid string) (responseDTO models.OrderDTO, err error) {
	responseDTO = in_memory[order_uid]
	if reflect.DeepEqual(responseDTO, models.OrderDTO{}) {
		return responseDTO, fmt.Errorf("order not found with id = %v", order_uid)
	}

	return responseDTO, nil
}

func (s *Order) RecoveryInMemory() error {
	ordersDTO, err := s.storageOrder.FindAll()
	if err != nil {
		return err
	}

	in_memory = ordersDTO

	return nil
}

func NewServiceOrder(storageOrder *pdb.Order) *Order {
	return &Order{
		storageOrder: storageOrder,
	}
}
