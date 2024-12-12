package ants

import (
	"lem-in/models"
)

// AssignAnts distributes a specified number of ants across multiple paths in a round-robin fashion.
// Each path receives ants until its capacity (sum of rooms and ants) exceeds the next path's room count, 
// at which point the function shifts to the next path. The process continues until all ants are assigned.
func AssignAnts(paths []models.Path, ants int) {
	currentAnt := 1
	pathIndex := 0

	for currentAnt <= ants {
		paths[pathIndex].Ants = append(paths[pathIndex].Ants, currentAnt)

		if len(paths[pathIndex].Rooms)+len(paths[pathIndex].Ants) > len(paths[(pathIndex+1)%len(paths)].Rooms) {
			pathIndex = (pathIndex + 1) % len(paths)
		}
		currentAnt++
	}
}