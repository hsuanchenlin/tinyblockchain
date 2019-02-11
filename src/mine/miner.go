package mine

import (
	"amis_test/src/core/types"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"time"
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

func (miner *Miner) mine(c chan uint32, thNum int, nonce uint32) (bool, uint32) {
	fmt.Printf("thread %d\n", thNum)
	for {
		if miner.nonce != 0 {
			break
		}
		nonce += uint32(thNum)
		seed := make([]byte, 40)
		seed = crypto.Keccak512(seed)
		seed = append(seed, byte(nonce))
		blkHash :=  miner.block.Header.BlockHash[:]
		hashRes := crypto.Keccak256(append(seed,  blkHash...))

		if miningHashCmp(miner.targetPrefixZeroes, hashRes) {
			miner.nonce = nonce
			s := fmt.Sprintf("%s", hex.EncodeToString(hashRes))
			fmt.Println("success routine")
			fmt.Println(s)
			fmt.Printf("%d %d\n",s[7], s[8])
			fmt.Println("------")
			c <- nonce
			return true, nonce
		}
	}
	fmt.Printf("fail routine %d\n", thNum)
	return false, 0
}

func (self *Miner) update() {

}

func (miner *Miner) start() {
	var i uint8
	c := make(chan uint32)

	for i = 0; i < miner.threadNumber; i++{
		go miner.mine(c, int(i), uint32(int(i)*1000000000))
	}
	timeEla := 0
	for {
		select {
			case n := <-c:
				fmt.Printf("done nonce:%d\n",n)
				return
			case <-time.After(time.Second * 1):
				timeEla++
				fmt.Printf("timeout %d\n", timeEla)
		}
	}

}

func miningHashCmp(targetPrefixZeroes uint8, miningHash []byte) bool{
	var i, prev uint8
	shash := hex.EncodeToString(miningHash)
	prev = shash[0]
	for i=0; i < targetPrefixZeroes; i++ {
		if  prev != shash[i] {
			return false
		}
		prev = shash[i]
	}
	return true
}


func hex2int(ch rune) int {
	chInt := int(ch)
	if chInt >= 48 && chInt <= 57 {
		return chInt - 48
	}
	if chInt >= 65 && chInt <= 70 {
		return chInt - 55
	}
	if chInt >= 97 && chInt <= 102 {
		return chInt - 87
	}
	return -1
}