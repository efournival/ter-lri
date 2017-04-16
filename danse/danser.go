package main

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	// Max genus to compute
	MAX_GENUS = 35

	// Inverse depth from which Cilk will be used
	CILK_BOUND = 14

	// When this number of tasks is reached, adding a task will be blocking
	MAX_TASKS = 1000
)

type Danser struct {
	workers  []*Worker
	server   *Server
	result   nm.MonoidResults
	tasks    chan nm.MonoidStorage
	results  chan nm.MonoidResults
	syncc    chan net.Conn
	finished chan bool
}

func NewDanser() (d *Danser) {
	d = &Danser{}

	d.tasks = make(chan nm.MonoidStorage, MAX_TASKS)
	d.results = make(chan nm.MonoidResults, MAX_TASKS)
	d.syncc = make(chan net.Conn, 1)
	d.finished = make(chan bool, 1)

	// Fail silently if the file does not exist
	d.LoadConfig("config.json")

	return
}

func (d *Danser) LoadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	var config struct {
		port      string
		addresses []string
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		return err
	}

	d.server = NewServer(":"+config.port, d.tasks, d.syncc)

	for _, address := range config.addresses {
		worker, err := NewWorker(address, d.tasks, d.results)

		if err != nil {
			return err
		}

		d.workers = append(d.workers, worker)
	}

	log.Println("Loaded config from", filename, "successfully")

	return nil
}

func (d *Danser) Danse(isMaster bool) {
	if d.server == nil {
		panic("DANSE configuration has not been loaded")
	}

	go func() {
		if err := d.server.Listen(); err != nil {
			panic(err)
		}
	}()

	d.schedule()

	if isMaster {
		log.Println("Starting DANSE as the master process")
		d.work(nm.NewMonoid())
	} else {
		log.Println("Starting DANSE as worker")
	}

	// Wait indefinitely if worker, or until computation is finished if master
	if <-d.finished {
		log.Println(d.result)
		log.Println("DANSE finished")
	}
}

func (d *Danser) work(m nm.GoMonoid) {
	if m.Genus() < MAX_GENUS-CILK_BOUND {
		it := m.NewIterator()
		var nbr uint64

		for it.MoveNext() {
			// TODO: m.RemoveIteratorGenerator(it)
			d.work(m.RemoveGenerator(it.GetGen()))
			nbr++
		}

		var res nm.MonoidResults
		res[m.Genus()] += nbr
		d.results <- res

		it.Free()
	} else {
		d.results <- m.WalkChildren()
	}
}

func sync(conn net.Conn, result nm.MonoidResults) {
	err := binary.Write(conn, binary.BigEndian, NewSyncAnswer(result))

	if err != nil {
		log.Println("Binary write to", conn.LocalAddr(), "failed:", err.Error())
	}

	for i := 0; i < len(result); i++ {
		result[i] = 0
	}
}

func (d *Danser) schedule() {
	go func() {
		for {
			select {
			case conn := <-d.syncc:
				// Requested to sync
				sync(conn, d.result)
			case result := <-d.results:
				// Reduce
				for k, v := range result {
					d.result[k] += v
				}
			}
		}
	}()

	go func() {
		for {
			time.Sleep(250 * time.Millisecond)

			if len(d.tasks) == 0 {
				// No tasks, no results, we are done
				if len(d.results) == 0 {
					// TODO: not if worker
					d.finished <- true
				} /*else {
					d.workers[rand.Intn(len(d.workers))].Steal(2)
				}*/
			}
		}
	}()
}
