package app

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
	"github.com/taniho0707/HagoniwaMouse/ui/widgets"
)

type Sidebar struct {
	Theme *hakoniwatheme.Theme

	flatButtons []*widgets.FlatButton
	Buttons     []*SideBarButton
	list        *widget.List

	cache *op.Ops

	clickables []*widget.Clickable

	selectedIndex int
}

type SideBarButton struct {
	Icon *widget.Icon
	Text string
}

func NewSidebar(theme *hakoniwatheme.Theme) *Sidebar {
	s := &Sidebar{
		Theme: theme,
		cache: new(op.Ops),

		Buttons: []*SideBarButton{
			{Icon: widgets.ForwardIcon, Text: "Sim"},
			{Icon: widgets.EditorFormatListNumberedIcon, Text: "Params"},
			// {Icon: widgets.WorkspacesIcon, Text: "Flash"},
			// {Icon: widgets.FileFolderIcon, Text: "Config"},
			{Icon: widgets.EditorShortText, Text: "Log"},
		},
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}

	s.clickables = make([]*widget.Clickable, 0)
	for range s.Buttons {
		s.clickables = append(s.clickables, &widget.Clickable{})
	}

	s.makeButtons(theme)

	return s
}

func (s *Sidebar) makeButtons(theme *hakoniwatheme.Theme) {
	s.flatButtons = make([]*widgets.FlatButton, 0)
	for i, b := range s.Buttons {
		s.flatButtons = append(s.flatButtons, &widgets.FlatButton{
			Icon:              b.Icon,
			Text:              b.Text,
			IconPosition:      widgets.FlatButtonIconTop,
			Clickable:         s.clickables[i],
			SpaceBetween:      unit.Dp(5),
			BackgroundPadding: unit.Dp(1),
			CornerRadius:      5,
			MinWidth:          unit.Dp(40),
			BackgroundColor:   theme.SideBarBgColor,
			TextColor:         theme.SideBarTextColor,
			ContentPadding:    unit.Dp(5),
		})
	}
}

func (s *Sidebar) SelectedIndex() int {
	return s.selectedIndex
}

func (s *Sidebar) Layout(gtx layout.Context, theme *hakoniwatheme.Theme) layout.Dimensions {
	for i, c := range s.clickables {
		for c.Clicked(gtx) {
			s.selectedIndex = i
		}
	}

	return layout.Background{}.Layout(gtx,
		// Background color
		func(gtx layout.Context) layout.Dimensions {
			defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, theme.SideBarBgColor)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
		// Sidebar content
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(2), Right: unit.Dp(2)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return s.list.Layout(gtx, len(s.Buttons), func(gtx layout.Context, i int) layout.Dimensions {
							btn := s.flatButtons[i]
							if s.selectedIndex == i {
								btn.TextColor = theme.SideBarTextColor
							} else {
								btn.TextColor = widgets.Disabled(theme.SideBarTextColor)
							}
							return btn.Layout(gtx, theme)
						})
					}),
					// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// 		return layout.Dimensions{}
					// }),
				)
			})
		},
	)
}
