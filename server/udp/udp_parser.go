package udp_server

import (
	"errors"
	"math"

	udp_domain "github.com/taniho0707/HagoniwaMouse/server/domain"
)

var ErrInvalidCommandLength = errors.New("invalid command length")
var ErrInvalidCommandCode = errors.New("invalid command code")

func ParseUdpCommand(buf []byte) (udp_domain.UdpCommand, error) {
	if len(buf) < 3 {
		return udp_domain.UdpCommand{}, ErrInvalidCommandLength
	}

	cmd := udp_domain.UdpCommand{
		Code: udp_domain.CommandCode(buf[2]),
	}
	length := uint16(buf[0])<<8 | uint16(buf[1]) + 2

	switch cmd.Code {
	case udp_domain.CommandInternalNoResponse:
		return udp_domain.UdpCommand{}, nil
	case udp_domain.CommandSetMousePosition:
		cmd.MousePositionX = math.Float32frombits(uint32(buf[3])<<24 | uint32(buf[4])<<16 | uint32(buf[5])<<8 | uint32(buf[6]))
		cmd.MousePositionY = math.Float32frombits(uint32(buf[7])<<24 | uint32(buf[8])<<16 | uint32(buf[9])<<8 | uint32(buf[10]))
		cmd.MousePositionAngle = math.Float32frombits(uint32(buf[11])<<24 | uint32(buf[12])<<16 | uint32(buf[13])<<8 | uint32(buf[14]))
	case udp_domain.CommandSetMaze:
		cmd.MazeName = buf[3:length]
	case udp_domain.CommandSetMouseModel:
		cmd.MouseModelName = buf[3:length]
	case udp_domain.CommandSetMouseWallsensorType:
		cmd.MouseWallsensorTypeName = buf[3:length]
	case udp_domain.CommandSetMouseWallsensorNum:
		cmd.MouseWallsensorNum = uint8(buf[3])
	case udp_domain.CommandSetMazeCellMarker:
		for i := 0; i < (len(buf)-3)/3; i++ {
			cmd.MazeCellMarkers = append(cmd.MazeCellMarkers, udp_domain.MazeCellMarkerSet{
				X:    buf[i*3+3],
				Y:    buf[i*3+3+1],
				Type: udp_domain.CellMarkerType(buf[i*3+3+2]),
			})
		}
	case udp_domain.CommandSetPathByCell:
		for i := 0; i < (len(buf)-3)/3; i++ {
			cmd.MazePathByCell = append(cmd.MazePathByCell, udp_domain.MazePathByCellSet{
				X:    buf[i*3+3],
				Y:    buf[i*3+3+1],
				Type: udp_domain.LineType(buf[i*3+3+2]),
			})
		}
	case udp_domain.CommandSetPathByPosition:
		for i := 0; i < (len(buf)-3)/9; i++ {
			cmd.MazePathByPosition = append(cmd.MazePathByPosition, udp_domain.MazePathByPositionSet{
				X:    math.Float32frombits(uint32(buf[i*9+3])<<24 | uint32(buf[i*9+4])<<16 | uint32(buf[i*9+5])<<8 | uint32(buf[i*9+6])),
				Y:    math.Float32frombits(uint32(buf[i*9+7])<<24 | uint32(buf[i*9+8])<<16 | uint32(buf[i*9+9])<<8 | uint32(buf[i*9+10])),
				Type: udp_domain.LineType(buf[i*9+11]),
			})
		}
	case udp_domain.CommandGetMouseWallsensorValue:
	case udp_domain.CommandGetMouseImuValue:
	case udp_domain.CommandGetMouseBatteryValue:
	case udp_domain.CommandGetMouseEncoderValue:
	case udp_domain.CommandDeletePathAll:
	case udp_domain.CommandResultSuccess:
	case udp_domain.CommandResultInvalidLength:
	case udp_domain.CommandResultInvalidCommand:
	case udp_domain.CommandResultInvalidData:
		break
	default:
		return udp_domain.UdpCommand{}, ErrInvalidCommandCode
	}
	return cmd, nil
}
