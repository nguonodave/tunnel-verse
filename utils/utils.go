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

var (
	startRoom string
	endRoom   string
	roomLinks []string
)

func HandleError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

// SliceContainsString checks if a given string (s) is present in the provided slice of strings (arr).
func SliceContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// SliceContainsSlice checks if any element from the second slice (arr2) is present in the first slice (arr1).
func SliceContainsSlice(arr1, arr2 []string) bool {
	for _, v := range arr2 {
		if SliceContainsString(arr1, v) {
			return true
		}
	}
	return false
}

// SliceInSlices checks if any element from the second slice (arr2) is present in any of the sub-slices of the first slice (arr1).
// The function iterates over each element of arr2 and, for each element, checks if it exists in any sub-slice of arr1
// using the helper function SliceContainsString. If a match is found, the function returns true.
// If no elements from arr2 are found in any sub-slice of arr1, it returns false.
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

// ValidCoordinates checks if the given coordinates (x, y) are valid by ensuring that no existing room
// in the global vars.Rooms slice has the same coordinates.
func ValidCoordinates(x, y int) bool {
	for _, v := range vars.Rooms {
		if v.X == x && v.Y == y {
			return false
		}
	}
	return true
}

// ValidRoomConnection checks whether a given line represents a valid room connection.
// The function splits the line by a hyphen ("-") and validates the following conditions:
// - The line contains exactly one hyphen separating two rooms.
// - Both rooms are non-empty and do not contain spaces.
// - The two rooms are not the same.
func ValidRoomConnection(line string) bool {
	rooms := strings.Split(line, "-")
	return len(rooms) == 2 &&
		strings.Contains(line, "-") &&
		rooms[0] != "" &&
		rooms[1] != "" &&
		string(rooms[0][0]) != "#" &&
		string(rooms[0][0]) != "L" &&
		string(rooms[1][0]) != "#" &&
		string(rooms[1][0]) != "L" &&
		!strings.Contains(rooms[0], " ") &&
		!strings.Contains(rooms[1], " ") &&
		rooms[0] != rooms[1]
}

// StoreConnectedRooms processes a line representing a connection between two rooms,
// and stores the unique room names in the global vars.ConnectedRooms slice.
// The function splits the line by a hyphen ("-"), iterates over the resulting room names,
// and appends each room to vars.ConnectedRooms only if it is not already present in the slice.
func StoreConnectedRooms(line string) {
	rooms := strings.Split(line, "-")
	for _, v := range rooms {
		if !SliceContainsString(vars.ConnectedRooms, v) {
			vars.ConnectedRooms = append(vars.ConnectedRooms, v)
		}
	}
}

// ValidColonyRooms validates room and connection data in the provided file.
// It checks for valid room connections, ensures no duplicate room names or coordinates, and stores valid rooms.
// After processing, it compares the number of connected rooms to the total room names, sorting and comparing them.
// The function returns false if any validation fails, and true if all checks pass.
// It uses file.Seek(0, 0) to reset the file pointer to the beginning for potential further operations.
func ValidColonyRooms(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if Comment(line) {
			continue
		}
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

			if !ValidCoordinates(x, y) {
				log.Fatalf("ERROR: invalid data format, duplicated coordinates on %s", line)
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

// ProcessNumberOfAnts parses a line representing the number of ants and updates the global ants count.
func ProcessNumberOfAnts(line string) error {
	number, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("invalid number of ants: %v", err)
	}
	vars.AntsNumber = number
	return nil
}

// GetRoom parses a line of input representing a room's details and returns the room's name and coordinates (x, y).
// The function expects the line to contain exactly three space-separated values:
// - The room name (string)
// - The x-coordinate (integer)
// - The y-coordinate (integer)
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

// StoreRoom strores and struct room with its name and x and y coordinates
func StoreRoom(name string, x, y int) {
	room := models.Room{
		Name: name,
		X:    x,
		Y:    y,
	}
	vars.Rooms = append(vars.Rooms, room)
}

// ProcessLine processes a single line of input, updating global variables based on its content.
// The function performs the following steps:
// - If the line represents a valid room connection, it updates the colony's room connectivity (graph).
// - The graph is in form of a map where the keys are the nodes, and values, all its neighbors.
// Each step also includes error handling for potential issues encountered during processing.
func ProcessLine(line string) {
	if vars.FirstLine {
		errNumberOfAnts := ProcessNumberOfAnts(line)
		HandleError(errNumberOfAnts)
		vars.FirstLine = false
	} else if vars.IsStartNode {
		start, _, _, errRoom := GetRoom(line)
		HandleError(errRoom)
		vars.StartRoom = start
		startRoom = line
		vars.IsStartNode = false
	} else if vars.IsEndNode {
		end, _, _, errRoom := GetRoom(line)
		HandleError(errRoom)
		endRoom = line
		vars.EndRoom = end
		vars.IsEndNode = false
	} else if ValidRoomConnection(line) {
		roomLinks = append(roomLinks, line)
		rooms := strings.Split(line, "-")
		vars.Colony[rooms[0]] = append(vars.Colony[rooms[0]], rooms[1])
		vars.Colony[rooms[1]] = append(vars.Colony[rooms[1]], rooms[0])
	}
}

// HasStartAndEnd check if the colony has a start and end room,
// It returns true if it does and false otherwise.
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

// MaxTurns calculates the nuber of turn in each selected paths
func MaxTurns(paths []models.Path) int {
	maxTurns := 1
	for _, path := range paths {
		rooms := path.Rooms[1 : len(path.Rooms)-1]
		ants := path.Ants
		turns := len(rooms) + len(ants)

		if turns > maxTurns {
			maxTurns = turns
		}
	}
	return maxTurns
}

func Comment(s string) bool {
	if s != "" && string(s[0]) == "#" && s != "##start" && s != "##end" {
		return true
	}
	return false
}

func PrintColony() {
	fmt.Println(vars.AntsNumber)

	fmt.Println("##start")
	fmt.Println(startRoom)

	for _, v := range vars.Rooms {
		room := fmt.Sprintf("%s %d %d", v.Name, v.X, v.Y)
		if room != startRoom && room != endRoom {
			fmt.Println(room)
		}
	}

	fmt.Println("##end")
	fmt.Println(endRoom)

	for _, v := range roomLinks {
		fmt.Println(v)
	}

	fmt.Println()
}
