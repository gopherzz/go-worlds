package inventory

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/gopherzz/go-worlds/internal/block"
	"github.com/gopherzz/go-worlds/internal/constants"
	"github.com/gopherzz/go-worlds/internal/properties"
)

type Inventory struct {
	slots      map[int]*InventorySlot
	slotsKeys  []int
	textWriter *text.Text
}

func NewInventory() *Inventory {
	text := text.New(pixel.V(float64(constants.WINDOW_WIDTH/2), float64(constants.WINDOW_HEIGHT)), properties.BasicAtlas)
	return &Inventory{
		slots:      make(map[int]*InventorySlot),
		textWriter: text,
		slotsKeys:  make([]int, 0),
	}
}

func (i *Inventory) AddItem(id int) {
	if _, ok := i.slots[id]; !ok {
		block := *block.BLOCKS[id]
		i.slots[id] = NewSlot(block)
		i.slotsKeys = append(i.slotsKeys, block.BlockId)
		return
	}
	i.slots[id].Count++
}

func (i *Inventory) AddItemWithCount(id, count int) {
	if _, ok := i.slots[id]; !ok {
		block := *block.BLOCKS[id]
		i.slots[id] = NewSlotWithCount(block, count)
		i.slotsKeys = append(i.slotsKeys, block.BlockId)
		return
	}
	i.slots[id].Count += count
}

func (i *Inventory) GetSlot(idx int) *InventorySlot {
	return i.slots[idx]
}

func (i *Inventory) Length() int {
	return len(i.slots)
}

func (i *Inventory) Draw(win *pixelgl.Window) {
	i.textWriter.Clear()
	for _, id := range i.slotsKeys {
		fmt.Fprintln(i.textWriter, i.slots[id].String())
	}
	i.textWriter.Draw(win, pixel.IM.Moved(pixel.V(-((i.textWriter.Bounds().Max.X-i.textWriter.Bounds().Min.X)/2), 0)))
}
