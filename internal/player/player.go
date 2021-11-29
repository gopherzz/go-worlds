package player

import (
	"fmt"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/gopherzz/go-worlds/internal/block"
	"github.com/gopherzz/go-worlds/internal/block/blocktypes"
	"github.com/gopherzz/go-worlds/internal/constants"
	"github.com/gopherzz/go-worlds/internal/player/inventory"
	"github.com/gopherzz/go-worlds/internal/utils"
	"github.com/gopherzz/go-worlds/internal/world"
)

type Player struct {
	screenPos, worldPos pixel.Vec
	currentWorld        *world.World
	sprite              pixel.Sprite
	inventory           *inventory.Inventory
	selected            *inventory.InventorySlot
	selectedText        *text.Text
	selectedIndex       int
	IsMove, IsPlace     chan int
	showInventory       bool
}

func NewPlayer(x, y float64, imagepath string, world *world.World, txt *text.Text) *Player {
	pic, err := utils.LoadPic(imagepath)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())

	inv := inventory.NewInventory()
	for i := range block.BLOCKS {
		inv.AddItemWithCount(i, 5)
	}

	return &Player{
		screenPos: pixel.V(x, y),
		worldPos:  pixel.V(float64(int(x)-(int(x)%constants.WORLD_WIDTH)), float64(int(y)-(int(y)%constants.WORLD_HEIGHT))),
		sprite:    *sprite,
		// Admin
		inventory:     inv,
		selected:      inv.GetSlot(0),
		currentWorld:  world,
		selectedText:  txt,
		showInventory: false,
	}
}

func (p *Player) UseController(win *pixelgl.Window) {
	posX, posY := p.worldPos.X, p.worldPos.Y
	if win.JustPressed(pixelgl.KeyW) && p.GetPos().Y+32 < constants.WINDOW_HEIGHT {
		p.Move(0, 1)
	}
	if win.JustPressed(pixelgl.KeyS) && p.GetPos().Y-32 > 0 {
		p.Move(0, -1)
	}
	if win.JustPressed(pixelgl.KeyA) && p.GetPos().X-32 > 0 {
		p.Move(-1, 0)
	}
	if win.JustPressed(pixelgl.KeyD) && p.GetPos().X+32 < constants.WINDOW_WIDTH {
		p.Move(1, 0)
	}
	if win.JustPressed(pixelgl.KeyH) && posX-1 >= 0 {
		p.placeBlock(pixel.V(posX-1, posY))
	}
	if win.JustPressed(pixelgl.KeyJ) && posY-1 >= 0 {

		p.placeBlock(pixel.V(posX, posY-1))
	}
	if win.JustPressed(pixelgl.KeyK) && posY+1 <= constants.WORLD_HEIGHT {
		p.placeBlock(pixel.V(posX, posY+1))
	}
	if win.JustPressed(pixelgl.KeyL) && posX+1 <= constants.WORLD_WIDTH {
		p.placeBlock(pixel.V(posX+1, posY))
	}
	if win.JustPressed(pixelgl.KeyTab) {
		p.changeBlock()
	}
	if win.JustPressed(pixelgl.KeyI) {
		p.showInventory = !p.showInventory
	}
}

func (p *Player) changeBlock() {
	if p.selectedIndex+1 == p.inventory.Length() {
		p.selectedText.Clear()
		p.selectedIndex = 0
		fmt.Fprint(p.selectedText, p.inventory.GetSlot(p.selectedIndex).Block.Name)
		p.selected = p.inventory.GetSlot(p.selectedIndex)
		return
	}
	p.selectedText.Clear()
	p.selectedIndex++
	fmt.Fprint(p.selectedText, p.inventory.GetSlot(p.selectedIndex).Block.Name)
	p.selected = p.inventory.GetSlot(p.selectedIndex)
}

func (p *Player) placeBlock(pos pixel.Vec) {
	if p.selected.Count <= 0 {
		return
	}
	fmt.Println(p.selected)
	p.selected.Count--
	cpy := p.selected.Block
	p.currentWorld.WorldMap[pos] = &cpy
}

func (p *Player) Move(moveX, moveY float64) {
	moveWorld := p.worldPos.Add(pixel.V(moveX, moveY))
	if p.currentWorld.WorldMap[moveWorld].BlockType == blocktypes.SOLID {
		canMove := p.punchBlock(moveWorld)
		if !canMove {
			return
		}
	}
	moveScreen := p.screenPos.Add(pixel.V(moveX*constants.SPRITE_HEIGHT, moveY*constants.SPRITE_HEIGHT))
	p.screenPos = moveScreen
	p.worldPos = moveWorld
}

func (p *Player) punchBlock(blockPos pixel.Vec) bool {
	p.currentWorld.WorldMap[blockPos].Punch()
	if p.currentWorld.WorldMap[blockPos].Strength <= 0 {
		p.inventory.AddItem(p.currentWorld.WorldMap[blockPos].BlockId)
		return true
	}
	return false
}

func (p *Player) Draw(win *pixelgl.Window) {
	p.sprite.Draw(win, pixel.IM.Moved(p.screenPos))
	if p.showInventory {
		p.inventory.Draw(win)
	}
}

func (p *Player) GetPos() pixel.Vec {
	return p.screenPos
}
