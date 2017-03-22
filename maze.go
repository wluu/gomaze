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
	if *width == 0 {
		log.Fatal("width cannot be 0.")
	}
	if *height == 0 {
		log.Fatal("height cannot be 0.")
	}

	g := tl.NewGame()

	l := tl.NewBaseLevel(tl.Cell{})
	g.Screen().SetLevel(l)

	gen.MazeFile(*width, *height)

	dat, err := ioutil.ReadFile(gen.MAZE_FILE)
	if err != nil {
		log.Fatal(err)
	}

	e := tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(dat)))
	l.AddEntity(e)
	g.Start()
}
