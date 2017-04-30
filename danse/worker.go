package main

import (
	"log"
	"net"
	"net/rpc"
	"time"

	nm "github.com/efournival/ter-lri/go-numeric-monoid"
)

// Worker is the RPC client type
type Worker struct {
	RPC        *rpc.Client
	Address    string
	TaskStream chan nm.GoMonoid
	Results    chan nm.MonoidResults
	lastSync   time.Time
}

// NewWorker will create a new worker and try to automatically connect it to the specified address
func NewWorker(address string, ts chan nm.GoMonoid, r chan nm.MonoidResults) *Worker {
	w := &Worker{nil, address, ts, r, time.Now()}

	// Loop until a connection is established
	go func() {
		for {
			conn, err := net.Dial("tcp", w.Address)

			if err == nil {
				w.RPC = rpc.NewClient(conn)
				return
			}

			time.Sleep(250 * time.Millisecond)
		}
	}()

	return w
}

// Steal has to be called when the current process is idling and we want some work
func (w *Worker) Steal() {
	var reply StealReply

	// TODO: change MaxTasksRPC
	if err := w.RPC.Call("Danse.StealRequest", MaxTasksRPC, &reply); err != nil {
		panic(err)
	}

	if reply.Count > 0 {
		log.Println("Stole", reply.Count, "tasks")

		for i := 0; i < reply.Count; i++ {
			w.TaskStream <- nm.FromBytes(reply.Tasks[i])
		}
	}
}

// Sync is called every 2 seconds in order to retrieve results
func (w *Worker) Sync() {
	var result nm.MonoidResults

	if err := w.RPC.Call("Danse.SyncRequest", true, &result); err != nil {
		panic(err)
	}

	empty := true

	for i := 0; i < len(result); i++ {
		if result[i] != 0 {
			empty = false
			break
		}
	}

	if !empty {
		w.lastSync = time.Now()

		if debug {
			log.Println("Got sync from", w.Address, result)
		}

		w.Results <- result
	}
}

// IsActive returns true when the worker synced in the last 5 seconds
func (w *Worker) IsActive() bool {
	return time.Now().Sub(w.lastSync) < 5*time.Second
}
