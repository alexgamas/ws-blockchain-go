package blockchain

import (
	"fmt"
)

//Chain ..
type Chain struct {
	Block Block
	Next  *Chain
	Prev  *Chain
}

func (c Chain) String() string {
	return fmt.Sprintf("[%d]%s", c.Block.Index, c.Block.Hash)
}
