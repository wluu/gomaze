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

// Player is an interface method from JoelOtter/termloop
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

// Draw is an interface method from JoelOtter/termloop
func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/150-x, screenHeight/100-y)
	player.Entity.Draw(screen)
}

// Tick is an interface method from JoelOtter/termloop
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func main() {
	flag.Parse()
	if *width == 0 {
		log.Fatal("width cannot be 0.")
	}
	if *height == 0 {
		log.Fatal("height cannot be 0.")
	}

	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{})

	player := Player{
		Entity: tl.NewEntity(0, 0, 1, 1),
		level:  level,
	}
	level.AddEntity(&player)
	game.Screen().SetLevel(level)

	gen.MazeFile(*width, *height)
	dat, err := ioutil.ReadFile(gen.FileMaze)
	if err != nil {
		log.Fatal(err)
	}
	maze := tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(dat)))
	level.AddEntity(maze)

	game.Start()
}
