package main

import (
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

type Worker struct {
	RPC        *rpc.Client
	Address    string
	TaskStream chan nm.GoMonoid
	Results    chan nm.MonoidResults
	lastSync   time.Time
}

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

func (w *Worker) Steal() {
	var reply StealReply

	if err := w.RPC.Call("Danse.StealRequest", STEAL_COUNT, &reply); err != nil {
		panic(err)
	}

	log.Println("Stole", reply.Count, "tasks")

	for i := 0; i < reply.Count; i++ {
		w.TaskStream <- nm.FromBytes(reply.Tasks[i])
	}
}

func (w *Worker) Sync() {
	var result nm.MonoidResults

	if err := w.RPC.Call("Danse.SyncRequest", true, &result); err != nil {
		panic(err)
	}

	w.lastSync = time.Now()
	w.Results <- result
}

func (w *Worker) IsActive() bool {
	// If last sync is younger than 5 seconds, then we still have results to sync
	return time.Now().Sub(w.lastSync) < 5*time.Second
}
