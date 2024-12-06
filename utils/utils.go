package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"lem-in/models"
	"lem-in/vars"
)

func HandleError(err error) {
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
		if !sliceContainsString(vars.ConnectedRooms, v) {
			vars.ConnectedRooms = append(vars.ConnectedRooms, v)
		}
	}
}

func ValidColonyRooms(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if validRoomConnection(line) {
			storeConnectedRooms(line)
		}

		if strings.Contains(line, " ") {
			name, x, y, errRoom := getRoom(line)
			HandleError(errRoom)

			if sliceContainsString(vars.RoomNames, name) {
				log.Fatal("ERROR: invalid data format, room definition repeated")
			}

			storeRoom(name, x, y)
			if !sliceContainsString(vars.RoomNames, name) {
				vars.RoomNames = append(vars.RoomNames, name)
			}
		}
	}

	if len(vars.ConnectedRooms) != len(vars.RoomNames) {
		return false
	}

	sort.Strings(vars.RoomNames)
	sort.Strings(vars.ConnectedRooms)

	for i := range vars.ConnectedRooms {
		if vars.ConnectedRooms[i] != vars.RoomNames[i] {
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
	vars.AntsNumber = number
	return nil
}

func getRoom(line string) (string, int, int, error) {
	room := strings.Split(line, " ")
	if len(room) != 3 {
		return "", 0, 0, fmt.Errorf("invalid room details, %s", line)
	}
	name := room[0]

	x, err := strconv.Atoi(room[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid x coordinate: %v", err)
	}

	y, err := strconv.Atoi(room[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid y coordinate: %v", err)
	}

	return name, x, y, nil
}

func storeRoom(name string, x, y int) {
	room := models.Room{
		Name: name,
		X:    x,
		Y:    y,
	}
	vars.Rooms = append(vars.Rooms, room)
}

func ProcessLine(line string) {
	if vars.FirstLine {
		errNumberOfAnts := processNumberOfAnts(line)
		HandleError(errNumberOfAnts)
		vars.FirstLine = false
	} else if vars.IsStartNode {
		start, _, _, errRoom := getRoom(line)
		HandleError(errRoom)
		vars.StartRoom = start
		vars.IsStartNode = false
	} else if vars.IsEndNode {
		end, _, _, errRoom := getRoom(line)
		HandleError(errRoom)
		vars.EndRoom = end
		vars.IsEndNode = false
	} else if validRoomConnection(line) {
		rooms := strings.Split(line, "-")
		vars.Colony[rooms[0]] = append(vars.Colony[rooms[0]], rooms[1])
		vars.Colony[rooms[1]] = append(vars.Colony[rooms[1]], rooms[0])
	}
}

func HasStartAndEnd(file *os.File) bool {
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
