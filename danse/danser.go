package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

type (
	WorkerFunc func(nm.GoMonoid) nm.MonoidResults
	MasterFunc func(chan bool)

	parameter struct {
		min, max, value int
	}

	Danser struct {
		workerFunc WorkerFunc
		masterFunc MasterFunc
		parameters map[string]parameter
		workers    []*worker
		server     *server
	}
)

func NewDanser() (d *Danser) {
	d = &Danser{}
	d.parameters = make(map[string]parameter)

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

	d.server = NewServer(config.port)

	for _, address := range config.addresses {
		worker, err := NewWorker(address)

		if err != nil {
			return err
		}

		d.workers = append(d.workers, worker)
	}

	log.Println("Loaded config from", filename, "successfully")

	return nil
}

func (d *Danser) RegisterParameter(name string, min, max, value int) {
	d.parameters[name] = parameter{min, max, value}
}

func (d *Danser) Parameter(name string) int {
	return d.parameters[name].value
}

func (d *Danser) WorkerFunc(wf WorkerFunc) {
	d.workerFunc = wf
}

func (d *Danser) MasterFunc(mf func()) {
	d.masterFunc = func(finished chan bool) {
		mf()
		close(finished)
	}
}

func (d *Danser) Work(gm nm.GoMonoid) nm.MonoidResults {
	return d.workerFunc(gm)
}

func (d *Danser) Danse(isMaster bool) {
	finishedServe, finishedMaster := make(chan bool), make(chan bool)

	if d.server == nil {
		panic("DANSE configuration has not been loaded")
	}

	err := d.server.Listen(finishedServe)

	if err != nil {
		panic(err)
	}

	if isMaster {
		log.Println("Starting DANSE as the master process")
		d.masterFunc(finishedMaster)
	} else {
		log.Println("Starting DANSE as worker")
	}

	select {
	case <-finishedServe:
	case <-finishedMaster:
	}

	log.Println("DANSE finished")
}
