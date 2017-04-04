package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

// When this number of tasks is reached, adding a task will be blocking
const MAX_TASKS = 15

type (
	Parameter struct {
		Min, Max, Value int
	}

	Danser struct {
		workerFunc func(nm.GoMonoid) nm.MonoidResults
		masterFunc func(chan bool)
		parameters map[string]Parameter
		tasks      chan Task
		workers    []*Worker
		server     *Server
	}
)

func NewDanser() (d *Danser) {
	d = &Danser{}
	d.parameters = make(map[string]Parameter)
	d.tasks = make(chan Task, MAX_TASKS)

	d.workerFunc = func(gm nm.GoMonoid) nm.MonoidResults {
		panic("Worker function is not defined")
		return nm.MonoidResults{}
	}

	d.masterFunc = func(cb chan bool) {
		panic("Master function is not defined")
	}

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

	d.server = NewServer(":"+config.port, d.tasks)

	for _, address := range config.addresses {
		worker, err := NewWorker(address, d.tasks)

		if err != nil {
			return err
		}

		d.workers = append(d.workers, worker)
	}

	log.Println("Loaded config from", filename, "successfully")

	return nil
}

func (d *Danser) RegisterParameter(name string, min, max, value int) {
	d.parameters[name] = Parameter{min, max, value}
}

func (d *Danser) Parameter(name string) int {
	return d.parameters[name].Value
}

func (d *Danser) WorkerFunc(wf func(nm.GoMonoid) nm.MonoidResults) {
	d.workerFunc = wf
}

func (d *Danser) MasterFunc(mf func()) {
	d.masterFunc = func(finished chan bool) {
		mf()
		finished <- true
	}
}

func (d *Danser) Danse(isMaster bool) {
	finishedMaster := make(chan bool, 1)

	if d.server == nil {
		panic("DANSE configuration has not been loaded")
	}

	go func() {
		if err := d.server.Listen(); err != nil {
			panic(err)
		}
	}()

	if isMaster {
		log.Println("Starting DANSE as the master process")
		d.masterFunc(finishedMaster)
	} else {
		log.Println("Starting DANSE as worker")
	}

	// Wait indefinitely if worker, or until computation is finished if master
	if <-finishedMaster {
		log.Println("DANSE finished")
	}
}

func (d *Danser) Work(gm nm.GoMonoid) nm.MonoidResults {
	// TODO: schedule
	return d.workerFunc(gm)

	//return <-d.queue(gm)
}

func (d *Danser) schedule() {
	go func() {
		for {
			if len(d.tasks) == 0 {
				// TODO: parameter
				d.workers[rand.Intn(len(d.workers))].Steal(2)
			}

			// TODO: parameter
			time.Sleep(500 * time.Millisecond)
		}
	}()
}
