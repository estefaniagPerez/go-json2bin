package main

import (
	"localhost/estefaniagPerez/json2bin"
	"log"
	"os"
)

func main() {

	// 1. Read Json File
	if len(os.Args) < 1 {
		log.Fatal("Not enough arguments")
	}

	// --gen_bin -> generate resource / --parse_file -> generate json
	typeOp := os.Args[1]
	jsonPath := os.Args[2]
	binaryOut := os.Args[3]

	switch typeOp {
	case "--gen_bin":
		err := json2bin.GenBin(jsonPath, binaryOut)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Success Generating Binary")
	case "--parse_file":
		err := json2bin.ParseConf(jsonPath, binaryOut)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Success Parsing to JSON")
	default:
		log.Fatal("Unknown Option")
	}

}
