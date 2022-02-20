package handlers

import (
	"cache/models"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

var Orders []models.Order


func CreateOrderInRedis(order models.Order) bool {
	fmt.Println("Creating order...")
	orderMarshal, _ := json.Marshal(order)
	Client.Do("SET", order.UserId, string(orderMarshal))
	return true
}

func GetOrder(UserId string) string {
	res, _ := Client.Do("GET", UserId)
	return res.(string)
}

func GetAllOrders(c echo.Context) error {
	toJson, _ := json.Marshal(Orders)
	return c.String(200, string(toJson))
}

func CreateOrder(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	json_order := models.Order{}
	json.NewDecoder(c.Request().Body).Decode(&json_order)
	CreateOrderInRedis(models.Order{
		UserId: json_order.UserId,
		Product: json_order.Product,
		HashTx: json_order.HashTx,
		Price: json_order.Price,
	})
	j, _ := json.Marshal(json_order)
	fmt.Println(string(j))
	return c.String(200, string(j))
}
