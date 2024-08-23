package logs

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
)

type C = layout.Context
type D = layout.Dimensions

type View struct {
	theme  *hakoniwatheme.Theme
	window *app.Window

	// textBox *widgets.CodeEditor
	text []string

	// callbacks
	onGetNewLog func(str string)
	onClearLog  func()
}

func NewView(w *app.Window, theme *hakoniwatheme.Theme) *View {
	return &View{
		theme:  theme,
		window: w,
	}
}

func (v *View) SetOnGetNewLog(f func(str string)) {
	v.onGetNewLog = f
}

func (v *View) SetOnClearLog(f func()) {
	v.onClearLog = f
}

func (v *View) AddLog(str string) {
	v.text = append(v.text, str)
}

func (v *View) Layout(gtx layout.Context, theme *hakoniwatheme.Theme) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(15), Right: unit.Dp(10)}

	items := []layout.FlexChild{}
	if len(v.text) == 0 {
		items = append(items,
			layout.Rigid(func(gtx C) D {
				return material.Body1(v.theme.Material(), "No Logs").Layout(gtx)
			}),
		)
	} else {
		for _, t := range v.text {
			items = append(items,
				layout.Rigid(func(gtx C) D {
					// return theme.Body1(t).Layout(gtx)
					return material.Body1(v.theme.Material(), t).Layout(gtx)
				}),
			)
		}
	}

	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Start,
		}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				// return v.textBox.Layout(gtx, theme)
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Start,
					Spacing:   layout.SpaceStart,
				}.Layout(gtx, items...)
			}),
		)
	})
}
