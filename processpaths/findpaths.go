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
