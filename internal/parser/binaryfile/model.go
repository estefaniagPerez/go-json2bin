package binaryfile

// Config Drone Configuration
type Config struct {
	DevInfo     DevInfo
	DevSettings DevSettings
}

type DevInfo struct {
	Version               uint32
	DeviceSn              [16]byte
	DeviceManufacturer    [16]byte
	DeviceManufactureDate [16]byte
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
