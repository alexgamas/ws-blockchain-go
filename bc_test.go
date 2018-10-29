package main

import (
	"fmt"
	"testing"
	"ws-blockchain-go/blockchain"
)

func TestValidateProof(t *testing.T) {

	hs := map[string]bool{
		"9a9e31f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": false,
		"0a9e31f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": false,
		"009e31f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": false,
		"000e31f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": false,
		"000031f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": true,
		"000001f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": true,
		"000000f149c594beb568d683186be966faf2f63a99c35c0a175d8ec625adb066": true,
	}

	for k, v := range hs {
		res := blockchain.ValidateProof(k)
		if res != v {
			t.Error(fmt.Sprintf("Para o hash %s esperado %v, recebido %v", k, v, res))
		}
	}

}
