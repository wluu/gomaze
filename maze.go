package main

import (
	"./gen"
	"flag"
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
	"log"
)

var (
	width  = flag.Int("width", 0, "width of maze")
	height = flag.Int("height", 0, "height of maze")
)

func main() {
	flag.Parse()

	g := tl.NewGame()

	l := tl.NewBaseLevel(tl.Cell{})
	g.Screen().SetLevel(l)

	gen.MazeFile(*width, *height)

	dat, err := ioutil.ReadFile(gen.MAZE_FILE)
	if err != nil {
		log.Fatal(err)
	}

	e := tl.NewEntityFromCanvas(1, 1, tl.CanvasFromString(string(dat)))
	l.AddEntity(e)
	g.Start()
}
