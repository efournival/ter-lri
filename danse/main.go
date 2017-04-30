package main

import "flag"

var debug = true

func main() {
	var master bool
	flag.BoolVar(&master, "master", false, "One to rule them all")
	flag.Parse()

	danser := NewDanser(master)
	danser.Danse()
}
