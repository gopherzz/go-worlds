package inventory

import (
	"fmt"

	"github.com/gopherzz/go-worlds/internal/block"
)

type InventorySlot struct {
	Count int
	Block block.Block
}

func NewSlot(b block.Block) *InventorySlot {
	return &InventorySlot{
		Count: 1,
		Block: b,
	}
}

func NewSlotWithCount(b block.Block, count int) *InventorySlot {
	return &InventorySlot{
		Count: count,
		Block: b,
	}
}

func (s InventorySlot) String() string {
	return fmt.Sprintf("%s: %d", s.Block.Name, s.Count)
}
