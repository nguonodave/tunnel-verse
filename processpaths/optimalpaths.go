package processpaths

import (
	"math"
	"sort"

	"lem-in/ants"
	"lem-in/models"
	"lem-in/utils"
	"lem-in/vars"
)

// GetOptimalPaths1 selects the optimal paths from a list of paths (arr) based on certain criteria.
// The function first sorts the paths by length, then iterates through each path to determine if it should be included in the result.
// A path is added if it meets one of the following conditions:
//   - If the number of rooms in the current path is less than or equal to half of the total number of ants (rounded),
//     and it shares some rooms with the first path but has a different length.
//   - If the current path does not already exist in any of the existing paths' room sets.
//
// The function returns a slice of paths that meet these conditions.
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

// GetOptimalPaths2 selects unique paths from a list of paths (arr) based on room uniqueness.
// The function first sorts the paths by length, then iterates through each path, adding it to the result if it is not already
// present in any of the previously selected paths' room sets.
// The function returns a slice of paths with unique room configurations.
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

// OptimalPathMovement assigns ants to the optimal paths based on movement efficiency.
// The function calculates two sets of optimal paths (path1 and path2) using GetOptimalPaths1 and GetOptimalPaths2.
// It then converts these paths into models.Path and assigns ants to each set using the AssignAnts function.
// The function compares the number of turns required to complete each path and selects the path with the fewest turns for movement.
// The optimal path is stored in the global vars.PathMovement.
func OptimalPathMovement() {
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

	ants.AssignAnts(pathComb1, vars.AntsNumber)
	ants.AssignAnts(pathComb2, vars.AntsNumber)

	turns1 := utils.MaxTurns(pathComb1)
	turns2 := utils.MaxTurns(pathComb2)

	if turns1 < turns2 {
		vars.PathMovement = pathComb1
		return
	}

	vars.PathMovement = pathComb2
}
