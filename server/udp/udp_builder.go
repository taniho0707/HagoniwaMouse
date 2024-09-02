package udp_server

import (
	"errors"
	"math"

	udp_domain "github.com/taniho0707/HagoniwaMouse/server/domain"
)

var ErrUnsupportedCommand = errors.New("unsupported command")
var ErrNoParameters = errors.New("no parameters")

func BuildBytesFromUdpCommand(u *udp_domain.UdpCommand) ([]byte, error) {
	switch u.Code {
	case udp_domain.CommandGetMouseWallsensorValue:
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
	case udp_domain.CommandGetMouseImuValue:
		data := []byte{0x00, 0x0F, byte(u.Code)}
		data = append(data, byte(u.MouseImuTemp>>8), byte(u.MouseImuTemp))
		data = append(data, byte(u.MouseGyroX>>8), byte(u.MouseGyroX))
		data = append(data, byte(u.MouseGyroY>>8), byte(u.MouseGyroY))
		data = append(data, byte(u.MouseGyroZ>>8), byte(u.MouseGyroZ))
		data = append(data, byte(u.MouseAccX>>8), byte(u.MouseAccX))
		data = append(data, byte(u.MouseAccY>>8), byte(u.MouseAccY))
		data = append(data, byte(u.MouseAccZ>>8), byte(u.MouseAccZ))
		return data, nil
	case udp_domain.CommandGetMouseBatteryValue:
		data := []byte{0x00, 0x05, byte(u.Code)}
		batteryValueBits := math.Float32bits(u.MouseBatteryValue)
		data = append(data, byte(batteryValueBits>>24), byte(batteryValueBits>>16), byte(batteryValueBits>>8), byte(batteryValueBits))
		return data, nil
	case udp_domain.CommandGetMouseEncoderValue:
		data := []byte{0x00, 0x05, byte(u.Code)}
		data = append(data, byte(u.MouseEncoderLeft>>8), byte(u.MouseEncoderLeft))
		data = append(data, byte(u.MouseEncoderRight>>8), byte(u.MouseEncoderRight))
		return data, nil
	case udp_domain.CommandResultSuccess:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	case udp_domain.CommandResultInvalidLength:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	case udp_domain.CommandResultInvalidCommand:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	case udp_domain.CommandResultInvalidData:
		return []byte{0x00, 0x01, byte(u.Code)}, nil
	default:
		return nil, ErrUnsupportedCommand
	}
}
