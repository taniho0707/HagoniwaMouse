package maze

import (
	"github.com/taniho0707/HagoniwaMouse/internal/mazedata"
	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
	"github.com/taniho0707/HagoniwaMouse/ui/widgets"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	giox "gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type View struct {
	window *app.Window
	split  widgets.SplitView

	// maze
	mazeWidget *widgets.Maze

	mazeData *mazedata.MazeData

	onMazeOpen  func()
	onAllReset  func()
	onZoomPlus  func()
	onZoomMinus func()
	onZoomReset func()

	// mouse on maze

	// path on maze

	// seekbar

	// info
	titleInfoLabel material.LabelStyle

	// input
	titleInputLabel material.LabelStyle

	// control
	titleControlLabel   material.LabelStyle
	mazeOpenButton      widget.Clickable
	mazeFileName        string
	allResetButton      widget.Clickable
	mazeZoomPlusButton  widget.Clickable
	mazeZoomMinusButton widget.Clickable
	mazeZoomResetButton widget.Clickable
}

func NewView(w *app.Window, theme *hakoniwatheme.Theme) *View {
	v := &View{
		window: w,

		split: widgets.SplitView{
			Resize: giox.Resize{
				Ratio: 0.8,
			},
			BarWidth: unit.Dp(2),
		},

		mazeWidget: widgets.NewMaze(),
		mazeData:   mazedata.NewMazeDataBlank(),

		titleInfoLabel:    material.H6(theme.Material(), "Maze"),
		titleInputLabel:   material.H6(theme.Material(), "Input"),
		titleControlLabel: material.H6(theme.Material(), "Control"),

		mazeFileName: "No file selected",
	}
	return v
}

func (v *View) SetOnMazeOpen(f func()) {
	v.onMazeOpen = f
}

func (v *View) SetOnAllReset(f func()) {
	v.onAllReset = f
}

func (v *View) SetOnZoomPlus(f func()) {
	v.onZoomPlus = f
}

func (v *View) SetOnZoomMinus(f func()) {
	v.onZoomMinus = f
}

func (v *View) SetOnZoomReset(f func()) {
	v.onZoomReset = f
}

func (v *View) SetMazeData(mazeData *mazedata.MazeData) {
	v.mazeData = mazeData
}

func (v *View) Layout(gtx C, theme *hakoniwatheme.Theme) D {
	return v.split.Layout(gtx, theme,
		func(gtx C) D {
			return v.layoutLeft(gtx, theme)
		},
		func(gtx C) D {
			return v.layoutRight(gtx, theme)
		},
	)
}

func (v *View) layoutLeft(gtx C, theme *hakoniwatheme.Theme) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return v.layoutMaze(gtx, theme)
		}),
		layout.Rigid(func(gtx C) D {
			return v.layoutSeekbar(gtx, theme)
		}),
	)
}

func (v *View) layoutRight(gtx C, theme *hakoniwatheme.Theme) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return v.layoutInfo(gtx, theme)
		}),
		layout.Rigid(func(gtx C) D {
			return v.layoutInput(gtx, theme)
		}),
		layout.Rigid(func(gtx C) D {
			return v.layoutControl(gtx, theme)
		}),
	)
}

func (v *View) layoutMaze(gtx C, theme *hakoniwatheme.Theme) D {
	return layout.Inset{}.Layout(gtx,
		func(gtx C) D {
			// return material.H1(theme.Material(), "Maze").Layout(gtx)
			return v.mazeWidget.Layout(gtx, theme, v.mazeData)
		},
	)
}

func (v *View) layoutSeekbar(gtx C, theme *hakoniwatheme.Theme) D {
	return layout.Inset{Top: unit.Dp(2)}.Layout(gtx,
		func(gtx C) D {
			return material.Slider(theme.Material(), &widget.Float{
				Value: 0.5,
			}).Layout(gtx)
		},
	)
}

func (v *View) layoutInfo(gtx C, _ *hakoniwatheme.Theme) D {
	return layout.Inset{}.Layout(gtx,
		func(gtx C) D {
			return v.titleInfoLabel.Layout(gtx)
		},
	)
}

func (v *View) layoutInput(gtx C, _ *hakoniwatheme.Theme) D {
	return layout.Inset{}.Layout(gtx,
		func(gtx C) D {
			return v.titleInputLabel.Layout(gtx)
		},
	)
}

func (v *View) layoutControl(gtx C, theme *hakoniwatheme.Theme) D {
	if v.onMazeOpen != nil {
		if v.mazeOpenButton.Clicked(gtx) {
			v.onMazeOpen()
		}
	}

	return layout.Inset{
		Top:    unit.Dp(2),
		Left:   unit.Dp(2),
		Right:  unit.Dp(2),
		Bottom: unit.Dp(2),
	}.Layout(gtx,
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return v.titleControlLabel.Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
								layout.Rigid(func(gtx C) D {
									return material.Body1(theme.Material(), v.mazeFileName).Layout(gtx)
								}),
							)
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
						layout.Rigid(func(gtx C) D {
							return material.Button(theme.Material(), &v.mazeOpenButton, "Open").Layout(gtx)
						}),
					)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),
				layout.Rigid(func(gtx C) D {
					if v.allResetButton.Clicked(gtx) {
						v.onAllReset()
					}
					return material.Button(theme.Material(), &v.allResetButton, "Reset").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),
				layout.Rigid(func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							if v.mazeZoomPlusButton.Clicked(gtx) {
								v.onZoomPlus()
							}
							return material.Button(theme.Material(), &v.mazeZoomPlusButton, "+").Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
						layout.Rigid(func(gtx C) D {
							if v.mazeZoomMinusButton.Clicked(gtx) {
								v.onZoomMinus()
							}
							return material.Button(theme.Material(), &v.mazeZoomMinusButton, "-").Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
						layout.Rigid(func(gtx C) D {
							if v.mazeZoomResetButton.Clicked(gtx) {
								v.onZoomReset()
							}
							return material.Button(theme.Material(), &v.mazeZoomResetButton, "0").Layout(gtx)
						}),
					)
				}),
			)
		},
	)
}
