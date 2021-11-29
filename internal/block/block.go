package block

import (
	"encoding/json"
	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/gopherzz/go-worlds/internal/constants"
	"github.com/gopherzz/go-worlds/internal/utils"
)

var (
	BLOCKS_FRAMES map[int]pixel.Rect = make(map[int]pixel.Rect)

	SPRITESHEET *pixel.PictureData

	BLOCKS = make([]*Block, 0)
)

func init() {

	var err error
	SPRITESHEET, err = utils.LoadPic("resources/img/blocks.png")
	if err != nil {
		panic(err)
	}

	id := 0
	for x := SPRITESHEET.Bounds().Min.X; x < SPRITESHEET.Bounds().Max.X; x += constants.SPRITE_WIDTH {
		for y := SPRITESHEET.Bounds().Min.Y; y < SPRITESHEET.Bounds().Max.Y; y += constants.SPRITE_HEIGHT {
			BLOCKS_FRAMES[id] = pixel.R(x, y, x+32, y+32)
			id++
		}
	}

	blocksJson, err := ioutil.ReadFile("resources/json/blocks.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(blocksJson, &BLOCKS)
	if err != nil {
		panic(err)
	}
}

type Blocks []Block

type Block struct {
	BlockType int    `json:"type"`
	BlockId   int    `json:"id"`
	Strength  int    `json:"strength"`
	Name      string `json:"name"`
	Sprite    *pixel.Sprite
}

func (b *Block) Draw(win *pixel.Batch, pos pixel.Vec, id int) {
	b.Sprite = pixel.NewSprite(SPRITESHEET, BLOCKS_FRAMES[id])
	b.Sprite.Draw(win, pixel.IM.Moved(pos))
}

func (b *Block) Punch() {
	b.Strength--
}
