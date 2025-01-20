package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"lem-in/ants"
	"lem-in/processpaths"
	"lem-in/utils"
	"lem-in/vars"
)

// - Checks if a valid file name is provided as a command-line argument.
// - Opens and reads the specified file, ensuring error handling for file operations.
// - Validates the file format, ensuring that start and end rooms are defined, and the room and connection data is valid.
// - Processes each line in the file, identifying special markers for the start and end rooms, and processing room data accordingly.
// - Verifies that at least one ant is available for movement.
// - Finds the optimal path from the start room to the end room and assigns ants to the paths.
// - Moves the ants along the selected optimal path.
func main() {
	if len(os.Args) != 2 {
		log.Println("ERROR: missing file name")
		return
	}

	file, errOpenFile := os.Open(os.Args[1])
	utils.HandleError(errOpenFile)
	defer file.Close()

	if !utils.HasStartAndEnd(file) {
		log.Fatal("ERROR: invalid data format, no start or end room found")
	}

	if !utils.ValidColonyRooms(file) {
		log.Fatal("ERROR: invalid data format, your rooms and room links do not match")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if utils.Comment(line) {
			continue
		}
		if line == "##start" {
			vars.IsStartNode = true
			continue
		} else if line == "##end" {
			vars.IsEndNode = true
			continue
		}
		utils.ProcessLine(line)
	}

	if vars.AntsNumber < 1 {
		log.Fatal("ERROR: invalid data format, no ants to move in colony")
	}

	processpaths.FindPaths(vars.StartRoom, vars.EndRoom)
	processpaths.OptimalPathMovement()

	utils.PrintColony()
	ants.MoveAnts(vars.PathMovement)
}
