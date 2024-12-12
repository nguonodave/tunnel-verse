package ants

import (
	"lem-in/models"
)

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