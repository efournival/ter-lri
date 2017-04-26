package main

import "flag"

func main() {
	var master bool
	flag.BoolVar(&master, "master", false, "One to rule them all")
	flag.Parse()

	danser := NewDanser(master)
	danser.Danse()
}
