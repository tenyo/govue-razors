package main

import (
	"flag"
	"log"
)

var dev bool
var port int

func init() {
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.BoolVar(&dev, "dev", false, "Run in dev mode, serving frontend from the file system")
}

func main() {
	flag.Parse()

	if dev {
		log.Print("Starting govue-razors in dev mode")
	} else {
		log.Print("Starting govue-razors")
	}

	if err := NewServer(port, dev); err != nil {
		log.Fatal(err)
	}
}
