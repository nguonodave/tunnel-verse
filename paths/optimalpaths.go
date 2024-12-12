package paths

import (
	"fmt"
	"math"
	"sort"

	"lem-in/ants"
	"lem-in/models"
	"lem-in/utils"
	"lem-in/vars"
)

func GetOptimalPaths1(arr [][]string) [][]string {
	paths := [][]string{}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})

	paths = append(paths, arr[0])

	for i := 1; i < len(arr); i++ {
		firstPath := paths[0]
		firstPathRooms := firstPath[1 : len(firstPath)-1]
		currentPathRooms := arr[i][1 : len(arr[i])-1]
		val := float64(vars.AntsNumber) / 2

		if len(currentPathRooms) <= int(math.Round(val)) && utils.SliceContainsSlice(paths[0], currentPathRooms) && len(currentPathRooms) != len(firstPathRooms) {
			paths = paths[1:]
			paths = append(paths, arr[i])
		} else if !utils.SliceInSlices(paths, currentPathRooms) {
			paths = append(paths, arr[i])
		}
	}

	return paths
}

func GetOptimalPaths2(arr [][]string) [][]string {
	paths := [][]string{}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})

	paths = append(paths, arr[0])

	for i := 1; i < len(arr); i++ {
		currentPathRooms := arr[i][1 : len(arr[i])-1]
		if !utils.SliceInSlices(paths, currentPathRooms) {
			paths = append(paths, arr[i])
		}
	}

	return paths
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

func OptimalPath() []models.Path {
	path1 := GetOptimalPaths1(vars.AllPaths)
	path2 := GetOptimalPaths2(vars.AllPaths)

	pathComb1 := []models.Path{}
	for _, v := range path1 {
		path := models.Path{
			Rooms: v,
		}
		pathComb1 = append(pathComb1, path)
	}

	pathComb2 := []models.Path{}
	for _, v := range path2 {
		path := models.Path{
			Rooms: v,
		}
		pathComb2 = append(pathComb2, path)
	}

	// assign ants
	// calculate number of turns
	// select the one with the least number of turns

	ants.AssignAnts(pathComb1, vars.AntsNumber)
	ants.AssignAnts(pathComb2, vars.AntsNumber)

	fmt.Println(pathComb1)
	fmt.Println(pathComb2)

	turns1 := MaxTurns(pathComb1)
	turns2 := MaxTurns(pathComb2)

	fmt.Println(turns1)
	fmt.Println(turns2)

	if turns1 < turns2 {
		return pathComb1
	}

	return pathComb2
}
