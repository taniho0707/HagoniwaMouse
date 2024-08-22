package fonts

import (
	"embed"
	"fmt"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed fonts/*
var fonts embed.FS

// func Prepare() ([]font.FontFace, error) {

// }

func getFont(path string) ([]byte, error) {
	data, err := fonts.ReadFile(fmt.Sprintf("fonts/%s", path))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func MustGetCodeEditorFont() font.FontFace {
	// data, err := getFont("source_sans_pro_regular.otf")
	data, err := getFont("HackGenConsole-Regular.ttf")
	if err != nil {
		panic(err)
	}

	monoFont, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}

	return font.FontFace{Font: font.Font{}, Face: monoFont}
}
