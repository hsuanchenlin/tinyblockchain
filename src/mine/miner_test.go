package mine

import (
	"amis_test/src/core/types"
	"fmt"
	"testing"
)

func init() {

}

func TestMineSimple(*testing.T) {
	s := "abcdefghijabcdefghijabcdefghijab"
	var bh [32]byte
	copy(bh[:], s)
	header := types.Header{nil, bh, bh, 0, 0}
	blk := types.Block{  &header}
	miner := Miner{0, blk.Header.BlockHash,0, 0}
	ch := make(chan int)
	res, nonce := miner.mine(ch, 1,0)
	fmt.Println(res,nonce)
}

func TestMineMulti(*testing.T) {
	s := "abcdefghijabcdefghijabcdefghijab"
	var bh [32]byte
	copy(bh[:], s)
	header := types.Header{nil, bh, bh, 0, 0}
	blk := types.Block{  &header}
	miner := Miner{0, blk.Header.BlockHash, 4, 0}
	ch := make(chan int)
	go miner.Start(ch, bh[:])
	ans := <-ch
	fmt.Printf("ans %d", ans)
}