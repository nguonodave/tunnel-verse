package vars

import "lem-in/models"

var (
	AntsNumber     int
	FirstLine      = true
	IsStartNode    = false
	IsEndNode      = false
	StartRoom      string
	EndRoom        string
	RoomName       string
	RoomNames      []string
	ConnectedRooms []string
	Colony         = make(map[string][]string)
	Rooms          []models.Room
)
