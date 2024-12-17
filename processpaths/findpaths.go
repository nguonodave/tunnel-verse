package processpaths

import (
	"lem-in/utils"
	"lem-in/vars"
)

var (
	visited = []string{}
	path    = []string{}
	current string
)

// FindPaths recursively searches for all possible paths from the 'start' room to the 'end' room.
// The function tracks the current path and visited rooms to avoid revisiting them.
// If the 'current' room is the same as the 'end' room, the current path is saved to vars.AllPaths.
// The function explores each neighboring room, and for unvisited rooms, it recursively calls itself to continue the search.
// After exploring all neighbors, it backtracks by removing the last room from the path and visited lists.
func FindPaths(start, end string) {
	visited = append(visited, start)
	path = append(path, start)
	current = start

	if current == end {
		vars.AllPaths = append(vars.AllPaths, append([]string(nil), path...))
	}

	for _, neighbor := range vars.Colony[current] {
		if !utils.SliceContainsString(visited, neighbor) {
			FindPaths(neighbor, end)
		}
	}

	path = path[0 : len(path)-1]
	visited = visited[0 : len(visited)-1]
}
