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
	pos        int
}

func NewInventory() *Inventory {
	text := text.New(pixel.V(float64(constants.WINDOW_WIDTH/2), float64(constants.WINDOW_HEIGHT)), properties.BasicAtlas)
	return &Inventory{
		slots:      make(map[int]*InventorySlot),
		textWriter: text,
		slotsKeys:  make([]int, 0),
		pos:        0,
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
	i.refreshSlotsKeys()
}

func (i *Inventory) AddItemWithCount(id, count int) {
	if _, ok := i.slots[id]; !ok {
		block := *block.BLOCKS[id]
		i.slots[id] = NewSlotWithCount(block, count)
		i.slotsKeys = append(i.slotsKeys, block.BlockId)
		return
	}
	i.slots[id].Count += count
	i.refreshSlotsKeys()
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
		//fmt.Println(i.slots[id].String())
	}
	i.textWriter.Draw(win, pixel.IM.Moved(pixel.V(-((i.textWriter.Bounds().Max.X-i.textWriter.Bounds().Min.X)/2), -16)))
}

func (i *Inventory) refreshSlotsKeys() {
	i.slotsKeys = []int{}
	for k, _ := range i.slots {
		i.slotsKeys = append(i.slotsKeys, k)
	}
	fmt.Println(i.pos)
}

func (i *Inventory) ClearSlot(id int) {
	delete(i.slots, id)
	i.removeSlotKey(id)
}

func (i *Inventory) removeSlotKey(id int) {
	keyIdx := 0
	for idx, el := range i.slotsKeys {
		if el == id {
			keyIdx = idx
		}
	}
	i.slotsKeys = append(i.slotsKeys[:keyIdx], i.slotsKeys[keyIdx+1:]...)
	i.refreshSlotsKeys()
	i.pos = 0
}

func (i *Inventory) Next() *InventorySlot {
	fmt.Println("Pos: ", i.pos)
	if i.pos+1 == i.Length() || i.pos == i.Length() {
		fmt.Println(i.slotsKeys)
		i.refreshSlotsKeys()
		i.pos = 0
		//return i.slots[i.pos]
	}
	if i.Length() == 0 {
		fmt.Println("len is 0")
		return nil
	}
	fmt.Println("Pos: ", i.pos)
	fmt.Println("Inv", i.slots)
	b := i.slots[i.slotsKeys[i.pos]]
	i.pos++
	if b == nil && i.Length() != 0 {
		fmt.Println("refreshed")
		i.refreshSlotsKeys()
	}
	return b
}
