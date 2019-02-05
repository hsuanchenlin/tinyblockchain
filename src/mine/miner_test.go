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
	res, nonce := miner.mine()
	fmt.Println(res,nonce)
}