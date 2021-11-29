package main

import (
	"image/color"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/gopherzz/go-worlds/internal/block"
	"github.com/gopherzz/go-worlds/internal/constants"
	gameobject "github.com/gopherzz/go-worlds/internal/gameObject"
	"github.com/gopherzz/go-worlds/internal/player"
	"github.com/gopherzz/go-worlds/internal/properties"
	"github.com/gopherzz/go-worlds/internal/world"
)

var (
	bgColor color.Color = color.RGBA{102, 57, 49, 255}
	game    *Game       = NewGame()
)

type Game struct {
	gameObjects []gameobject.GameObject
}

func NewGame() *Game {
	return &Game{
		gameObjects: []gameobject.GameObject{},
	}
}

func (g *Game) Controllers(win *pixelgl.Window) {
	for _, o := range g.gameObjects {
		o.UseController(win)
	}
}

func (g *Game) DrawObjects(win *pixelgl.Window) {
	for _, o := range g.gameObjects {
		go o.Draw(win)
	}
}

func (g *Game) AddObject(o gameobject.GameObject) {
	g.gameObjects = append(g.gameObjects, o)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Go-Worlds",
		Bounds: pixel.R(0, 0, constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	world := world.NewWorld("test")
	basicTxt := text.New(pixel.V(constants.WINDOW_WIDTH/2, 0), properties.BasicAtlas)

	player := player.NewPlayer(0+constants.SPRITE_WIDTH/2, 0+constants.SPRITE_HEIGHT/2, "resources/img/player.png", world, basicTxt)
	batch := pixel.NewBatch(&pixel.TrianglesData{}, block.SPRITESHEET)

	game.AddObject(player)

	for !win.Closed() {
		win.Update()
		win.Clear(bgColor)

		game.Controllers(win)

		world.Draw(win, batch)
		player.Draw(win)
		basicTxt.Draw(win, pixel.IM)
	}
}

func main() {
	pixelgl.Run(run)
}
