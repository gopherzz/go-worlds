package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gopherzz/go-worlds/internal/block"
	"github.com/gopherzz/go-worlds/internal/constants"
)

type World struct {
	name     string
	WorldMap map[pixel.Vec]*block.Block
}

func NewWorld(name string) *World {
	wmap := make(map[pixel.Vec]*block.Block)
	for x := 0; x < constants.WORLD_WIDTH; x++ {
		for y := 0; y < constants.WORLD_HEIGHT; y++ {
			wmap[pixel.V(float64(x), float64(y))] = block.BLOCKS[0]
		}
	}
	return &World{
		name:     name,
		WorldMap: wmap,
	}
}

func (w *World) Draw(win *pixelgl.Window, batch *pixel.Batch) {
	batch.Clear()
	for p, b := range w.WorldMap {
		posX, posY := p.X, p.Y
		pos := pixel.V(float64(posX*constants.SPRITE_WIDTH+15), float64(posY*constants.SPRITE_HEIGHT+15))
		if b.Strength <= 0 {
			w.WorldMap[p] = block.BLOCKS[0]
		}
		b.Draw(batch, pos, b.BlockId)
	}
	batch.Draw(win)
}

func (w *World) UseController(win *pixelgl.Window) {
}
