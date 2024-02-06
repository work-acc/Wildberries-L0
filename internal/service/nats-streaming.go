package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/stan.go"

	"work-acc/wildberries-L0/internal/models"
	"work-acc/wildberries-L0/internal/storage/pdb"
)

type Nats struct {
	order *pdb.Order
}

var in_memory = make(map[string]models.OrderDTO)

func (s *Nats) Subscribe(sc stan.Conn) (err error) {
	_, err = sc.Subscribe("test", func(m *stan.Msg) {
		var data models.OrderDTO
		if err := json.Unmarshal(m.Data, &data); err != nil {
			log.Println(err)

			return
		}

		if !s.validate(data) {
			log.Println("the data is not valid")

			return
		}

		jData, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("%v", err)

			return
		}

		if err := s.order.Insert(data.OrderUID, jData); err != nil {
			log.Fatalf("error to insert data to DB: %v", err)
		}

		in_memory[data.OrderUID] = data
	})

	if err != nil {
		log.Fatalf("error for subscribe to channel: %v", err)
	}

	return nil
}

func (s *Nats) validate(data models.OrderDTO) (isOk bool) {
	if data.OrderUID != "" &&
		data.TrackNumber != "" &&
		data.Entry != "" &&
		data.Delivery.Name != "" &&
		data.Delivery.Phone != "" &&
		data.Delivery.Zip != "" &&
		data.Delivery.City != "" &&
		data.Delivery.Address != "" &&
		data.Delivery.Region != "" &&
		data.Delivery.Email != "" &&
		data.Payment.Transaction != "" &&
		data.Payment.Currency != "" &&
		data.Payment.Provider != "" &&
		data.Payment.PaymentDT >= 0 &&
		data.Payment.Bank != "" &&
		len(data.Items) > 0 &&
		data.Locale != "" &&
		data.CustomerID != "" &&
		data.SMID >= 0 &&
		data.DeliveryService != "" &&
		data.ShardKey != "" &&
		data.DateCreated != "" &&
		data.OOFShard != "" {

		for _, val := range data.Items {
			if val.TrackNumber != "" &&
				val.ChrtID >= 0 &&
				val.Price >= 0 &&
				val.TotalPrice >= 0 &&
				val.NmID >= 0 &&
				val.Status >= 0 &&
				val.RID != "" &&
				val.Name != "" &&
				val.Size != "" &&
				val.Brand != "" {
			} else {
				return
			}
		}

		return true
	}

	return
}

func NewServiceNatsStreaming(order *pdb.Order) *Nats {
	return &Nats{
		order: order,
	}
}
