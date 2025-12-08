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
		return err
	}

	var s = buf.String()
	fmt.Println(s)
	err = os.WriteFile(file, buf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (j JSONFileParser) Read(jsonPath string) (Config, error) {
	var dataj Config

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return dataj, err
	}
	err = json.Unmarshal([]byte(data), &dataj)

	if err != nil {
		return dataj, err
	}
	return dataj, nil
}
