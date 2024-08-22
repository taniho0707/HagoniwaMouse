package logs

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
	"github.com/taniho0707/HagoniwaMouse/ui/widgets"
)

type View struct {
	theme  *hakoniwatheme.Theme
	window *app.Window

	textBox *widgets.CodeEditor

	// callbacks
	onGetNewLog func(str string)
	onClearLog  func()
}

func NewView(w *app.Window, theme *hakoniwatheme.Theme) *View {
	return &View{
		theme:   theme,
		window:  w,
		textBox: widgets.NewCodeEditor("", theme),
	}
}

func (v *View) SetOnGetNewLog(f func(str string)) {
	v.onGetNewLog = f
}

func (v *View) SetOnClearLog(f func()) {
	v.onClearLog = f
}

func (v *View) Layout(gtx layout.Context, theme *hakoniwatheme.Theme) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(15), Right: unit.Dp(10)}
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Start,
		}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return v.textBox.Layout(gtx, theme)
			}),
		)
	})
}
