package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"lem-in/models"
)

var (
	antsNumber     int
	firstLine      = true
	startNode      = false
	endNode        = false
	startRoom string
	endRoom string
	roomName       string
	roomNames      []string
	connectedRooms []string
	xCord          int
	yCord          int
	errCord        error
	colony         = make(map[string][]string)
	rooms          []models.Room
)

func handleError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func sliceContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func validRoomConnection(line string) bool {
	rooms := strings.Split(line, "-")
	return len(rooms) == 2 &&
		strings.Contains(line, "-") &&
		rooms[0] != "" &&
		rooms[1] != "" &&
		!strings.Contains(rooms[0], " ") &&
		!strings.Contains(rooms[1], " ")
}

func storeConnectedRooms(line string) {
	rooms := strings.Split(line, "-")
	for _, v := range rooms {
		if !sliceContainsString(connectedRooms, v) {
			connectedRooms = append(connectedRooms, v)
		}
	}
}

func validColonyRooms(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if validRoomConnection(line) {
			storeConnectedRooms(line)
		}

		if strings.Contains(line, " ") {
			roomName, x, y, errRoom := getRoom(line)
			handleError(errRoom)

			if sliceContainsString(roomNames, roomName) {
				log.Fatal("ERROR: invalid data format, room definition repeated")
			}

			storeRoom(roomName, x, y)
			if !sliceContainsString(roomNames, roomName) {
				roomNames = append(roomNames, roomName)
			}
		}
	}

	if len(connectedRooms) != len(roomNames) {
		return false
	}

	sort.Strings(roomNames)
	sort.Strings(connectedRooms)

	for i := range connectedRooms {
		if connectedRooms[i] != roomNames[i] {
			return false
		}
	}
	file.Seek(0, 0)
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
	room := models.Room{
		Name: name,
		X:    x,
		Y:    y,
	}
	rooms = append(rooms, room)
}

func processLine(line string) {
	if firstLine {
		errNumberOfAnts := processNumberOfAnts(line)
		handleError(errNumberOfAnts)
		firstLine = false
	} else if startNode {
		start, _, _, errRoom := getRoom(line)
		handleError(errRoom)
		startRoom = start
		startNode = false
	} else if endNode {
		end, _, _, errRoom := getRoom(line)
		handleError(errRoom)
		endRoom = end
		endNode = false
	} else if validRoomConnection(line) {
		rooms := strings.Split(line, "-")
		colony[rooms[0]] = append(colony[rooms[0]], rooms[1])
		colony[rooms[1]] = append(colony[rooms[1]], rooms[0])
	}
}

func hasStartAndEnd(file *os.File) bool {
	hasStart := false
	hasEnd := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "##start" {
			hasStart = true
		} else if line == "##end" {
			hasEnd = true
		}
		if hasStart && hasEnd {
			file.Seek(0, 0)
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		log.Println("ERROR: missing file name")
		return
	}

	file, errOpenFile := os.Open(os.Args[1])
	handleError(errOpenFile)
	defer file.Close()

	if !hasStartAndEnd(file) {
		log.Fatal("ERROR: invalid data format, no start or end room found")
	}

	if !validColonyRooms(file) {
		log.Fatal("ERROR: invalid data format, your room links diverge from the rooms")
	}

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
	fmt.Println(roomNames)
	fmt.Println(connectedRooms)
	fmt.Println(startRoom)
	fmt.Println(endRoom)
}
