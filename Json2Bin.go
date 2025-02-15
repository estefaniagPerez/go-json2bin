package main

import (
	"log"
	"os"
)

func main() {

	// 1. Read Json File
	if len(os.Args) < 1 {
		log.Fatal("Not enough arguments")
	}

	// --gen_bin -> generate resource / --parse_file -> generate json
	type_op := os.Args[1]
	json_path := os.Args[2]
	binary_out := os.Args[3]

	if type_op == "--gen_bin" {
		err := GenBin(json_path, binary_out)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Success Generating Binary")
	} else if type_op == "--parse_file" {
		err := ParseCfg(json_path, binary_out)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Success Parsing to JSON")
	} else {
		log.Fatal("Unknown Option")
	}

}

func ParseCfg(json_path string, binary_out string) error {
	file_data, err := os.ReadFile(binary_out)
	if err != nil {
		return err
	}
	// 2. Read bin file
	data, err := LoadCfgFile(file_data)
	if err != nil {
		return err
	}

	// 3. Parse bin to json
	json_data, err := ConfigToJson(data)
	if err != nil {
		return err
	}

	// 4. Save Json to file
	err = WriteJsonFile(json_path, json_data)
	if err != nil {
		return err
	}

	return nil
}

func GenBin(json_path string, binary_out string) error {
	dat, err := os.ReadFile(json_path)
	if err != nil {
		return err
	}

	// 2. Parse JSON
	config_json, err := ParseJson(string(dat))
	if err != nil {
		return err
	}
	cfg, err := LoadData(config_json)
	if err != nil {
		return err
	}

	// 3. Write to Binary File
	err = WriteData(binary_out, cfg)
	if err != nil {
		return err
	}
	return nil
}
