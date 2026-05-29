package models

type OrderDetail struct {
	id           int     `json:"id"`
	order_id     int     `json:"order_id"`
	menu_item_id int     `json:"menu_item_id"`
	quantity     int     `json:"quantity"`
	price        float32 `json:"price"`
}
