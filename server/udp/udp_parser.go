package udp_server

import (
	"errors"
	"math"
)

var ErrInvalidCommandLength = errors.New("invalid command length")
var ErrInvalidCommandCode = errors.New("invalid command code")

func ParseUdpCommand(buf []byte) (UdpCommand, error) {
	if len(buf) < 3 {
		return UdpCommand{}, ErrInvalidCommandLength
	}

	cmd := UdpCommand{
		Code: CommandCode(buf[2]),
	}
	length := uint16(buf[0])<<8 | uint16(buf[1]) + 2

	switch cmd.Code {
	case CommandInternalNoResponse:
		return UdpCommand{}, nil
	case CommandSetMousePosition:
		cmd.MousePositionX = math.Float32frombits(uint32(buf[3])<<24 | uint32(buf[4])<<16 | uint32(buf[5])<<8 | uint32(buf[6]))
		cmd.MousePositionY = math.Float32frombits(uint32(buf[7])<<24 | uint32(buf[8])<<16 | uint32(buf[9])<<8 | uint32(buf[10]))
		cmd.MousePositionAngle = math.Float32frombits(uint32(buf[11])<<24 | uint32(buf[12])<<16 | uint32(buf[13])<<8 | uint32(buf[14]))
	case CommandSetMaze:
		cmd.MazeName = buf[3:length]
	case CommandSetMouseModel:
		cmd.MouseModelName = buf[3:length]
	case CommandSetMouseWallsensorType:
		cmd.MouseWallsensorTypeName = buf[3:length]
	case CommandSetMouseWallsensorNum:
		cmd.MouseWallsensorNum = uint8(buf[3])
	case CommandSetMazeCellMarker:
		for i := 0; i < (len(buf)-3)/3; i++ {
			cmd.MazeCellMarkers = append(cmd.MazeCellMarkers, MazeCellMarkerSet{
				X:    buf[i*3+3],
				Y:    buf[i*3+3+1],
				Type: CellMarkerType(buf[i*3+3+2]),
			})
		}
	case CommandSetPathByCell:
		for i := 0; i < (len(buf)-3)/3; i++ {
			cmd.MazePathByCell = append(cmd.MazePathByCell, MazePathByCellSet{
				X:    buf[i*3+3],
				Y:    buf[i*3+3+1],
				Type: LineType(buf[i*3+3+2]),
			})
		}
	case CommandSetPathByPosition:
		for i := 0; i < (len(buf)-3)/9; i++ {
			cmd.MazePathByPosition = append(cmd.MazePathByPosition, MazePathByPositionSet{
				X:    math.Float32frombits(uint32(buf[i*9+3])<<24 | uint32(buf[i*9+4])<<16 | uint32(buf[i*9+5])<<8 | uint32(buf[i*9+6])),
				Y:    math.Float32frombits(uint32(buf[i*9+7])<<24 | uint32(buf[i*9+8])<<16 | uint32(buf[i*9+9])<<8 | uint32(buf[i*9+10])),
				Type: LineType(buf[i*9+11]),
			})
		}
	case CommandGetMouseWallsensorValue:
	case CommandGetMouseImuValue:
	case CommandGetMouseBatteryValue:
	case CommandGetMouseEncoderValue:
	case CommandDeletePathAll:
	case CommandResultSuccess:
	case CommandResultInvalidLength:
	case CommandResultInvalidCommand:
	case CommandResultInvalidData:
		break
	default:
		return UdpCommand{}, ErrInvalidCommandCode
	}
	return cmd, nil
}
