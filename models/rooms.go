package models

type Room struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Status     int      `json:"status"`
	MinPlayers int      `json:"minPlayers"`
	MaxPlayers int      `json:"maxPlayers"`
	Mode       string   `json:"mode"`
	Players    []Player `json:"players"`
}
