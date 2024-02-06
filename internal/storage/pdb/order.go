package pdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"work-acc/wildberries-L0/internal/models"
)

type Order struct {
	database *sql.DB
}

func (st *Order) Insert(orderID string, data []byte) error {
	_, err := st.database.Exec(`
		INSERT INTO Orders (
			Order_Uid,
			Data
		) VALUES ($1, $2)
	`,
		orderID,
		data)

	if err != nil {
		return err
	}

	return nil
}

func (st *Order) FindAll() (ordersDTO map[string]models.OrderDTO, err error) {
	ordersDTO = make(map[string]models.OrderDTO)
	rows, err := st.database.Query(`
		SELECT
    		Order_Uid,
    		Data
		FROM Orders;
	`)

	if err != nil {
		return nil, fmt.Errorf("findAll: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.OrderUID,
			&order.Data,
		)
		if err != nil {
			return nil, fmt.Errorf("findAllOrder: %v", err)
		}

		var data models.OrderDTO
		err = json.Unmarshal(order.Data, &data)
		if err != nil {
			log.Println(err)

			return
		}

		ordersDTO[order.OrderUID] = data
	}

	return ordersDTO, nil
}

func NewStorageOrder(database *sql.DB) *Order {
	return &Order{
		database: database,
	}
}
