package models

type MenuItem struct {
	ID    int     `json:"id"`
	name  string  `json:"name"`
	price float32 `json:"price"`
}
