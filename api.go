package main

import (
	"cache/handlers"

	"github.com/labstack/echo/v4"
)

func StartServer() bool {
	Router := echo.New()
	Router.GET("/rooms", handlers.GetRooms)
	Router.GET("/rooms/:roomId/players", handlers.GetRoomPlayers)
	Router.POST("/rooms/:roomId/join", handlers.JoinRoom)
	Router.POST("/rooms/:roomId/quit", handlers.QuitRoom)
	Router.POST("/rooms", handlers.CreateRoom)
	Router.GET("/orders", handlers.GetAllOrders)
	Router.POST("/orders", handlers.CreateOrder)
	Router.GET("/players", handlers.GetAllPlayers)
	Router.GET("/players/:userId", handlers.GetOnePlayer)
	Router.PUT("/players/:userId", handlers.UpdateOnePlayer)
	Router.POST("/players/:userId/cards", handlers.AddCardToUser)
	Router.DELETE("/players/:userId/cards", handlers.DeleteCardFromUser)
	Router.GET("/players/:userId/cards", handlers.GetCardsByUser)
	Router.POST("/players", handlers.CreatePlayer)
	Router.Start(":1323")
	return true
}