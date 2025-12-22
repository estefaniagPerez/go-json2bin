package mapper

import (
	"fmt"
	"strconv"

	"github.com/estefaniagPerez/json2bin/internal/helper"
	"github.com/estefaniagPerez/json2bin/internal/parser/binaryfile"
	"github.com/estefaniagPerez/json2bin/internal/parser/jsonfile"
)

func JSONToConfig(data jsonfile.Config) (binaryfile.Config, error) {
	var conf binaryfile.Config
	var err error
	conf.DevInfo, err = loadDevInfo(data)
	if err != nil {
		return conf, fmt.Errorf("error joading JSON device information: %w", err)
	}

	conf.DevSettings = loadSettings(data)
	return conf, nil
}

func loadDevInfo(data jsonfile.Config) (binaryfile.DevInfo, error) {
	var devinfo binaryfile.DevInfo
	versionU, err := strconv.ParseUint(data.DevInfo.Version, 10, 32)
	if err != nil {
		return devinfo, fmt.Errorf("error converting version string to uint: %w", err)
	}
	devinfo.Version = uint32(versionU)
	copy(devinfo.DeviceSn[:], data.DevInfo.SerialNumber)
	copy(devinfo.DeviceManufacturer[:], data.DevInfo.DeviceManufacturer)
	copy(devinfo.DeviceManufactureDate[:], data.DevInfo.Fecha)
	return devinfo, nil
}

func loadSettings(data jsonfile.Config) binaryfile.DevSettings {
	var devsettings binaryfile.DevSettings

	// Camera
	devsettings.Camera.ShutterSpeed = uint16(data.DeviceSetting.Camera.ShutterSpeed)
	devsettings.Camera.ISO = uint16(data.DeviceSetting.Camera.ISO)
	devsettings.Camera.WhiteBalance = uint16(data.DeviceSetting.Camera.WhiteBalance)
	devsettings.Camera.Aperture = uint16(data.DeviceSetting.Camera.Aperture)
	devsettings.Camera.DNG = helper.BoolToUint16(data.DeviceSetting.Camera.DNG)
	devsettings.Camera.ImageSize = uint16(data.DeviceSetting.Camera.ImageSize)

	// Flying
	devsettings.Fly.MaxAltitude = uint16(data.DeviceSetting.Fly.MaxAltitude)
	devsettings.Fly.MaxDistance = uint16(data.DeviceSetting.Fly.MaxDistance)
	devsettings.Fly.AutoReturnAltitude = uint16(data.DeviceSetting.Fly.AutoReturnAltitude)
	devsettings.Fly.PitchSpeed = uint16(data.DeviceSetting.Fly.PitchSpeed)
	devsettings.Fly.YawRotationSpeed = uint16(data.DeviceSetting.Fly.YawRotationSpeed)

	return devsettings
}
