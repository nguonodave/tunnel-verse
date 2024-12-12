package models

type Room struct {
	Name string
	X    int
	Y    int
}

type Path struct {
	Rooms []string
	Ants  []int
}