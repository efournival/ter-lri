package main

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	// Max genus to compute
	MAX_GENUS = 42

	// Inverse depth from which Cilk will be used
	CILK_BOUND = 20

	// When this number of tasks is reached, adding a task will be blocking
	MAX_TASKS = 10000000
)

type Danser struct {
	isMaster bool
	ready    bool
	workers  []*Worker
	server   *Server
	result   nm.MonoidResults
	tasks    chan nm.GoMonoid
	results  chan nm.MonoidResults
	syncc    chan net.Conn
	finished chan bool
}

func NewDanser(master bool) (d *Danser) {
	d = &Danser{}

	d.isMaster = master
	d.ready = false
	d.tasks = make(chan nm.GoMonoid, MAX_TASKS)
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
		Port     string   `json:"port"`
		Machines []string `json:"machines"`
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Fatalln("ERROR: unmarshaling JSON failed:", err.Error())
	}

	d.server = NewServer(":"+config.Port, d.tasks, d.syncc)
	log.Println("Server is now listening on port", config.Port)

	for _, address := range config.Machines {
		d.workers = append(d.workers, NewWorker(address, d.tasks, d.results))
	}

	log.Println("Loaded config from", filename, "successfully")

	if len(d.workers) == 0 {
		log.Println("WARNING: no worker has been found")
		d.ready = true
	} else {
		go d.waitForWorkers()
	}

	return nil
}

func (d *Danser) waitForWorkers() {
	for {
		ok := true

		for _, worker := range d.workers {
			ok = ok && (worker.connection != nil)
		}

		if ok {
			d.ready = true
			log.Println("All workers are ready")
			return
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (d *Danser) Danse() {
	if d.server == nil {
		panic("DANSE configuration has not been loaded")
	}

	go func() {
		if err := d.server.Listen(); err != nil {
			panic(err)
		}
	}()

	go d.schedule()

	if d.isMaster {
		log.Println("Starting DANSE as the master process")
		d.work(nm.NewMonoid())
	} else {
		log.Println("Starting DANSE as worker")
	}

	// TODO: option (verbose)
	/*go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			log.Println(len(d.tasks), "tasks queued")
		}
	}()*/

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
			d.tasks <- m.RemoveGenerator(it.GetGen())
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

func (d *Danser) schedule() {
	go func() {
		for {
			select {
			case conn := <-d.syncc:
				sync(conn, d.result)
			case result := <-d.results:
				// Reduce
				for k, v := range result {
					d.result[k] += v
				}
			default:
				if d.isMaster && len(d.tasks) == 0 {
					finished := true

					for _, worker := range d.workers {
						if time.Now().Sub(worker.lastSync) < 2*time.Second {
							finished = false
							break
						}
					}

					// In case there is no worker, check if d.finished has not already been set to true
					if finished && len(d.finished) == 0 {
						d.finished <- true
					}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case task := <-d.tasks:
				d.work(task)
			}
		}
	}()

	if d.isMaster {
		go func() {
			for {
				time.Sleep(time.Second)

				if d.ready {
					log.Println("Asking all the workers to sync")

					for _, worker := range d.workers {
						worker.Sync()
					}
				}
			}
		}()
	}

	if len(d.workers) > 0 && !d.isMaster {
		go func() {
			for {
				// TODO: 200 from configuration
				time.Sleep(400 * time.Millisecond)

				if d.ready && len(d.tasks) == 0 {
					// TODO: 50 from configuration
					w := d.workers[rand.Intn(len(d.workers))]
					w.Steal(5)
					log.Println("Sending steal request to", w.connection.RemoteAddr())
				}
			}
		}()
	}
}

func sync(conn net.Conn, result nm.MonoidResults) {
	empty := true

	for i := 0; i < len(result); i++ {
		if result[i] != 0 {
			empty = false
			break
		}
	}

	if empty {
		return
	}

	err := binary.Write(conn, binary.BigEndian, NewSyncAnswer(result))

	if err != nil {
		log.Panicln("WORKER Binary write to", conn.LocalAddr(), "failed:", err.Error())
	}

	log.Println("Sent sync information")

	for i := 0; i < len(result); i++ {
		result[i] = 0
	}
}
