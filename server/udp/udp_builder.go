package udp_server

import (
	"errors"
	"math"
)

var ErrUnsupportedCommand = errors.New("unsupported command")
var ErrNoParameters = errors.New("no parameters")

func (u *UdpCommand) Build() ([]byte, error) {
	switch u.Code {
	case CommandGetMouseWallsensorValue:
		data := []byte{0x00, 0x00, byte(u.Code)}
		if len(u.MouseWallsensorValue) == 0 {
			return nil, ErrNoParameters
		}
		for _, v := range u.MouseWallsensorValue {
			data = append(data, byte(v>>8), byte(v))
		}
		data[0] = byte(len(data) >> 8)
		data[1] = byte(len(data))
		return data, nil
	case CommandGetMouseImuValue:
		data := []byte{0x00, 0x0F, byte(u.Code)}
		data = append(data, byte(u.MouseImuTemp>>8), byte(u.MouseImuTemp))
		data = append(data, byte(u.MouseGyroX>>8), byte(u.MouseGyroX))
		data = append(data, byte(u.MouseGyroY>>8), byte(u.MouseGyroY))
		data = append(data, byte(u.MouseGyroZ>>8), byte(u.MouseGyroZ))
		data = append(data, byte(u.MouseAccX>>8), byte(u.MouseAccX))
		data = append(data, byte(u.MouseAccY>>8), byte(u.MouseAccY))
		data = append(data, byte(u.MouseAccZ>>8), byte(u.MouseAccZ))
		return data, nil
	case CommandGetMouseBatteryValue:
		data := []byte{0x00, 0x05, byte(u.Code)}
		batteryValueBits := math.Float32bits(u.MouseBatteryValue)
		data = append(data, byte(batteryValueBits>>24), byte(batteryValueBits>>16), byte(batteryValueBits>>8), byte(batteryValueBits))
		return data, nil
	case CommandGetMouseEncoderValue:
		data := []byte{0x00, 0x05, byte(u.Code)}
		data = append(data, byte(u.MouseEncoderLeft>>8), byte(u.MouseEncoderLeft))
		data = append(data, byte(u.MouseEncoderRight>>8), byte(u.MouseEncoderRight))
		return data, nil
	case CommandResultSuccess:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	case CommandResultInvalidLength:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	case CommandResultInvalidCommand:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	case CommandResultInvalidData:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	default:
		return nil, ErrUnsupportedCommand
	}
}
