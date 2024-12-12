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

func SliceContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func SliceContainsSlice(arr1, arr2 []string) bool {
	for _, v := range arr2 {
		if SliceContainsString(arr1, v) {
			return true
		}
	}
	return false
}

func SliceInSlices(arr1 [][]string, arr2 []string) bool {
	for _, v := range arr2 {
		for _, w := range arr1 {
			if SliceContainsString(w, v) {
				return true
			}
		}
	}
	return false
}

func ValidRoomConnection(line string) bool {
	rooms := strings.Split(line, "-")
	return len(rooms) == 2 &&
		strings.Contains(line, "-") &&
		rooms[0] != "" &&
		rooms[1] != "" &&
		!strings.Contains(rooms[0], " ") &&
		!strings.Contains(rooms[1], " ") &&
		rooms[0] != rooms[1]
}

func StoreConnectedRooms(line string) {
	rooms := strings.Split(line, "-")
	for _, v := range rooms {
		if !SliceContainsString(vars.ConnectedRooms, v) {
			vars.ConnectedRooms = append(vars.ConnectedRooms, v)
		}
	}
}

func ValidColonyRooms(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if ValidRoomConnection(line) {
			StoreConnectedRooms(line)
		} else if strings.Contains(line, "-") && !ValidRoomConnection(line) {
			log.Fatalf("ERROR: invalid data format, invalid room link %s", line)
		}

		if strings.Contains(line, " ") {
			name, x, y, errRoom := GetRoom(line)
			HandleError(errRoom)

			if SliceContainsString(vars.RoomNames, name) {
				log.Fatal("ERROR: invalid data format, room definition repeated")
			}

			StoreRoom(name, x, y)
			if !SliceContainsString(vars.RoomNames, name) {
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

func ProcessNumberOfAnts(line string) error {
	number, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("invalid number of ants: %v", err)
	}
	vars.AntsNumber = number
	return nil
}

func GetRoom(line string) (string, int, int, error) {
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

func StoreRoom(name string, x, y int) {
	room := models.Room{
		Name: name,
		X:    x,
		Y:    y,
	}
	vars.Rooms = append(vars.Rooms, room)
}

func ProcessLine(line string) {
	if vars.FirstLine {
		errNumberOfAnts := ProcessNumberOfAnts(line)
		HandleError(errNumberOfAnts)
		vars.FirstLine = false
	} else if vars.IsStartNode {
		start, _, _, errRoom := GetRoom(line)
		HandleError(errRoom)
		vars.StartRoom = start
		vars.IsStartNode = false
	} else if vars.IsEndNode {
		end, _, _, errRoom := GetRoom(line)
		HandleError(errRoom)
		vars.EndRoom = end
		vars.IsEndNode = false
	} else if ValidRoomConnection(line) {
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

func MaxTurns(paths []models.Path) int {
	maxTurns := 1
	for _, path := range paths {
		rooms := path.Rooms[1:len(path.Rooms)-1]
		ants := path.Ants
		turns := len(rooms) + len(ants)

		if turns > maxTurns {
			maxTurns = turns
		}
	}
	return maxTurns
}
