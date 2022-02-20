package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetCardsByUser (c echo.Context) error {
	userId := c.Param("userId")
	cards, err := Client.Do("SMEMBERS", fmt.Sprintf("%v-cards", userId))
	cardsList := []string{}
	for _, card := range cards.([]interface{}) {
		data := fmt.Sprintf("%s", card)
		cardsList = append(cardsList, data)
	}
	if err != nil {
		return c.JSONPretty(500, "Error getting cards", "  ")
	}
	return c.JSONPretty(200, cardsList, "  ")
}

func AddCardToUser (c echo.Context) error {
	userId := c.Param("userId")
	json_map := map[string]interface{}{
		"cardId": "",
	}
	json.NewDecoder(c.Request().Body).Decode(&json_map)
	player, _ := Client.Do("GET", fmt.Sprintf("%v-profile", userId))
	if player == nil {
		return c.JSONPretty(500, "Error adding card to player because it doesn't exist", "  ")
	}
	value, _ := Client.Do("SADD", fmt.Sprintf("%v-cards", userId), json_map["cardId"].(string))
	if value == int64(0) {
		return c.JSONPretty(500, "Error adding card because it was added before", "  ")
	}
	return c.String(200, "Card added!")
}

func DeleteCardFromUser (c echo.Context) error {
	userId := c.Param("userId")
	json_map := map[string]interface{}{
		"cardId": "",
	}
	json.NewDecoder(c.Request().Body).Decode(&json_map)
	player, _ := Client.Do("GET", fmt.Sprintf("%v-profile", userId))
	if player == nil {
		return c.JSONPretty(500, "Error removing card from player because it doesn't exist", "  ")
	}
	value, _ := Client.Do("SREM", fmt.Sprintf("%v-cards", userId), json_map["cardId"].(string))
	if value == int64(0) {
		return c.JSONPretty(500, "Error removing card because it wasn't added before", "  ")
	}
	return c.String(200, "Card removed!")
}