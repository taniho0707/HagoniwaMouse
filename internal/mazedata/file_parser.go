package mazedata

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidMazeData = errors.New("invalid maze data")

func parseFromKerikun11(file []byte) (*MazeData, error) {
	// 1行目を読んで迷路サイズを推定
	fileString := string(file)
	fileString = strings.Replace(fileString, "\r\n", "\n", -1)
	fileString = strings.Replace(fileString, "\r", "\n", -1)
	lines := strings.Split(fileString, "\n")
	reHorizontalWall := regexp.MustCompile(`\+---`)
	widthMatch := reHorizontalWall.FindAll([]byte(lines[0]), -1)
	if widthMatch == nil {
		return nil, ErrInvalidMazeData
	}
	width := len(widthMatch)
	if len(lines) < width*2+1 {
		return nil, ErrInvalidMazeData
	}

	// 2*n / 2*n+1 行目を読んで壁情報を取得
	mazeData := &MazeData{}
	for n := 0; n < width; n++ {
		invN := width - n - 1 // 迷路のY座標に相当する
		// 縦壁
		for x := 0; x < width-1; x++ {
			target := lines[2*n+1][4*x+4] // 右側の壁
			if target == ' ' {
				mazeData.VerticalWalls[x][invN] = WallNotExist
			} else if target == '|' {
				mazeData.VerticalWalls[x][invN] = WallExist
			} else if target == '.' {
				mazeData.VerticalWalls[x][invN] = WallUnknown
			} else {
				mazeData.VerticalWalls[x][invN] = WallInvalid
			}
		}
	}
	for n := 0; n < width-1; n++ {
		invN := width - n - 2 // 迷路のY座標に相当する
		// 横壁
		for x := 0; x < width; x++ {
			target := lines[2*n+2][4*x+2] // 上側の壁
			if target == ' ' {
				mazeData.HorizontalWalls[x][invN] = WallNotExist
			} else if target == '-' {
				mazeData.HorizontalWalls[x][invN] = WallExist
			} else if target == '.' {
				mazeData.HorizontalWalls[x][invN] = WallUnknown
			} else {
				mazeData.HorizontalWalls[x][invN] = WallInvalid
			}
		}
	}
	for n := 0; n < width; n++ {
		invN := width - n // 迷路のY座標に相当する
		// S/G
		for x := 0; x < width; x++ {
			target := lines[2*n+1][4*x+2]
			if target == 'S' {
				mazeData.Start = CoordinatePosition{X: x, Y: invN}
			} else if target == 'G' {
				mazeData.Goal = append(mazeData.Goal, CoordinatePosition{X: x, Y: invN})
			}
		}
	}

	return mazeData, nil
}
