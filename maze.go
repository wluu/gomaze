package main

import (
	"fmt"
	_ "github.com/JoelOtter/termloop"
)

func main() {
	// g := tl.NewGame()

	// empty maze for now
	// l := tl.NewBaseLevel(tl.Cell{})
	// g.Screen().SetLevel(l)

	maze := genMaze(5, 3)
	fmt.Println(maze)

	// populate the maze
	// l.AddEntity(tl.NewRectangle(i, j, 1, 1, tl.ColorWhite))

	// g.Start()
}

// track current cell
type curcell struct {
	x, y int
}

type cell struct {
	visited bool
	nWall, sWall, eWall, wWall bool
}

func genMaze(w, h int) [][]cell {
	maze := make([][]cell, w)
	for i, _ := range maze {
		col := make([]cell, h)
		maze[i] = col
	}

	/* Recursive backtracker */

	// Make the initial cell the current cell and mark it as visited
	cur := curcell{}
	maze[cur.x][cur.y].visited = true

	/*
	2. While there are unvisited cells
		1. If the current cell has any neighbours which have not been visited
			1. Choose randomly one of the unvisited neighbours
			2. Push the current cell to the stack
			3. Remove the wall between the current cell and the chosen cell
			4. Make the chosen cell the current cell and mark it as visited
		2. Else if stack is not empty
			1. Pop a cell from the stack
			2. Make it the current cell
	*/

	return maze
}