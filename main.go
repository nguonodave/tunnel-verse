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

	ants.MoveAnts(vars.PathMovement)
}
