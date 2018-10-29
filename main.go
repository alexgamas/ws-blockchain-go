package main

import (
	"fmt"
	"math/rand"
	"time"

	bc "github.com/alexgamas/ws-blockchain-go/blockchain"
)

var chain bc.Chain

func main() {

	bloco := bc.NewBlock()

	chain = bc.Chain{Prev: nil, Next: nil, Block: bloco}

	s := false

	td0 := time.Now()
	t0 := time.Now()

	perf := 0

	var hash string

	for !s {

		t1 := time.Now()

		bloco.Nonce = rand.Int63()

		hash = bloco.GenerateHash()

		// s = blockchain.ValidateProof(hash)

		if bc.ValidateProof(hash) {
			fmt.Printf("%s, total: %s\n", hash, time.Now().Sub(td0))
			bloco.Hash = hash

			bloco = bc.NewBlockFromBlock(bloco)

			td0 = time.Now()

		}

		if t1.Sub(t0).Nanoseconds() >= int64(time.Second) {
			fmt.Printf("%v H/s\n", perf)
			t0 = time.Now()
			perf = 0
		}

		perf++

	}

	// fmt.Printf("%s, total: %s\n", hash, time.Now().Sub(td0))
}
