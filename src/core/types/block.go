package types

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"time"
)

type Hash [32]byte
type Address [32]byte

type Header struct {
	Transactions []*Transaction
	BlockHash Hash
	ParentHash Hash
	BlockHeight uint64
	Nonce int
}

type Block struct {
	Header *Header
}

type HeaderLog struct {
	Transactions []*Transaction
	BlockHash string
	ParentHash string
	BlockHeight uint64
	Nonce int
}

type BlockLog struct {
	HeaderLog *HeaderLog
}

func (blk *Block) addTX(tx *Transaction) (timestamp time.Time, err error) {
	blk.Header.Transactions = append(blk.Header.Transactions, tx)
	return time.Now(), nil
}

func (blk *Block) WriteFile(filePath string) error {
	hashString := hex.EncodeToString(blk.Header.BlockHash[:])
	file, err := os.Create(filePath+"/" + hashString)
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(file)
	headerlog := HeaderLog{
		Transactions:blk.Header.Transactions,
		BlockHash: hashString,
		ParentHash: hex.EncodeToString(blk.Header.ParentHash[:]),
		BlockHeight:blk.Header.BlockHeight,
		Nonce:blk.Header.Nonce,
	}
	blklog := BlockLog{HeaderLog:&headerlog}
	enc.Encode(blklog)
	file.Close()
	return nil
}