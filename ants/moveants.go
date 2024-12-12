package ants

import (
	"fmt"
	"strings"

	"lem-in/models"
	"lem-in/utils"
)

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
				// fmt.Println(moves)
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
