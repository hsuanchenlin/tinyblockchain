package mine

import (
	"net"
	"sync"
)

type Master struct {
	sync.Mutex

	address         string
	registerChannel chan string
	doneChannel     chan bool
	workers         []string // protected by the mutex

	// Per-task information
	jobName string   // Name of currently executing job
	files   []string // Input files
	nReduce int      // Number of reduce partitions

	shutdown chan struct{}
	l        net.Listener
	stats    []int
}
