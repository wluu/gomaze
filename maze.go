package main

import (
	"fmt"
	_ "os"
	"time"
	"math/rand"
	tl "github.com/JoelOtter/termloop"
)

const (
	WIDTH  = 50
	HEIGHT = 50
)

func main() {
	g := tl.NewGame()

	l := tl.NewBaseLevel(tl.Cell{})
	g.Screen().SetLevel(l)

	maze := genMaze(WIDTH, HEIGHT)

	// drawing the maze
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			cell := ""
			wall := maze[x][y].w

			if wall.north {
				cell += "+--+\n"
			} else {
				cell += "\n"
			}

			if wall.west || wall.east {
				if wall.west {
					cell += "╎"
					if wall.east {
						cell += "  ╎\n"
					} else {
						cell += "\n"
					}
				} else if wall.east {
					cell += "   ╎\n"
				}
			} else {
				cell += "\n\n"
			}

			if wall.south {
				cell += "+--+"
			}

			fmt.Println(cell)

			canvasCell := tl.CanvasFromString(cell)
			cellWidth := len(canvasCell) - 1
			cellHeight := len(canvasCell[0]) - 1
			coordX, coordY := 0, 0

			if x != 0 && y != 0 {
				coordX = x + cellWidth
				coordY = y + cellHeight
			} else if x == 0 && y != 0 {
				coordY = y + cellHeight
			} else if x != 0 && y == 0 {
				coordX = x + cellWidth
			}

			e := tl.NewEntityFromCanvas(coordX, coordY, canvasCell)
			l.AddEntity(e)
		}
	}

	g.Start()
}

// x-y coordinates for the cells
type coord struct {
	x, y int
}

type walls struct {
	north bool
	east bool
	south bool
	west bool
}

type cell struct {
	visited bool
	w walls
}

func genMaze(w, h int) [][]cell {
	maze := make([][]cell, w)
	for i, _ := range maze {
		col := make([]cell, h)
		for j, _ := range col {
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

	/*
		something simple: recursive backtracker
		https://en.wikipedia.org/wiki/Maze_generation_algorithm#Depth-first_search
	*/

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
		} else if len(backtrack) > 0 { 	// 2. Else if stack is not empty
			// 1. Pop a cell from the stack i.e. get the last element from the slice
			// 2. Make it the current cell
			cur = backtrack[len(backtrack) - 1]
			backtrack = backtrack[:len(backtrack) - 1]
		}
	}

	return maze
}

// remove the wall between the current cell and the chosen neighbor
func markWall(cur, neighbor coord, maze [][]cell) {
	// NOTE: just want to see how this looks
	switch {
	// on the north/south plane
	case cur.x - neighbor.x == 0:
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
	case cur.y - neighbor.y == 0:
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
	if x, y := cur.x, cur.y - 1; y >= 0 {
		addNeighbor(x, y)
	}

	// didn't pass east border
	if x, y := cur.x + 1, cur.y; x < width {
		addNeighbor(x, y)
	}

	// didn't pass south border
	if x, y := cur.x, cur.y + 1; y < height {
		addNeighbor(x, y)
	}

	// didn't pass west border
	if x, y := cur.x - 1, cur.y; x >= 0 {
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