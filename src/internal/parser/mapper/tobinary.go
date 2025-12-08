package mapper

import (
	"bytes"
	"localhost/estefaniagPerez/json2bin/internal/helper"
	"localhost/estefaniagPerez/json2bin/internal/parser/binaryfile"
	"localhost/estefaniagPerez/json2bin/internal/parser/jsonfile"
	"strconv"
)

func ConfigToJSON(conf binaryfile.Config) (jsonfile.Config, error) {
	var data jsonfile.Config
	data = devInfoToJSON(conf.DevInfo, data)
	data = devSettingsToJSON(conf.DevSettings, data)
	return data, nil
}

func devInfoToJSON(devinfo binaryfile.DevInfo, data jsonfile.Config) jsonfile.Config {
	data.DevInfo.Version = strconv.FormatUint(uint64(devinfo.Version), 10)
	data.DevInfo.SerialNumber = string(devinfo.DeviceSn[:bytes.Index(devinfo.DeviceSn[:], []byte{0})])
	data.DevInfo.DeviceManufacturer = string(devinfo.DeviceManufacturer[:bytes.Index(devinfo.DeviceManufacturer[:], []byte{0})])
	data.DevInfo.Fecha = string(devinfo.DeviceManufactureDate[:bytes.Index(devinfo.DeviceManufactureDate[:], []byte{0})])
	return data
}

func devSettingsToJSON(DevSettings binaryfile.DevSettings, data jsonfile.Config) jsonfile.Config {
	// Camera
	data.DeviceSetting.Camera.ShutterSpeed = int(DevSettings.Camera.ShutterSpeed)
	data.DeviceSetting.Camera.ISO = int(DevSettings.Camera.ISO)
	data.DeviceSetting.Camera.WhiteBalance = int(DevSettings.Camera.WhiteBalance)
	data.DeviceSetting.Camera.Aperture = int(DevSettings.Camera.Aperture)
	data.DeviceSetting.Camera.DNG = helper.Uint16ToBool(DevSettings.Camera.DNG)
	data.DeviceSetting.Camera.ImageSize = int(DevSettings.Camera.ImageSize)

	// Fly
	data.DeviceSetting.Fly.MaxAltitude = int(DevSettings.Fly.MaxAltitude)
	data.DeviceSetting.Fly.MaxDistance = int(DevSettings.Fly.MaxDistance)
	data.DeviceSetting.Fly.AutoReturnAltitude = int(DevSettings.Fly.AutoReturnAltitude)
	data.DeviceSetting.Fly.PitchSpeed = int(DevSettings.Fly.PitchSpeed)
	data.DeviceSetting.Fly.YawRotationSpeed = int(DevSettings.Fly.YawRotationSpeed)

	return data
}
