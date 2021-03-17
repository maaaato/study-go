package main

import (
	"flag"
)

func main() {
	flag.IntVar(&port, "port", envPort, "port to use")
	func IntVar(p *int, name string, value int, usage string){
		CommandLine.Var(newIntValue(value, p), name, usage)
	}
	flags := flag.NewFlagSet("example", flag.ContinueOnError)
}
