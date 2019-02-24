package server

import (
	"amis_test/src/core/types"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	if _, ok := accs[args.To]; !ok {
		//do something here
		accs[args.To] = 100
	}

	if val, ok := accs[args.From]; ok {
		//do something here
		if args.Value > val {
			*reply = "not enough"
			valid = false
		} else {
			*reply = "commit success"
			valid = true
		}
	} else {
		if args.Value > 100 {
			*reply = "not enough"
			valid = false
		}
	}
	if valid {
		fmt.Println(valid)
		acc.Nonce++
		writeAcc(acc)


		tx := types.Transaction{From:args}
		addTX()
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
		fmt.Println("writeerr")
		log.Fatal(err)
	}
	enc := json.NewEncoder(file)
	enc.Encode(acc)
	file.Close()
}

func addTX(tx types.Transaction) {
	file, err := os.Create("txs/state.json")
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(file)
	enc.Encode(tx)
	file.Close()
}



