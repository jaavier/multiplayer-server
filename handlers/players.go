package handlers

import (
	"cache/models"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
)

var Client, _ = redis.Dial("tcp", ":6379")

func CreatePlayer(c echo.Context) error {
	var player models.Player
	json.NewDecoder(c.Request().Body).Decode(&player)
	toBytes, _ := json.Marshal(player)
	res, err := Client.Do("SET", fmt.Sprintf("%v-profile", player.UserId), string(toBytes))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	print("res", res)
	print("err", err)
	fmt.Println("CreatePlayer: ", player)
	return c.String(200, "Player Created!")
}

func GetAllPlayers(c echo.Context) error {
	players, err := Client.Do("KEYS", "*-profile")
	if err != nil {
		return c.JSONPretty(500, "Error getting players", "  ")
	}
	fmt.Println("GetAllPlayers: ", players)
	members := []string{}
	for _, player := range players.([]interface{}) {
		data := fmt.Sprintf("%s", player)
		members = append(members, strings.Replace(data, "-profile", "", 1))
	}
	return c.JSONPretty(200, members, "  ")
}

func GetOnePlayer(c echo.Context) error {
	var pModel models.Player
	userId := c.Param("userId")
	fmt.Println("GetOnePlayer: ", fmt.Sprintf("%v-profile", userId))
	player, _ := Client.Do("GET", fmt.Sprintf("%v-profile", userId))
	if player == nil {
		return c.JSONPretty(200, "Error", "  ")
	}
	toJSON := string(player.([]byte))
	json.Unmarshal([]byte(toJSON), &pModel)
	return c.JSONPretty(200, pModel, "  ")
}

func UpdateOnePlayer(c echo.Context) error {
	userId := c.Param("userId")
	var player models.Player
	json.NewDecoder(c.Request().Body).Decode(&player)
	toBytes, _ := json.Marshal(player)
	res, err := Client.Do("SET", userId, string(toBytes))
	print("res", res)
	print("err", err)
	return c.String(200, "Player Updated!")
}

func GetPlayersPlaying(c echo.Context) error {
	players, err := Client.Do("KEYS", "players-playing")
	if err != nil {
		return c.JSONPretty(500, "Error getting players playing", "  ")
	}
	return c.JSONPretty(200, players, "  ")
}

func RemovePlayerPlaying(c echo.Context) error {
	userId := c.Param("userId")
	players, err := Client.Do("SREM", "players-playing", userId)
	if err != nil {
		return c.JSONPretty(500, "Error getting players playing", "  ")
	}
	return c.JSONPretty(200, players, "  ")
}

