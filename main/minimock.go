package main

import (
	"minimock/server"
	"flag"
	"runtime"
)

func main() {
	port := flag.String("port", "8888", "port number")
	procs := flag.Int("procs", runtime.NumCPU(), "GOMAXPROCS value")
	flag.Parse()

	runtime.GOMAXPROCS(*procs)
	server.Server(*port)
}
