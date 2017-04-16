package main

import "flag"

func main() {
	danser := NewDanser()

	var master bool
	flag.BoolVar(&master, "master", false, "One to rule them all")
	flag.Parse()

	danser.Danse(master)
}
