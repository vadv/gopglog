package main

import (
	"flag"
	"log"
	"parser"
)

func main() {

	log_file := flag.String("log-file", "postgresql.log", "Path to log file.")

	if !flag.Parsed() {
		flag.Parse()
	}

	parser, err := parser.Create(*log_file)
	if err != nil {
		log.Fatalf("Can't create parser: %s\n", err)
	}
	parser.DebugPrintAll()
}
