package ants

import (
	"fmt"
	"strings"

	"lem-in/models"
	"lem-in/utils"
)

// MoveAnts simulates the movement of ants across multiple paths for a given number of turns.
// For each turn in every path, it calculates the new positions of ants on their respective paths, 
// and generates a list of movements formatted as "L<ant_number>-<room_name>". 
// After all the ants move in a given turn, their movements are printed as a space-separated string.
func MoveAnts(paths []models.Path) {
	maxTurns := utils.MaxTurns(paths)

	for turn := 1; turn <= maxTurns; turn++ {
		moves := []string{}

		for _, path := range paths {
			for antIndex, ant := range path.Ants {
				position := turn - antIndex - 1
				if position >= 0 && position < len(path.Rooms[1:]) {
					moves = append(moves, fmt.Sprintf("L%d-%s", ant, path.Rooms[1:][position]))
				}
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
