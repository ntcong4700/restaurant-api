package models

type Table struct {
	ID           int    `json:"id"`
	Is_available string `json:"status"`
}
