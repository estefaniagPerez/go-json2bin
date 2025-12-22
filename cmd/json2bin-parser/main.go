package main

import (
	"flag"
	"log"
	"os"

	"github.com/estefaniagPerez/json2bin"
)

type Color string

const (
	ColorBlack  = "\u001b[30m"
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
)

type Options string

const (
	OptionGenBin    = "gen_bin"
	OptionParseFile = "parse_file"
)

func main() {

	// 1. Read Json File
	if len(os.Args) < 1 {
		log.Fatal("Not enough arguments")
	}

	var typeOp string
	var jsonPath string
	var binaryOut string
	flag.StringVar(&typeOp, "op", OptionGenBin, "Type of operation: gen_bin or parse_file")
	flag.StringVar(&jsonPath, "json", "./assets/Conf.json", "Path to the JSON file")
	flag.StringVar(&binaryOut, "bin", "./assets/Conf.dat", "Path to the binary output file")
	flag.Parse()

	switch typeOp {
	case OptionGenBin:
		err := json2bin.GenBin(jsonPath, binaryOut)
		if err != nil {
			log.Fatal(string(ColorRed), "✗ ", err, string(ColorReset))
		}
		log.Println(string(ColorGreen), "✓ Success Generating Binary", string(ColorReset))
	case OptionParseFile:
		err := json2bin.ParseConf(jsonPath, binaryOut)
		if err != nil {
			log.Fatal(string(ColorRed), "✗ ", err, string(ColorReset))
		}
		log.Println(string(ColorGreen), "✓ Success Parsing to JSON", string(ColorReset))
	default:
		log.Println(string(ColorRed), "✗ Unknown Option", string(ColorReset))
	}
}
