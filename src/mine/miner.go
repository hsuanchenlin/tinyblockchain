package mine

import (
	"amis_test/src/core/types"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type Miner struct {
	//worker *Worker
	halt               uint8
	targetPrefixZeroes uint8
	block              *types.Block
	blockHash          types.Hash
	threadNumber       uint8
	nonce              uint32
}
func (miner *Miner) mine() (bool, int) {
	nonce := 0
	for {
		nonce++
		seed := make([]byte, 40)
		seed = crypto.Keccak512(seed)
		seed = append(seed, byte(nonce))
		blkHash :=  miner.block.Header.BlockHash[:]
		hashRes := crypto.Keccak256(append(seed,  blkHash...))
		s := fmt.Sprintf("%s", hex.EncodeToString(hashRes))
		fmt.Println(s)
		fmt.Println("------")
		if miningHashCmp(miner.targetPrefixZeroes, hashRes) {
			return true, nonce
		}
	}
	return false, 0
}

func (self *Miner) update() {

}

func (miner *Miner) start() {

}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func miningHashCmp(targetPrefixZeroes uint8, miningHash []byte) bool{
	var i uint8
	shash := hex.EncodeToString(miningHash)
	for i=0; i < targetPrefixZeroes; i++ {
		if shash[i] != '0' {
			return false
		}
	}
	return true
}