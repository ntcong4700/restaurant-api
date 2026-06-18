package models

type OrderDetail struct {
	Id           int     `json:"id"`
	Order_id     int     `json:"order_id"`
	Menu_item_id int     `json:"menu_item_id"`
	Quantity     int     `json:"quantity"`
	Price        float32 `json:"price"`
}
