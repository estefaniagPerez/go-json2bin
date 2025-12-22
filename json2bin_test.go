package json2bin

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	// 1. Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("./assets", "json2bin-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	// 2. Define file paths
	jsonInPath := filepath.Join(tmpDir, "test_in.json")
	binPath := filepath.Join(tmpDir, "test.bin")
	jsonOutPath := filepath.Join(tmpDir, "test_out.json")

	// 3. Create a sample JSON file
	jsonContent := `{
		"DevInfo":{
			"Version":"1",
			"SerialNumber":"sn123",
			"Fecha":"2023",
			"DeviceManufacturer":"manuf"
		},
		"DeviceSetting":{
			"Camera":{
				"ShutterSpeed":100,
				"ISO":200,
				"WhiteBalance":300,
				"Aperture":400,
				"DNG":true,
				"ImageSize":2
			},
			"Fly":{
				"MaxAltitude":1000,
				"MaxDistance":2000,
				"AutoReturnAltitude":300,
				"PitchSpeed":40,
				"YawRotationSpeed":50
			}
		}
	}`
	err = os.WriteFile(jsonInPath, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write sample JSON file: %v", err)
	}

	// 4. Run GenBin to convert JSON to binary
	err = GenBin(jsonInPath, binPath)
	if err != nil {
		t.Fatalf("GenBin failed: %v", err)
	}

	// 5. Run ParseConf to convert binary back to JSON
	err = ParseConf(jsonOutPath, binPath)
	if err != nil {
		t.Fatalf("ParseConf failed: %v", err)
	}

	// 6. Read the original and generated JSON files
	jsonInContent, err := os.ReadFile(jsonInPath)
	if err != nil {
		t.Fatalf("Failed to read original JSON file: %v", err)
	}
	jsonOutContent, err := os.ReadFile(jsonOutPath)
	if err != nil {
		t.Fatalf("Failed to read generated JSON file: %v", err)
	}

	// 7. Compare the JSON files
	assert.JSONEq(t, string(jsonInContent), string(jsonOutContent))
}
