package widgets

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	"image/png"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/taniho0707/HagoniwaMouse/internal/mazedata"
	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
)

const (
	PillarWidth = 6
	WallWidth   = 84
)

//go:embed maze/mouse.png
var mouseImageFile embed.FS

type zoommode int

const (
	Zoom32 zoommode = iota
	Zoom16
	Zoom8
	Zoom4
	Zoom2
)

type zoomcenter int

const (
	ZoomCenterMaze zoomcenter = iota
	ZoomCenterMouse
)

type Position struct {
	X float64
	Y float64
}

type Maze struct {
	// MazeData *mazedata.MazeData

	Zoom       zoommode
	ZoomCenter zoomcenter

	BackgroundColor  color.NRGBA
	ExistWallColor   color.NRGBA
	UnknownWallColor color.NRGBA

	MouseImage     *image.Image
	mouseImageOp   paint.ImageOp
	MouseAngle     float64 // degree
	MouseCentorPos image.Point

	MouseAbsolutePos  Position // mm 系での絶対座標
	MouseAbsoluteSize Position // mm 系での絶対サイズ
}

func NewMaze() *Maze {
	m := &Maze{
		Zoom:             Zoom32,
		ZoomCenter:       ZoomCenterMaze,
		BackgroundColor:  color.NRGBA{0xFF, 0xFF, 0xCC, 0xFF},
		ExistWallColor:   color.NRGBA{0xFF, 0x00, 0x00, 0xFF},
		UnknownWallColor: color.NRGBA{0x44, 0x88, 0xFF, 0xFF},
	}

	imgFile, err := mouseImageFile.ReadFile("maze/mouse.png")
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(bytes.NewReader(imgFile))
	if err != nil {
		panic(err)
	}
	m.MouseImage = &img
	m.mouseImageOp = paint.NewImageOp(*m.MouseImage)
	m.MouseAbsolutePos = Position{X: 45 + PillarWidth/2, Y: 45 + PillarWidth/2} // mm
	m.MouseAbsoluteSize = Position{X: 38, Y: 58}                                // mm
	m.MouseCentorPos = image.Point{X: 197, Y: 363}                              // pixel

	return m
}

func (m *Maze) mmToPixelRatio(minWidthHeight int) float64 {
	var ratio float64
	switch m.Zoom {
	case Zoom32:
		ratio = float64(minWidthHeight) / 2886.0
	case Zoom16:
		ratio = float64(minWidthHeight) / 1446.0
	case Zoom8:
		ratio = float64(minWidthHeight) / 726.0
	case Zoom4:
		ratio = float64(minWidthHeight) / 366.0
	case Zoom2:
		ratio = float64(minWidthHeight) / 186.0
	}
	return ratio
}

func (m *Maze) convertMmToPixelMaze(mmX, mmY float64, windowX, windowY int, center zoomcenter) image.Point {
	var minWidthHeight int
	if windowX > windowY {
		minWidthHeight = windowY
	} else {
		minWidthHeight = windowX
	}

	mmToPixelRatio := m.mmToPixelRatio(minWidthHeight)

	switch center {
	case ZoomCenterMaze:
		return image.Point{
			X: int(mmX * mmToPixelRatio),
			Y: windowY - int(mmY*mmToPixelRatio),
		}
	case ZoomCenterMouse:
		return image.Point{
			X: windowX/2 + int((mmX-m.MouseAbsolutePos.X)*mmToPixelRatio),
			Y: windowY/2 - int((m.MouseAbsolutePos.Y-mmY)*mmToPixelRatio),
		}
	default:
		return image.Point{}
	}
}

func (m *Maze) convertMmToPixelMouse(mmX, mmY float64, windowX, windowY int, center zoomcenter) image.Point {
	var minWidthHeight int
	if windowX > windowY {
		minWidthHeight = windowY
	} else {
		minWidthHeight = windowX
	}

	mmToPixelRatio := m.mmToPixelRatio(minWidthHeight)
	currentOffset := image.Point{
		X: m.MouseCentorPos.X,
		Y: windowY - m.mouseImageOp.Size().Y + m.MouseCentorPos.Y,
	}

	switch center {
	case ZoomCenterMaze:
		return image.Point{
			X: int(mmX*mmToPixelRatio) - currentOffset.X,
			Y: windowY - int(mmY*mmToPixelRatio) - currentOffset.Y,
		}
	case ZoomCenterMouse:
		return image.Point{
			X: windowX/2 - m.MouseCentorPos.X - currentOffset.X,
			Y: windowY/2 - m.MouseCentorPos.Y - currentOffset.Y,
		}
	default:
		return image.Point{}
	}
}

