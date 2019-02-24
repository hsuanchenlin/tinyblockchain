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
	TargetPrefixZeroes uint8
	PreviousBlockHash  types.Hash
	ThreadNumber       int
	Nonce              int
}

func (miner *Miner) mine(c chan int, thNum int, nonce int) (bool, int) {
	fmt.Printf("thread %d\n", thNum)
	for {
		if miner.Nonce != 0 {
			break
		}
		nonce += thNum
		seed := make([]byte, 40)
		seed = append(seed, byte(nonce))
		seed = append(seed, miner.PreviousBlockHash[:]...)
		hashRes := crypto.Keccak256(seed)

		if miningHashCmp(miner.TargetPrefixZeroes, hashRes) {
			miner.Nonce = nonce
			//s := fmt.Sprintf("%s", hex.EncodeToString(hashRes))
			fmt.Println("success routine")
			//fmt.Println(s)
			//fmt.Printf("%d %d\n",s[7], s[8])
			//fmt.Println("------")
			c <- nonce
			return true, nonce
		}
	}
	fmt.Printf("fail routine %d\n", thNum)
	return false, 0
}

func (miner *Miner) Start(ch chan int, rootHash []byte) {
	var i int
	c := make(chan int)

	for i = 0; i < miner.ThreadNumber; i++{
		go miner.mine(c, i, i*1000000000)
	}
	timeEla := 0
	for {
		select {
			case n := <-c:
				fmt.Printf("done nonce:%d\n",n)
				ch <- n
				return
			case <-time.After(time.Second * 1):
				timeEla++
				fmt.Printf("timeout %d\n", timeEla)
		}
	}

}

func miningHashCmp(targetPrefixZeroes uint8, miningHash []byte) bool{
	var i uint8
	shash := hex.EncodeToString(miningHash)

	for i=0; i < targetPrefixZeroes; i++ {
		//rune 48 - 57 => 0 -> 9
		if  shash[i] != 48 {
			return false
		}

	}
	fmt.Printf("hash %s\n", shash)
	return true
}

//
//func hex2int(ch rune) int {
//	chInt := int(ch)
//	if chInt >= 48 && chInt <= 57 {
//		return chInt - 48
//	}
//	if chInt >= 65 && chInt <= 70 {
//		return chInt - 55
//	}
//	if chInt >= 97 && chInt <= 102 {
//		return chInt - 87
//	}
//	return -1
//}