package mazedata

type WallState int

const (
	WallExist WallState = iota
	WallNotExist
	WallUnknown
	WallInvalid
)

type Angle int

const (
	North Angle = iota
	East
	South
	West
)

type CoordinatePosition struct {
	X int
	Y int
}

type MazeData struct {
	// サイズは32x32固定
	Start CoordinatePosition
	Goal  []CoordinatePosition

	// 縦壁
	// 外周は常に WallExist
	VerticalWalls [31][32]WallState

	// 横壁
	// 外周は常に WallExist
	HorizontalWalls [32][31]WallState
}

func NewMazeDataBlank() *MazeData {
	m := &MazeData{}
	m.Start = CoordinatePosition{X: 0, Y: 0}
	m.Goal = []CoordinatePosition{
		{X: 15, Y: 15},
		{X: 16, Y: 15},
		{X: 15, Y: 16},
		{X: 16, Y: 16},
	}
	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			if x < 31 {
				m.VerticalWalls[x][y] = WallUnknown
			}
			if y < 31 {
				m.HorizontalWalls[x][y] = WallUnknown
			}
		}
	}
	return m
}

func NewMazeDataFromFile(file []byte) (*MazeData, error) {
	mazeData, err := parseFromKerikun11(file)
	if err != nil {
		return nil, err
	}
	return mazeData, nil
}

func (m *MazeData) GetWallState(x, y int, a Angle) WallState {
	if x < 0 || x > 32 || y < 0 || y > 32 {
		return WallInvalid
	}
	if x == 0 && a == West {
		return WallExist
	} else if x == 32 && a == East {
		return WallExist
	} else if y == 0 && a == South {
		return WallExist
	} else if y == 32 && a == North {
		return WallExist
	}
	if a == North {
		return m.HorizontalWalls[x][y]
	} else if a == East {
		return m.VerticalWalls[x][y]
	} else if a == South {
		return m.HorizontalWalls[x][y-1]
	} else if a == West {
		return m.VerticalWalls[x-1][y]
	} else {
		return WallInvalid
	}
}

func (m *MazeData) SetWallState(x, y int, a Angle, state WallState) {
	if x < 0 || x > 32 || y < 0 || y > 32 {
		return
	}
	if x == 0 && a == West {
		return
	} else if x == 32 && a == East {
		return
	} else if y == 0 && a == South {
		return
	} else if y == 32 && a == North {
		return
	}
	if a == North {
		m.HorizontalWalls[x][y] = state
	} else if a == East {
		m.VerticalWalls[x][y] = state
	} else if a == South {
		m.HorizontalWalls[x][y-1] = state
	} else if a == West {
		m.VerticalWalls[x-1][y] = state
	}
}
