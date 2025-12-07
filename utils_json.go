package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type JSONData struct {
	DevInfo       JSONDeviceInfo
	DeviceSetting JSONDeviceSetting
}

type JSONDeviceInfo struct {
	Version            string
	SerialNumber       string
	Fecha              string
	DeviceManufacturer string
}

type JSONDeviceSetting struct {
	Camera JSONCameraSettings
	Fly    JSONFlySetting
}

type JSONCameraSettings struct {
	ShutterSpeed int  // 1 -> 1 | 4 -> 1/4 | 15 -> 1/15 | 60 -> 1/60 | 125 -> 1/125 | 500 -> 1/500 | 1000 -> 1/1000
	ISO          int  // 100 | 200 | 400 | 800 | 1600
	WhiteBalance int  // 0 -> Auto | 3000 | 5500 | 6500
	Aperture     int  // 14 -> f/1.4 | 20 -> f/2 | 28 -> f/2.8 | 40 -> f/4 | 56 -> f/56 | 80 -> f/8 | 110 -> f/11 | 160 -> f/16
	DNG          bool // True -> Enable | False -> Disable
	ImageSize    int  // 0 -> 4:3 | 1 -> 3:2 | 2 -> 16:9
}

type JSONFlySetting struct {
	MaxAltitude        int
	MaxDistance        int
	AutoReturnAltitude int
	PitchSpeed         int
	YawRotationSpeed   int
}

func ParseJson(data string) (JSONData, error) {
	var dataj JSONData
	err := json.Unmarshal([]byte(data), &dataj)

	if err != nil {
		return dataj, err
	}
	return dataj, nil
}

func JsonToString(dataj JSONData) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&dataj)

	if err != nil {
		return "{\"Status\":\"-1\"}", err
	}

	return buf.String(), nil
}

func WriteJsonFile(file string, data JSONData) error {
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
