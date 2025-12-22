package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type JSONFileParser struct{}

func (j JSONFileParser) Write(file string, data Config) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&data)
	if err != nil {
		return fmt.Errorf("error encoding configuration to JSON: %w", err)
	}

	err = os.WriteFile(file, buf.Bytes(), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}

	return nil
}

func (j JSONFileParser) Read(jsonPath string) (Config, error) {
	var dataj Config

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return dataj, fmt.Errorf("error reading JSON file: %w", err)
	}
	err = json.Unmarshal([]byte(data), &dataj)

	if err != nil {
		return dataj, fmt.Errorf("error converting to JSON structure: %w", err)
	}
	return dataj, nil
}
