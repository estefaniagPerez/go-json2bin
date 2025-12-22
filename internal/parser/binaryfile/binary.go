package binaryfile

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// Handler satisfies parser.Manager[Config]
type BinaryFileParser struct{}

func (b BinaryFileParser) Write(file string, data Config) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("error creating the binary file: %w", err)
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	err = binary.Write(f, binary.BigEndian, data)
	if err != nil {
		return fmt.Errorf("error wrinting the configuration data: %w", err)
	}

	return err
}

func (b BinaryFileParser) Read(file string) (Config, error) {
	conf := Config{}
	data, err := os.ReadFile(file)
	if err != nil {
		return conf, fmt.Errorf("error opening the binary file: %w", err)
	}

	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &conf)
	if err != nil {
		return conf, fmt.Errorf("error reading the binary data into the configuration structure: %w", err)
	}
	return conf, nil
}
