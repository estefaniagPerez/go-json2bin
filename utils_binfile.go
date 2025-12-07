package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"strconv"
)

// Drone Configuration
type Config struct {
	Dev_info     DevInfo
	Dev_settings DevSettings
}

type DevInfo struct {
	Version                 uint32
	Device_sn               [16]byte
	Device_manufacturer     [16]byte
	Device_manufacture_date [16]byte
}

type DevSettings struct {
	Camera CameraSettings
	Fly    FlySettings
}

type CameraSettings struct {
	ShutterSpeed uint16
	ISO          uint16
	WhiteBalance uint16
	Aperture     uint16
	DNG          uint16
	ImageSize    uint16
}

type FlySettings struct {
	MaxAltitude        uint16
	MaxDistance        uint16
	AutoReturnAltitude uint16
	PitchSpeed         uint16
	YawRotationSpeed   uint16
}

func LoadCfgFile(data []byte) (Config, error) {
	cfg := Config{}
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func LoadData(data JSONData) (Config, error) {
	var cfg Config
	var err error
	cfg.Dev_info, err = LoadDevInfo(data)
	if err != nil {
		return cfg, err
	}

	cfg.Dev_settings = LoadSettings(data)
	return cfg, nil
}

func LoadDevInfo(data JSONData) (DevInfo, error) {
	var devinfo DevInfo
	version_u, err := strconv.ParseUint(data.DevInfo.Version, 10, 32)
	if err != nil {
		return devinfo, err
	}
	devinfo.Version = uint32(version_u)
	copy(devinfo.Device_sn[:], data.DevInfo.SerialNumber)
	copy(devinfo.Device_manufacturer[:], data.DevInfo.DeviceManufacturer)
	copy(devinfo.Device_manufacture_date[:], data.DevInfo.Fecha)
	return devinfo, nil
}

func LoadSettings(data JSONData) DevSettings {
	var devsettings DevSettings

	// Camera
	devsettings.Camera.ShutterSpeed = uint16(data.DeviceSetting.Camera.ShutterSpeed)
	devsettings.Camera.ISO = uint16(data.DeviceSetting.Camera.ISO)
	devsettings.Camera.WhiteBalance = uint16(data.DeviceSetting.Camera.WhiteBalance)
	devsettings.Camera.Aperture = uint16(data.DeviceSetting.Camera.Aperture)
	devsettings.Camera.DNG = BoolToUint16(data.DeviceSetting.Camera.DNG)
	devsettings.Camera.ImageSize = uint16(data.DeviceSetting.Camera.ImageSize)

	// Flying
	devsettings.Fly.MaxAltitude = uint16(data.DeviceSetting.Fly.MaxAltitude)
	devsettings.Fly.MaxDistance = uint16(data.DeviceSetting.Fly.MaxDistance)
	devsettings.Fly.AutoReturnAltitude = uint16(data.DeviceSetting.Fly.AutoReturnAltitude)
	devsettings.Fly.PitchSpeed = uint16(data.DeviceSetting.Fly.PitchSpeed)
	devsettings.Fly.YawRotationSpeed = uint16(data.DeviceSetting.Fly.YawRotationSpeed)

	return devsettings
}

func ConfigToJson(cfg Config) (JSONData, error) {
	var data JSONData
	data = DevInfoToJson(cfg.Dev_info, data)
	data = DevSettingsToJson(cfg.Dev_settings, data)
	return data, nil
}

func DevInfoToJson(devinfo DevInfo, data JSONData) JSONData {
	data.DevInfo.Version = strconv.FormatUint(uint64(devinfo.Version), 10)
	data.DevInfo.SerialNumber = string(devinfo.Device_sn[:bytes.Index(devinfo.Device_sn[:], []byte{0})])
	data.DevInfo.DeviceManufacturer = string(devinfo.Device_manufacturer[:bytes.Index(devinfo.Device_manufacturer[:], []byte{0})])
	data.DevInfo.Fecha = string(devinfo.Device_manufacture_date[:bytes.Index(devinfo.Device_manufacture_date[:], []byte{0})])
	return data
}

func DevSettingsToJson(Dev_settings DevSettings, data JSONData) JSONData {
	// Camera
	data.DeviceSetting.Camera.ShutterSpeed = int(Dev_settings.Camera.ShutterSpeed)
	data.DeviceSetting.Camera.ISO = int(Dev_settings.Camera.ISO)
	data.DeviceSetting.Camera.WhiteBalance = int(Dev_settings.Camera.WhiteBalance)
	data.DeviceSetting.Camera.Aperture = int(Dev_settings.Camera.Aperture)
	data.DeviceSetting.Camera.DNG = uint16ToBool(Dev_settings.Camera.DNG)
	data.DeviceSetting.Camera.ImageSize = int(Dev_settings.Camera.ImageSize)

	// Fly
	data.DeviceSetting.Fly.MaxAltitude = int(Dev_settings.Fly.MaxAltitude)
	data.DeviceSetting.Fly.MaxDistance = int(Dev_settings.Fly.MaxDistance)
	data.DeviceSetting.Fly.AutoReturnAltitude = int(Dev_settings.Fly.AutoReturnAltitude)
	data.DeviceSetting.Fly.PitchSpeed = int(Dev_settings.Fly.PitchSpeed)
	data.DeviceSetting.Fly.YawRotationSpeed = int(Dev_settings.Fly.YawRotationSpeed)

	return data
}

func WriteData(file string, data Config) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = binary.Write(f, binary.BigEndian, data)
	if err != nil {
		return err
	}

	return nil
}

func ParseToString(byteArray []byte, length int) string {
	n := bytes.Index(byteArray, []byte{0})
	if n == -1 || n > length-1 {
		n = length - 1
	}
	return string(byteArray[:n])
}

func BoolToUint16(value bool) uint16 {
	if value {
		return 1
	}
	return 0
}

func uint16ToBool(value uint16) bool {
	return value == 1
}
