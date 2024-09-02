package simulator

import udp_domain "github.com/taniho0707/HagoniwaMouse/server/domain"

type SimState struct {
	timeMs uint32

	mousePositionX     float32
	mousePositionY     float32
	mousePositionAngle float32

	mouseWallsensorValue []uint16

	mouseGyroX int16
	mouseGyroY int16
	mouseGyroZ int16
	mouseAccX  int16
	mouseAccY  int16
	mouseAccZ  int16

	mouseBatteryValue float32

	mouseEncoderLeft  int16
	mouseEncoderRight int16

	mouseWallsensorNum uint8
}

type Simulator struct {
	last    SimState
	current SimState
}

func NewSimulator() *Simulator {
	s := &Simulator{}
	s.Reset()
	return s
}

func (s *Simulator) Reset() {
	s.last.timeMs = 0
	s.last.mousePositionX = 45
	s.last.mousePositionY = 45
	s.last.mousePositionAngle = 0
	s.last.mouseWallsensorValue = []uint16{0, 0, 0, 0}
	s.last.mouseGyroX = 0
	s.last.mouseGyroY = 0
	s.last.mouseGyroZ = 0
	s.last.mouseAccX = 0
	s.last.mouseAccY = 0
	s.last.mouseAccZ = 0
	s.last.mouseBatteryValue = 4.0
	s.last.mouseEncoderLeft = 0
	s.last.mouseEncoderRight = 0
	s.last.mouseWallsensorNum = 4
	s.current.timeMs = 0
	s.current.mousePositionX = 45
	s.current.mousePositionY = 45
	s.current.mousePositionAngle = 0
	s.current.mouseWallsensorValue = []uint16{0, 0, 0, 0}
	s.current.mouseGyroX = 0
	s.current.mouseGyroY = 0
	s.current.mouseGyroZ = 0
	s.current.mouseAccX = 0
	s.current.mouseAccY = 0
	s.current.mouseAccZ = 0
	s.current.mouseBatteryValue = 4.0
	s.current.mouseEncoderLeft = 0
	s.current.mouseEncoderRight = 0
	s.current.mouseWallsensorNum = 4
}

func (s *Simulator) Next(cmd *udp_domain.UdpCommand) error {
	// s.last = s.current
	// s.current.timeMs += 100
	// s.current.mousePositionX += 1
	// s.current.mousePositionY += 1
	// s.current.mousePositionAngle += 1
	// s.current.mouseWallsensorValue = []uint16{1, 2, 3, 4}
	// s.current.mouseGyroX += 1
	// s.current.mouseGyroY += 1
	// s.current.mouseGyroZ += 1
	// s.current.mouseAccX += 1
	// s.current.mouseAccY += 1
	// s.current.mouseAccZ += 1
	// s.current.mouseBatteryValue += 1
	// s.current.mouseEncoderLeft += 1
	// s.current.mouseEncoderRight += 1
	// s.current.mouseWallsensorNum = 4
	// cmd.MouseImuTemp = s.current.mouseGyroX
	// cmd.MouseGyroX = s.current.mouseGyroX
	// cmd.MouseGyroY = s.current.mouseGyroY
	// cmd.MouseGyroZ = s.current.mouseGyroZ
	// cmd.MouseAccX = s.current.mouseAccX
	// cmd.MouseAccY = s.current.mouseAccY
	// cmd.MouseAccZ = s.current.mouseAccZ
	return nil
}
