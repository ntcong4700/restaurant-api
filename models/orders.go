package models

import "time"

type Order struct {
	id           int       `json:"id"`
	table_id     int       `json:"table_id"`
	created_at   time.Time `json:"created_at"`
	is_paid      bool      `json:"is_paid"`
	total_amount float32   `json:"total_amount"`
}
