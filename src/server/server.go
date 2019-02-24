package server

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"os"
	"amis_test/src/core/types"
	"math/rand"
)

type Args struct {
	From string
	To string
	Value int
}

type Acc struct  {
	Accs map[string]int
	Nonce int
}

type Dealer struct {

}

func (d *Dealer) AcceptTX(args *Args,  reply *string) error {

	*reply = "Yi ha"
	valid := true
	acc, err := getAcc()
	if err != nil {
		log.Fatal(err)
	}

	accs := acc.Accs
	fmt.Printf("from %s v %d\n", args.From, accs[args.From])
	if _, ok := accs[args.To]; !ok {
		//do something here
		accs[args.To] = 100
	}

	if val, ok := accs[args.From]; ok {
		if args.Value > val {
			*reply = "not enough"
			valid = false
		} else {
			accs[args.From] -= args.Value
			*reply = "commit success"
			valid = true
		}
	} else {
		if args.Value > 100 {
			*reply = "not enough"
			valid = false
		} else {
			accs[args.From] = 100 - args.Value
			accs[args.To] += args.Value
			*reply = "valid"
			valid = true
		}
	}
	fmt.Printf("after from %s v %d\n", args.From, accs[args.From])
	if valid {
		fmt.Println(valid)
		acc.Nonce++
		writeAcc(acc)
		//random hash

		seed := make([]byte, 40)
		seed = append(seed, byte(rand.Int()))
		hash := crypto.Keccak256(seed)
		var txh [32]byte
		copy(txh[:], hash)
		tx := types.Transaction{From:args.From, To:args.To, Value:args.Value, Hash:txh}
		writeTX(tx)
	}
	return nil
}

func getAcc() (Acc, error) {
	file, err := os.Open("state/state.json")
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(file)

	var kv Acc
	err = dec.Decode(&kv)
	if err != nil {
		log.Fatal(err)
	}
	return kv, nil
}

func writeAcc(acc Acc) {
	file, err := os.Create("state/state.json")
	if err != nil {
		fmt.Println("write  err")
		log.Fatal(err)
	}
	enc := json.NewEncoder(file)
	enc.Encode(acc)
	file.Close()
}

func writeTX(tx types.Transaction) {
	hash := hex.EncodeToString(tx.Hash[:])
	file, err := os.Create("txs/"+hash+".json")
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(file)
	enc.Encode(tx)
	file.Close()
}



