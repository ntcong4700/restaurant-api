package repositories

import (
	"context"
	"restaurant-api/config"
	"restaurant-api/models"
)

func CreateOrderTx(order *models.Order, items []models.OrderDetail) error {
	tx, err := config.Db.Begin()

}
