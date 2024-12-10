package paths

import (
	"sort"

	"lem-in/utils"
	"lem-in/vars"
)

func GetOptimalPaths(arr [][]string) [][]string {
	paths := [][]string{}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})

	paths = append(paths, arr[0])
	
	for i := 1; i < len(arr); i++ {
		initialPath := paths[len(paths)-1]
		initialPathRooms := initialPath[1:len(initialPath)-1]
		currentPathRooms := arr[i][1:len(arr[i])-1]

		if len(currentPathRooms) <= vars.AntsNumber/2 && !utils.SliceContainsSlice(initialPathRooms, currentPathRooms) {
			paths = append(paths, arr[i])
		}
	}

	return paths
}
