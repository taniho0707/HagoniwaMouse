package widgets

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"

	giovieweditor "github.com/oligo/gioview/editor"
	"github.com/taniho0707/HagoniwaMouse/ui/fonts"
	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
)

type CodeEditor struct {
	editor *giovieweditor.Editor
	code   string

	font font.FontFace

	border widget.Border
}

func NewCodeEditor(code string, theme *hakoniwatheme.Theme) *CodeEditor {
	c := &CodeEditor{
		editor: new(giovieweditor.Editor),
		code:   code,
		font:   fonts.MustGetCodeEditorFont(),
	}

	c.border = widget.Border{
		Color:        theme.BorderColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}

	return c
}

func (c *CodeEditor) AppendCodeTail(code string) {
	c.editor.Insert(code)
	c.code += code
}

func (c *CodeEditor) SetCode(code string) {
	c.editor.SetText(code, false)
	c.code = code
}

func (c *CodeEditor) Code() string {
	return c.editor.Text()
}

func (c *CodeEditor) Layout(gtx layout.Context, theme *hakoniwatheme.Theme) layout.Dimensions {
	flexH := layout.Flex{Axis: layout.Horizontal}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return c.border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return flexH.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top:    unit.Dp(4),
							Bottom: unit.Dp(4),
							Left:   unit.Dp(8),
							Right:  unit.Dp(4),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							editorConf := &giovieweditor.EditorConf{
								Shaper:          theme.Shaper,
								TextColor:       theme.Fg,
								Bg:              theme.Bg,
								SelectionColor:  theme.TextSelectionColor,
								TypeFace:        c.font.Font.Typeface,
								TextSize:        unit.Sp(14),
								LineHeightScale: 1.2,
								ShowLineNum:     true,
								LineNumPadding:  unit.Dp(10),
							}

							return giovieweditor.NewEditor(c.editor, editorConf, "").Layout(gtx)
						})
					}),
				)
			})
		}),
	)
}
