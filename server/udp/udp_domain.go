package udp_server

type CellMarkerType uint8

const (
	CellMarkerTypeStart CellMarkerType = iota
	CellMarkerTypeGoal
	CellMarkerTypeTarget
)

type MazeCellMarkerSet struct {
	X    uint8
	Y    uint8
	Type CellMarkerType
}

type LineType uint8

const (
	LineTypeNormal LineType = iota
	LineTypeSub
)

type MazePathByCellSet struct {
	X    uint8
	Y    uint8
	Type LineType
}

type MazePathByPositionSet struct {
	X    float32
	Y    float32
	Type LineType
}

type UdpCommand struct {
	Code CommandCode

	// Code によって下記のいずれかの値を設定する
	// TODO: 微妙すぎるのでアーキテクチャを見直す

	// SetMousePosition
	MousePositionX     float32
	MousePositionY     float32
	MousePositionAngle float32

	// GetMouseWallsensorValue
	MouseWallsensorValue []uint16

	// GetMouseImuValue
	MouseImuTemp int16
	MouseGyroX   int16
	MouseGyroY   int16
	MouseGyroZ   int16
	MouseAccX    int16
	MouseAccY    int16
	MouseAccZ    int16

	// GetMouseBatteryValue
	MouseBatteryValue float32

	// GetMouseEncoderValue
	MouseEncoderLeft  int16
	MouseEncoderRight int16

	// SetMaze
	MazeName []byte

	// SetMouseModel
	MouseModelName []byte

	// SetMouseWallsensorType
	MouseWallsensorTypeName []byte

	// SetMouseWallsensorNum
	MouseWallsensorNum uint8

	// SetMazeCellMarker
	MazeCellMarkers []MazeCellMarkerSet

	// SetPathByCell
	MazePathByCell []MazePathByCellSet

	// SetPathByPosition
	MazePathByPosition []MazePathByPositionSet
}

type CommandCode uint8

// Get / Set はマウス視点での命名
const (
	CommandInternalNoResponse CommandCode = 0x00
	// 0x01 - 0x0F : Mouse Physical Status
	CommandSetMousePosition CommandCode = 0x01
	// reserve 0x20 - 0x3F for future use
	// 0x40 - 0x4F : Mouse Sensor Status
	CommandGetMouseWallsensorValue CommandCode = 0x40
	CommandGetMouseImuValue        CommandCode = 0x41
	CommandGetMouseBatteryValue    CommandCode = 0x42
	CommandGetMouseEncoderValue    CommandCode = 0x43
	// reserve 0x50 - 0x5F for future use
	// 0x60 - 0x6F : Mouse Setting
	CommandSetMaze                CommandCode = 0x60
	CommandSetMouseModel          CommandCode = 0x64
	CommandSetMouseWallsensorType CommandCode = 0x65
	CommandSetMouseWallsensorNum  CommandCode = 0x66
	// reserve 0x70 - 0x7F for future use
	// 0x80 - 0x8F : Maze Cell Setting
	CommandSetMazeCellMarker CommandCode = 0x80
	// 0x90 - 0x9F : Maze Path Setting
	CommandSetPathByCell     CommandCode = 0x90
	CommandSetPathByPosition CommandCode = 0x91
	CommandDeletePathAll     CommandCode = 0x9F
	// reserve 0xA0 - 0xAF for parameter monitor commands
	// reserve 0xB0 - 0xBF for flash monitor commands
	// reserve 0xC0 - 0xCF for future use
	// reserve 0xD0 - 0xDF for future use
	// reserve 0xE0 - 0xEF for future use
	// 0xF0 - 0xFF : Command Result
	CommandResultSuccess        CommandCode = 0xF0
	CommandResultInvalidLength  CommandCode = 0xF1
	CommandResultInvalidCommand CommandCode = 0xF2
	CommandResultInvalidData    CommandCode = 0xFF
)
