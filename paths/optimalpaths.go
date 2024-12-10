package paths

import (
	"math"
	"sort"

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
