package server

type Args struct {
	//X, Y int
	From string
	To string
	Value int
}

type Calculator struct{}

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}


type Dealer struct {}

func (d *Dealer) AcceptTX(args *Args,  reply *int) error {

}

