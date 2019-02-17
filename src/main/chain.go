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

func main() {

	var accs []types.Account
	blockchain := &types.BlockChain{nil, &accs}

	tmp, err := hex.DecodeString("deadbeefdeadbeefdeadbeefdeadbeef")
	if err != nil {
		log.Fatal(err)
	}
	var gh [32]byte
	copy(gh[:], tmp)
	genesisHeader := types.Header{Transactions: nil, BlockHash: gh,  ParentHash:gh, BlockHeight:0, Nonce:0}
	genesisBlk := types.Block{  &genesisHeader}
	blockchain.Last = &genesisBlk

	for {
		t := time.NewTimer(10*time.Second)

		txs := wrapping("../../txs/", )
		txRootHash := getTXRootHash(txs, blockchain.Last.Header.BlockHash)

		var bh [32]byte
		copy(bh[:], txRootHash)

		miner := mine.Miner{TargetPrefixZeroes:0, BlockHash:bh, ThreadNumber:4, Nonce:0}
		ch := make(chan int)
		go miner.Start(ch, bh[:])
		header := types.Header{nil, bh, blockchain.Last.Header.BlockHash, 0, 0}
		blk := types.Block{  &header}
		blk.Header.Nonce = <-ch
		blockchain.Last = &blk
		expire := <- t.C
		fmt.Printf("Expiration time: %v.\n", expire)
	}
}

type TXS struct {
	TXS []Transaction `json:"txs"`
}

type Transaction struct {
	Hash   string `json:"hash"`
	From   string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

func wrapping(path string, fileName string)  []Transaction{
	jsonFile, err := os.Open("../../txs/1.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var txs TXS

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &txs)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(txs.TXS); i++ {
		//txhash := hex.EncodeToString(txs.TXS[i].Hash[:])
		fmt.Printf("hash: %s\n", txs.TXS[i].Hash)
		fmt.Printf("from: %s\n", txs.TXS[i].From)
		fmt.Printf("to: %s\n", txs.TXS[i].To)
		fmt.Printf("value %d\n", txs.TXS[i].Value)
	}

	err = os.Rename(fmt.Sprintf("%s/%s", path, fileName), fmt.Sprintf("../../history/%s", fileName))
	if err != nil {
		panic(err)
	}

	return txs.TXS
}

func moveTX2History() {
	os.RemoveAll("/tmp/")
}

func getTXRootHash(txs []Transaction, prevHash [32]byte) []byte{
	seed := make([]byte, 40)
	seed = crypto.Keccak512(seed)
	seed = append(seed, prevHash[:]...)
	for _, tx := range(txs) {
		decoded, err := hex.DecodeString(tx.Hash)
		if err != nil {
			log.Fatal(err)
		}
		seed = append(seed, decoded...)
	}
	return crypto.Keccak256(seed)
}

func applyBalanceChange() {

}