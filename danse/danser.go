package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	nm "github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	// MaxGenus defines the maximum depth of our exploration
	MaxGenus = 42

	// CilkBound is the inverse depth from which Cilk (through bindings) will be used
	CilkBound = 30

	// MaxTasks is the size of the tasks queue, adding more tasks than that will be blocking
	MaxTasks = 10000000
)

// Danser is the core of our experiment
// In french, 'danse' means 'dance'
// It also means Distributed Array of Numerical Semigroup Explorations
type Danser struct {
	isMaster bool
	ready    bool
	workers  []*Worker
	server   *Server
	result   nm.MonoidResults
	tasks    chan nm.GoMonoid
	results  chan nm.MonoidResults
	syncc    chan chan nm.MonoidResults
	finished chan bool
}

// NewDanser creates a new Danser objet and try to load a default configuration (from ./config.json)
func NewDanser(master bool) (d *Danser) {
	d = &Danser{}

	d.isMaster = master
	d.ready = false
	d.tasks = make(chan nm.GoMonoid, MaxTasks)
	d.results = make(chan nm.MonoidResults, MaxTasks)
	d.syncc = make(chan chan nm.MonoidResults, 1)
	d.finished = make(chan bool, 1)

	// Fail silently if the file does not exist
	d.LoadConfig("config.json")

	return
}

// LoadConfig loads the configuration (no shit)
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

// Danse will start computation
func (d *Danser) Danse() {
	if d.server == nil {
		panic("DANSE configuration has not been loaded")
	}

	go d.schedule()

	if d.isMaster {
		log.Println("Starting DANSE as the master process")
		// Start exploration from the root
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
		log.Println("Results:", d.result)

		var accumulator uint64

		for _, atGenus := range d.result {
			accumulator += uint64(atGenus)
		}

		log.Println("Total:", accumulator)

		log.Println("DANSE finished")
	}
}

func (d *Danser) waitForWorkers() {
	for {
		ok := true

		// Wait for all workers to have a RPC client object defined (Dial succeeded)
		for _, worker := range d.workers {
			ok = ok && (worker.RPC != nil)
		}

		if ok {
			d.ready = true
			log.Println("All workers are ready")
			return
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (d *Danser) work(m nm.GoMonoid) {
	if m.Genus() < MaxGenus-CilkBound {
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

	m.Free()
}

func (d *Danser) schedule() {
	go func() {
		for {
			select {
			case sc := <-d.syncc:
				// Asked to sync, answer if we have non-null results
				// TODO: timeout if no result
				sync(sc, &d.result)
			case result := <-d.results:
				// Reduce
				for k, v := range result {
					d.result[k] += v
				}
			default:
				if d.isMaster && len(d.tasks) == 0 {
					finished := true

					for _, worker := range d.workers {
						if worker.IsActive() {
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
				// Explore as long as new tasks are coming into this channel
				d.work(task)
			}
		}
	}()

	if d.isMaster {
		go func() {
			for {
				time.Sleep(2 * time.Second)

				if d.ready {
					for _, worker := range d.workers {
						// Sync all the workers every 2 seconds
						go worker.Sync()
					}
				}
			}
		}()
	}

	if len(d.workers) > 0 && !d.isMaster {
		go func() {
			for {
				time.Sleep(250 * time.Millisecond)

				if d.ready && len(d.tasks) == 0 {
					// Pick one randomly
					w := d.workers[rand.Intn(len(d.workers))]
					log.Println("Sending steal request to", w.Address)
					w.Steal()
				}
			}
		}()
	}
}

func sync(sc chan nm.MonoidResults, result *nm.MonoidResults) {
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

	sc <- *result

	for i := 0; i < len(result); i++ {
		result[i] = 0
	}
}
