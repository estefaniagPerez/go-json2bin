package json2bin

import (
	"github.com/estefaniagPerez/json2bin/internal/parser"
	"github.com/estefaniagPerez/json2bin/internal/parser/binaryfile"
	"github.com/estefaniagPerez/json2bin/internal/parser/jsonfile"
	"github.com/estefaniagPerez/json2bin/internal/parser/mapper"
)

func ParseConf(jsonPath string, binaryOut string) error {
	// 1. Handlers
	var binHandler parser.Manager[binaryfile.Config] = binaryfile.BinaryFileParser{}
	var jsonHandler parser.Manager[jsonfile.Config] = jsonfile.JSONFileParser{}

	// 2. Read bin file
	data, err := binHandler.Read(binaryOut)
	if err != nil {
		return err
	}

	// 3. Parse bin to json
	jsonData, err := mapper.ConfigToJSON(data)
	if err != nil {
		return err
	}

	// 4. Write to Json file
	err = jsonHandler.Write(jsonPath, jsonData)
	if err != nil {
		return err
	}

	return nil
}

func GenBin(jsonPath string, binaryOut string) error {
	// 1. Handlers
	var binHandler parser.Manager[binaryfile.Config] = binaryfile.BinaryFileParser{}
	var jsonHandler parser.Manager[jsonfile.Config] = jsonfile.JSONFileParser{}

	// 2. Parse JSON
	configJSON, err := jsonHandler.Read(jsonPath)
	if err != nil {
		return err
	}

	// 3. Convert to Bin Struct
	conf, err := mapper.JSONToConfig(configJSON)
	if err != nil {
		return err
	}

	// 4. Write to Binary File
	err = binHandler.Write(binaryOut, conf)
	if err != nil {
		return err
	}
	return nil
}
