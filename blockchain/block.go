package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

//Block ..
type Block struct {
	Index        int64
	Timestamp    time.Time
	Nonce        int64
	Data         string
	Hash         string
	PreviousHash string
}

func (b Block) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v",
		b.Index, b.Timestamp.Format(time.RFC3339), b.Nonce, b.Data, b.Hash, b.PreviousHash)
}

func (b Block) GenerateHash() string {
	blockString := fmt.Sprintf("%s", b.String())
	sum := sha256.Sum256([]byte(blockString))
	return fmt.Sprintf("%x", sum)
}

func NewBlock() Block {
	return Block{Timestamp: time.Now(), Index: 0}
}

func NewBlockFromBlock(b Block) Block {
	return Block{Timestamp: time.Now(), PreviousHash: b.Hash, Index: b.Index + 1}
}
