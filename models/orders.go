package models

import "time"

type Order struct {
	Id           int       `json:"id"`
	Table_id     int       `json:"table_id"`
	Created_at   time.Time `json:"created_at"`
	Is_paid      bool      `json:"is_paid"`
	Total_amount float32   `json:"total_amount"`
}