func (m *Maze) Layout(gtx C, theme *hakoniwatheme.Theme, maze *mazedata.MazeData) D {
	// 1. setup layout
	//   - get size
	width := gtx.Constraints.Max.X
	height := gtx.Constraints.Max.Y
	//   - draw background
	backgroundClip := clip.Rect(image.Rect(0, 0, width, height))
	paint.FillShape(gtx.Ops, m.BackgroundColor, backgroundClip.Op())

	// 2. setup maze
	//   - draw piller
	for x := 0; x < 33; x++ {
		for y := 0; y < 33; y++ {
			pos0 := m.convertMmToPixelMaze(float64(x)*90.0, float64(y)*90.0, width, height, m.ZoomCenter)
			pos1 := m.convertMmToPixelMaze(float64(x)*90.0+PillarWidth, float64(y)*90.0+PillarWidth, width, height, m.ZoomCenter)
			pillarClip := image.Rect(pos0.X, pos0.Y, pos1.X, pos1.Y)
			if insideRect(image.Rect(0, 0, width, height), pillarClip) {
				// paint.FillShape(gtx.Ops, m.WallColor, pillarClip.Op())
				paint.FillShape(gtx.Ops, m.ExistWallColor, clip.Rect(pillarClip).Op())
			}
		}
	}
	//   - draw horizontal  wall
	for x := 0; x < 32; x++ {
		for y := 0; y < 33; y++ {
			if y == 0 || y == 32 || maze.HorizontalWalls[x][y-1] == mazedata.WallExist || maze.HorizontalWalls[x][y-1] == mazedata.WallUnknown {
				pos0 := m.convertMmToPixelMaze(float64(x)*90.0+PillarWidth, float64(y)*90.0, width, height, m.ZoomCenter)
				pos1 := m.convertMmToPixelMaze(float64(x)*90.0+PillarWidth+WallWidth, float64(y)*90.0+PillarWidth, width, height, m.ZoomCenter)
				horizontalWallClip := image.Rect(pos0.X, pos0.Y, pos1.X, pos1.Y)
				if insideRect(image.Rect(0, 0, width, height), horizontalWallClip) {
					color := m.UnknownWallColor
					if y == 0 || y == 32 || maze.HorizontalWalls[x][y-1] == mazedata.WallExist {
						color = m.ExistWallColor
					}
					paint.FillShape(gtx.Ops, color, clip.Rect(horizontalWallClip).Op())
				}
			}
		}
	}
	//   - draw vertical wall
	for x := 0; x < 33; x++ {
		for y := 0; y < 32; y++ {
			if x == 0 || x == 32 || maze.VerticalWalls[x-1][y] == mazedata.WallExist || maze.VerticalWalls[x-1][y] == mazedata.WallUnknown {
				pos0 := m.convertMmToPixelMaze(float64(x)*90.0, float64(y)*90.0+PillarWidth, width, height, m.ZoomCenter)
				pos1 := m.convertMmToPixelMaze(float64(x)*90.0+PillarWidth, float64(y)*90.0+PillarWidth+WallWidth, width, height, m.ZoomCenter)
				verticalWallClip := image.Rect(pos0.X, pos0.Y, pos1.X, pos1.Y)
				if insideRect(image.Rect(0, 0, width, height), verticalWallClip) {
					color := m.UnknownWallColor
					if x == 0 || x == 32 || maze.VerticalWalls[x-1][y] == mazedata.WallExist {
						color = m.ExistWallColor
					}
					paint.FillShape(gtx.Ops, color, clip.Rect(verticalWallClip).Op())
				}
			}
		}
	}

	// 3. setup mouse
	m.mouseLayout(gtx)

	// 4. setup path
	// TODO: draw path

	return D{
		Size: image.Point{
			X: width,
			Y: height,
		},
	}
}

func (m *Maze) mouseLayout(gtx C) {
	width := gtx.Constraints.Max.X
	height := gtx.Constraints.Max.Y
	var minWidthHeight int
	if width > height {
		minWidthHeight = height
	} else {
		minWidthHeight = width
	}
	mmToPixelRatio := m.mmToPixelRatio(minWidthHeight)

	tr := f32.Affine2D{}

	// rotate
	rotateRadian := m.MouseAngle * math.Pi / 180.0
	tr = tr.Rotate(
		f32.Point{
			X: float32(m.MouseCentorPos.X),
			Y: float32(height - m.mouseImageOp.Size().Y + m.MouseCentorPos.Y),
		},
		float32(rotateRadian))

	// scale
	scaledPixelSize := image.Point{
		X: int(m.MouseAbsoluteSize.X * mmToPixelRatio),
		Y: int(m.MouseAbsoluteSize.Y * mmToPixelRatio),
	}
	tr = tr.Scale(
		f32.Point{
			X: float32(m.MouseCentorPos.X),
			Y: float32(height - m.mouseImageOp.Size().Y + m.MouseCentorPos.Y),
		},
		f32.Point{
			X: float32(scaledPixelSize.X) / float32((*m.MouseImage).Bounds().Size().X),
			Y: float32(scaledPixelSize.Y) / float32((*m.MouseImage).Bounds().Size().Y),
		})

	// offset
	to := m.convertMmToPixelMouse(m.MouseAbsolutePos.X+PillarWidth/2, m.MouseAbsolutePos.Y+PillarWidth/2, width, height, m.ZoomCenter)
	tr = tr.Offset(f32.Point{
		X: float32(to.X),
		Y: float32(to.Y),
	})

	// image
	w := widget.Image{
		Src:      m.mouseImageOp,
		Fit:      widget.Unscaled,
		Position: layout.SW,
		Scale:    1,
	}

	// draw
	op.Affine(tr).Add(gtx.Ops)
	w.Layout(gtx)
}

func insideRect(whole image.Rectangle, target image.Rectangle) bool {
	return whole.Min.X <= target.Max.X && whole.Min.Y <= target.Max.Y && whole.Max.X >= target.Min.X && whole.Max.Y >= target.Min.Y
}
