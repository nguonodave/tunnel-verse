
# Tunnel Verse: Ant Farm Simulator

Lem-in is a Go program that simulates ants navigating through a digital ant farm (a network of rooms and tunnels). The program finds the most efficient path for ants to travel from the start room to the end room while respecting movement constraints.

## Features

- Parses ant farm descriptions from input files
- Validates colony structure and ant movements
- Implements pathfinding algorithms to determine optimal routes
- Displays ant movements step-by-step
- Handles various error cases with appropriate messages

## Input Format

The input file must describe:
- Number of ants
- Rooms (with coordinates)
- Tunnels connecting rooms
- Start and end markers (`##start` and `##end`)

Example:
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

## Output Format

The program displays:
1. The input file content
2. Each step showing ant movements in format `Lx-y` where:
   - `x` is the ant number
   - `y` is the destination room

Example:
```
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
```

## Installation & Usage

1. Clone the repository
2. Build and run with Go:

```bash
go run main.go [input_file]
```

Example:

```bash
go run main.go test1.txt
```

## Testing
Run unit tests with:

```bash
go test ./...
```

## Constraints
- Only one ant per room (except start/end)
- Each tunnel can be used once per turn
- Ants must take the shortest available path
- All moves must be valid

## Error Handling
The program will return `ERROR: invalid data format` for:

1. Missing start/end rooms
2. Invalid room formats
3. Duplicate rooms
4. Invalid tunnels
5. Other formatting issues
