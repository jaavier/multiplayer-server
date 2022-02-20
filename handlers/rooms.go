package handlers

import (
	"cache/models"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

var Rooms []models.Room

func GetRooms (c echo.Context) error {
	json, _ := json.Marshal(Rooms)
	return c.String(200, string(json));
}

func CreateRoom (c echo.Context) error {
	var room models.Room
	json.NewDecoder(c.Request().Body).Decode(&room)
	_, err := Client.Do("SADD", "rooms", string(room.Id))
	if err != nil {
		return c.JSONPretty(500, "Error creating room. Id already exists", "  ")
	}
	return c.String(200, "Room Created!")
}

func JoinRoom (c echo.Context) error {
	roomId := c.Param("roomId")
	var json_map = map[string]interface{}{
		"userId": "",
	}
	json.NewDecoder(c.Request().Body).Decode(&json_map)
	_, err := Client.Do("LINDEX", "rooms", roomId)
	if err != nil {
		return c.JSONPretty(500, "Error joining room (Doesn't exist)", "  ")
	}
	sadd, _ := Client.Do("SADD", roomId, json_map["userId"].(string))
	if sadd == int64(0) {
		return c.JSONPretty(500, "Error joining room, you're already in", "  ")
	}
	return c.String(200, "Now you're in the room!")
}

func GetRoomPlayers (c echo.Context) error {
	roomId := c.Param("roomId")
	players, err := Client.Do("SMEMBERS", roomId)
	members := []string{}
	for _, player := range players.([]interface{}) {
		data := fmt.Sprintf("%s", player)
		members = append(members, data)
	}
	if err != nil {
		return c.JSONPretty(500, "Error getting room players", "  ")
	}
	return c.JSONPretty(200, members, "  ")
}

func QuitRoom (c echo.Context) error {
	roomId := c.Param("roomId")
	var json_map = map[string]interface{}{
		"userId": "",
	}
	json.NewDecoder(c.Request().Body).Decode(&json_map)
	_, err := Client.Do("LINDEX", "rooms", roomId)
	if err != nil {
		return c.JSONPretty(500, "Error quitting room (Doesn't exist)", "  ")
	}
	srem, _ := Client.Do("SREM", roomId, json_map["userId"].(string))
	if srem == int64(0) {
		return c.JSONPretty(500, "Error quitting room, you're not in", "  ")
	}
	return c.String(200, "You're out of the room!")
}