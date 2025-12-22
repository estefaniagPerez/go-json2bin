package mapper

import (
	"reflect"
	"testing"

	"github.com/estefaniagPerez/json2bin/internal/parser/binaryfile"
	"github.com/estefaniagPerez/json2bin/internal/parser/jsonfile"
)

func TestConfigToJSON(t *testing.T) {
	// 1. Mock binary config
	binConfig := binaryfile.Config{
		DevInfo: binaryfile.DevInfo{
			Version:               1,
			DeviceSn:              [16]byte{'s', 'n', '1', '2', '3'},
			DeviceManufacturer:    [16]byte{'m', 'a', 'n', 'u', 'f'},
			DeviceManufactureDate: [16]byte{'2', '0', '2', '3'},
		},
		DevSettings: binaryfile.DevSettings{
			Camera: binaryfile.CameraSettings{
				ShutterSpeed: 100,
				ISO:          200,
				WhiteBalance: 300,
				Aperture:     400,
				DNG:          1,
				ImageSize:    2,
			},
			Fly: binaryfile.FlySettings{
				MaxAltitude:        1000,
				MaxDistance:        2000,
				AutoReturnAltitude: 300,
				PitchSpeed:         40,
				YawRotationSpeed:   50,
			},
		},
	}

	// 2. Expected JSON config
	expectedJSONConfig := jsonfile.Config{
		DevInfo: jsonfile.DeviceInfo{
			Version:            "1",
			SerialNumber:       "sn123",
			DeviceManufacturer: "manuf",
			Fecha:              "2023",
		},
		DeviceSetting: jsonfile.DeviceSetting{
			Camera: jsonfile.CameraSettings{
				ShutterSpeed: 100,
				ISO:          200,
				WhiteBalance: 300,
				Aperture:     400,
				DNG:          true,
				ImageSize:    2,
			},
			Fly: jsonfile.FlySetting{
				MaxAltitude:        1000,
				MaxDistance:        2000,
				AutoReturnAltitude: 300,
				PitchSpeed:         40,
				YawRotationSpeed:   50,
			},
		},
	}

	// 3. Run the conversion
	jsonConfig, err := ConfigToJSON(binConfig)
	if err != nil {
		t.Fatalf("ConfigToJSON failed with error: %v", err)
	}

	// 4. Compare the results
	if !reflect.DeepEqual(jsonConfig, expectedJSONConfig) {
		t.Errorf("ConfigToJSON() = %v, want %v", jsonConfig, expectedJSONConfig)
	}
}

func TestJSONToConfig(t *testing.T) {
	// 1. Mock JSON config
	jsonConfig := jsonfile.Config{
		DevInfo: jsonfile.DeviceInfo{
			Version:            "1",
			SerialNumber:       "sn123",
			DeviceManufacturer: "manuf",
			Fecha:              "2023",
		},
		DeviceSetting: jsonfile.DeviceSetting{
			Camera: jsonfile.CameraSettings{
				ShutterSpeed: 100,
				ISO:          200,
				WhiteBalance: 300,
				Aperture:     400,
				DNG:          true,
				ImageSize:    2,
			},
			Fly: jsonfile.FlySetting{
				MaxAltitude:        1000,
				MaxDistance:        2000,
				AutoReturnAltitude: 300,
				PitchSpeed:         40,
				YawRotationSpeed:   50,
			},
		},
	}

	// 2. Expected binary config
	expectedBinConfig := binaryfile.Config{
		DevInfo: binaryfile.DevInfo{
			Version:               1,
			DeviceSn:              [16]byte{'s', 'n', '1', '2', '3'},
			DeviceManufacturer:    [16]byte{'m', 'a', 'n', 'u', 'f'},
			DeviceManufactureDate: [16]byte{'2', '0', '2', '3'},
		},
		DevSettings: binaryfile.DevSettings{
			Camera: binaryfile.CameraSettings{
				ShutterSpeed: 100,
				ISO:          200,
				WhiteBalance: 300,
				Aperture:     400,
				DNG:          1,
				ImageSize:    2,
			},
			Fly: binaryfile.FlySettings{
				MaxAltitude:        1000,
				MaxDistance:        2000,
				AutoReturnAltitude: 300,
				PitchSpeed:         40,
				YawRotationSpeed:   50,
			},
		},
	}

	// 3. Run the conversion
	binConfig, err := JSONToConfig(jsonConfig)
	if err != nil {
		t.Fatalf("JSONToConfig failed with error: %v", err)
	}

	// 4. Compare the results
	if !reflect.DeepEqual(binConfig, expectedBinConfig) {
		t.Errorf("JSONToConfig() = %v, want %v", binConfig, expectedBinConfig)
	}
}
