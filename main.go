package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	antsNumber int
	firstLine  = true
	startNode  = false
	endNode    = false
	roomName string
	xCord int
	yCord int
	errCord error
	colony = make(map[string][]string)
	rooms []Room
)

type Room struct {
	Name string
	X int
	Y int
}

func handleError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func validRoomConnection(line string) bool {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return false
	}
	room1, room2 := parts[0], parts[1]
	if room1 == "" || room2 == "" {
		return false
	}
	if strings.Contains(room1, " ") || strings.Contains(room2, " ") {
		return false
	}
	return true
}

func processNumberOfAnts(line string) error {
	number, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("invalid number of ants: %v", err)
	}
	antsNumber = number
	return nil
}

func getRoom(line string) (string, int, int, error) {
	room := strings.Split(line, " ")
	if len(room) != 3 {
		return "", 0, 0, fmt.Errorf("invalid room details, %s", line)
	}
	roomName = room[0]

	x, err := strconv.Atoi(room[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid x coordinate: %v", err)
	}

	y, err := strconv.Atoi(room[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid y coordinate: %v", err)
	}

	return roomName, x, y, nil
}

func storeRoom(name string, x, y int) {
	room := Room {
		Name: name,
		X: x,
		Y: y,
	}
	rooms = append(rooms, room)
}

func processLine(line string) {
	if firstLine {
		errNumberOfAnts := processNumberOfAnts(line)
		handleError(errNumberOfAnts)
		firstLine = false
	} else if startNode {
		startRoom, x, y, errRoom := getRoom(line)
		handleError(errRoom)
		fmt.Println(startRoom)
		storeRoom(startRoom, x, y)
		startNode = false
	} else if endNode {
		endRoom, x, y, errRoom := getRoom(line)
		handleError(errRoom)
		fmt.Println(endRoom)
		storeRoom(endRoom, x, y)
		endNode = false
	} else if strings.Contains(line, "-") {
		if validRoomConnection(line) {
			rooms := strings.Split(line, "-")
			colony[rooms[0]] = append(colony[rooms[0]], rooms[1])
			colony[rooms[1]] = append(colony[rooms[1]], rooms[0])
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Println("ERROR: missing file name")
		return
	}

	file, errOpenFile := os.Open(os.Args[1])
	handleError(errOpenFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "##start" {
			startNode = true
			continue
		} else if line == "##end" {
			endNode = true
			continue
		}
		processLine(line)
	}

	fmt.Println(colony)
	fmt.Println(rooms)
}
