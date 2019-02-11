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
	header := types.Header{nil, bh, bh, 0}
	blk := types.Block{  &header}
	miner := Miner{0, 1, &blk, bh, 0, 0}
	ch := make(chan uint32)
	res, nonce := miner.mine(ch, 1,0)
	fmt.Println(res,nonce)
}

func TestMineMulti(*testing.T) {
	s := "abcdefghijabcdefghijabcdefghijab"
	var bh [32]byte
	copy(bh[:], s)
	header := types.Header{nil, bh, bh, 0}
	blk := types.Block{  &header}
	miner := Miner{0, 4, &blk, bh, 4, 0}
	miner.start()
}