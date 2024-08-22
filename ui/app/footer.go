package app

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
	"github.com/taniho0707/HagoniwaMouse/ui/widgets"
)

type Footer struct {
	theme *hakoniwatheme.Theme

	infoLabel   material.LabelStyle
	information string
}

func NewFooter(theme *hakoniwatheme.Theme) *Footer {
	f := &Footer{
		theme: theme,
	}
	f.infoLabel = widgets.MaterialIcons("info", theme)
	f.information = "Display loaded"
	return f
}

func (f *Footer) SetText(text string) {
	f.information = text
}

func (f *Footer) SetIcon(icon string) {
	f.infoLabel = widgets.MaterialIcons(icon, f.theme)
}

func (f *Footer) Layout(gtx layout.Context, theme *hakoniwatheme.Theme) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(4)}

	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// return f.infoLabel.Layout(gtx, f.theme.TextSize, f.information)
		return material.Label(f.theme.Theme, f.theme.TextSize, f.information).Layout(gtx)
	})
}
