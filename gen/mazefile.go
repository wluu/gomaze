package gen

import (
	"log"
	"math/rand"
	"os"
	"time"
)

// FileMaze is the text file which will contain the randomly generated maze
const FileMaze = "./maze.txt"

// MazeFile will generate a random maze and write it MazeFile; width and height
// dimensions need to be specified
func MazeFile(width, height int) {
	// if the file already exists, delete it
	if _, err := os.Stat(FileMaze); err == nil {
		os.Remove(FileMaze)
	}

	maze := genMaze(width, height)

	// write the maze to a file first
	for y := 0; y < height; y++ {
		northWalls := ""
		sideWalls := ""
		southWalls := ""

		for x := 0; x < width; x++ {
			walls := maze[x][y].w

			if walls.north {
				northWalls += "+--+"
			} else {
				northWalls += "+  +"
			}

			if walls.west {
				sideWalls += "|"
			} else {
				sideWalls += " "
			}

			if walls.east {
				sideWalls += "  |"
			} else {
				sideWalls += "   "
			}

			if walls.south {
				southWalls += "+--+"
			} else {
				southWalls += "+  +"
			}
		}

		file, err := os.OpenFile(MAZE_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteString(northWalls + "\n")
		file.WriteString(sideWalls + "\n")
		file.WriteString(southWalls + "\n")
	}
}

// x-y coordinates for the cells
type coord struct {
	x, y int
}

type walls struct {
	north bool
	east  bool
	south bool
	west  bool
}

type cell struct {
	visited bool
	w       walls
}

/*
	using a simple algorithm to randomly generate the maze in memory:
	https://en.wikipedia.org/wiki/Maze_generation_algorithm#Depth-first_search
*/
func genMaze(w, h int) [][]cell {
	maze := make([][]cell, w)
	for i := range maze {
		col := make([]cell, h)
		for j := range col {
			// by default, the booleans in the walls struct will be initialized to false
			// setting them to true to be more idiomatic
			col[j].w.north = true
			col[j].w.east = true
			col[j].w.south = true
			col[j].w.west = true
		}
		maze[i] = col
	}

	backtrack := []coord{}
	cur := coord{}

	// While there are unvisited cells
	for unvisitedCellsIn(maze) {
		// Make the initial cell the current cell and mark it as visited
		maze[cur.x][cur.y].visited = true

		// If the current cell has any neighbors which have not been visited
		if neighbors, ok := unvisitedNeighbors(cur, maze); ok {
			// 1. Choose randomly one of the unvisited neighbors
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randNeighbor := neighbors[r.Intn(len(neighbors))]

			// 2. Push the current cell to the stack
			backtrack = append(backtrack, cur)

			// 3. Remove the wall between the current cell and the chosen cell
			markWall(cur, randNeighbor, maze)

			// 4. Make the chosen cell the current cell and mark it as visited
			cur = randNeighbor
		} else if len(backtrack) > 0 { // 2. Else if stack is not empty
			// 1. Pop a cell from the stack i.e. get the last element from the slice
			// 2. Make it the current cell
			cur = backtrack[len(backtrack)-1]
			backtrack = backtrack[:len(backtrack)-1]
		}
	}

	return maze
}

// remove the wall between the current cell and the chosen neighbor
func markWall(cur, neighbor coord, maze [][]cell) {
	// NOTE: just want to see how this looks
	switch {
	// on the north/south plane
	case cur.x-neighbor.x == 0:
		switch {
		// south wall (relative to cur)
		case cur.y < neighbor.y:
			maze[cur.x][cur.y].w.south = false
			maze[neighbor.x][neighbor.y].w.north = false

		// north wall (relative to cur)
		case neighbor.y < cur.y:
			maze[cur.x][cur.y].w.north = false
			maze[neighbor.x][neighbor.y].w.south = false
		}

	// on the east/west plane
	case cur.y-neighbor.y == 0:
		switch {
		// east wall (relative to cur)
		case cur.x < neighbor.x:
			maze[cur.x][cur.y].w.east = false
			maze[neighbor.x][neighbor.y].w.west = false

		// west wall (relative to cur)
		case neighbor.x < cur.x:
			maze[cur.x][cur.y].w.west = false
			maze[neighbor.x][neighbor.y].w.east = false
		}
	}
}

// returns a list of unvisited cells
func unvisitedNeighbors(cur coord, maze [][]cell) ([]coord, bool) {
	width := len(maze)
	height := len(maze[0])
	neighbors := []coord{}

	addNeighbor := func(x, y int) {
		if !maze[x][y].visited {
			neighbors = append(neighbors, coord{x: x, y: y})
		}
	}

	// didn't pass north border
	if x, y := cur.x, cur.y-1; y >= 0 {
		addNeighbor(x, y)
	}

	// didn't pass east border
	if x, y := cur.x+1, cur.y; x < width {
		addNeighbor(x, y)
	}

	// didn't pass south border
	if x, y := cur.x, cur.y+1; y < height {
		addNeighbor(x, y)
	}

	// didn't pass west border
	if x, y := cur.x-1, cur.y; x >= 0 {
		addNeighbor(x, y)
	}

	return neighbors, len(neighbors) > 0
}

// returns true if there are any unvisited cells in the entire maze; otherwise, returns false
func unvisitedCellsIn(maze [][]cell) bool {
	for _, col := range maze {
		for _, cell := range col {
			if !cell.visited {
				return true
			}
		}
	}
	return false
}
