package widgets

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/x/component"

	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type SplitView struct {
	BarWidth unit.Dp
	component.Resize
}

const defaultBarWidth = 2

func (s *SplitView) Layout(gtx C, theme *hakoniwatheme.Theme, left, right layout.Widget) D {
	bar := gtx.Dp(s.BarWidth)
	if bar <= 1 {
		bar = gtx.Dp(defaultBarWidth)
	}

	return s.Resize.Layout(gtx,
		func(gtx C) D {
			return left(gtx)
		},
		func(gtx C) D {
			return right(gtx)
		},
		func(gtx C) D {
			rect := image.Rectangle{
				Max: image.Point{
					X: bar,
					Y: gtx.Constraints.Max.Y,
				},
			}
			paint.FillShape(gtx.Ops, theme.SeparatorColor, clip.Rect(rect).Op())
			return D{Size: rect.Max}
		},
	)
}
