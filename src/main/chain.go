package main

import (
	"amis_test/src/core/types"
	"amis_test/src/mine"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	InitBalance = 100
	BlockGenTime = 10 * time.Second
)


func main() {

	accs := make(map[string]int)
	blockchain := &types.BlockChain{nil, accs}

	tmp, err := hex.DecodeString("deadbeefdeadbeefdeadbeefdeadbeef")
	if err != nil {
		log.Fatal(err)
	}
	var gh [32]byte
	copy(gh[:], tmp)
	genesisHeader := types.Header{Transactions: nil, BlockHash: gh,  ParentHash:gh, BlockHeight:0, Nonce:0}
	genesisBlk := types.Block{  &genesisHeader}
	blockchain.Last = &genesisBlk

	blockHeight := 0
	for {
		t := time.NewTimer(BlockGenTime)

		files, err := ioutil.ReadDir("../../txs/")
		if err != nil {
			log.Fatal(err)
		}
		var txs []types.Transaction

		for _, f := range files {
			fmt.Println(f.Name())
			tx, errtx := wrapping("../../txs/", f.Name())
			if errtx != nil {
				log.Fatal(err)
			}
			txs = append(txs, tx)
		}
		fmt.Println("---txs---")
		for _, t := range txs {
			fmt.Printf("tx %s %s %d\n", t.From,t.To,t.Value)
		}
		txRootHash := getTXRootHash(txs, blockchain.Last.Header.BlockHash)

		var bh [32]byte
		copy(bh[:], txRootHash)

		miner := mine.Miner{TargetPrefixZeroes:1, PreviousBlockHash:bh, ThreadNumber:4, Nonce:0}
		ch := make(chan int)
		go miner.Start(ch, bh[:])
		header := types.Header{Transactions:txs, BlockHash:bh, ParentHash:blockchain.Last.Header.BlockHash, BlockHeight:blockHeight, Nonce:0}
		blk := types.Block{  &header}
		blk.Header.Nonce = <-ch
		blk.Header.BlockHash = calHashByNonce(blk.Header.Nonce, bh)
		blockchain.Last = &blk
		expire := <- t.C
		fmt.Printf("Expiration time: %v.\n", expire)
		for _, t := range txs {
			fmt.Printf("tx %s %s %d\n", t.From,t.To,t.Value)
		}

		updateBalance(accs, txs)
		blk.WriteFile("../../blks")
		blockHeight++
	}
}



func wrapping(path string, fileName string) (types.Transaction, error) {

	file, err := os.Open(path+fileName)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(file)

	var tx types.Transaction
	err = dec.Decode(&tx)
	//fmt.Printf("%s\n",tx.From)
	err = os.Rename(fmt.Sprintf("%s/%s", path, fileName), fmt.Sprintf("../../history/%s", fileName))
	if err != nil {
		panic(err)
	}

	return tx,nil
}

func getTXRootHash(txs []types.Transaction, prevHash [32]byte) []byte{
	seed := make([]byte, 40)
	seed = crypto.Keccak512(seed)
	seed = append(seed, prevHash[:]...)
	for _, tx := range txs {
		seed = append(seed, tx.Hash[:]...)
	}
	return crypto.Keccak256(seed)
}


func updateBalance(accs map[string]int, txs []types.Transaction) error {
	for _, tx := range txs {
		_, ok := accs[tx.From]
		if !ok {
			accs[tx.From] = InitBalance
		}
		_, ok = accs[tx.To]
		if !ok {
			accs[tx.From] = InitBalance
		}
		if accs[tx.From] < tx.Value {
			return fmt.Errorf("Balance Error")
		}
		accs[tx.From] -= tx.Value
		accs[tx.To] += tx.Value
	}
	return nil
}

func calHashByNonce(nonce int, blockhash types.Hash) types.Hash{

	seed := make([]byte, 40)
	seed = append(seed, byte(nonce))
	seed = append(seed, blockhash[:]...)
	hashRes := crypto.Keccak256(seed)
	var gh [32]byte
	copy(gh[:], hashRes)
	return gh
}