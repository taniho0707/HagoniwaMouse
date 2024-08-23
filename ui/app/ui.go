package app

import (
	"image"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/taniho0707/HagoniwaMouse/ui/hakoniwatheme"
	"github.com/taniho0707/HagoniwaMouse/ui/pages/logs"
)

type UI struct {
	Theme  *hakoniwatheme.Theme
	window *app.Window

	sidebar *Sidebar
	footer  *Footer

	// mazeView *maze.View
	logView *logs.View

	// mazeController *maze.Controller
	logController *logs.Controller

	// mazeState *state.Maze

	// repo repository.Repository
}

// New creates a new UI using the Go Fonts.
func New(w *app.Window) (*UI, error) {
	u := &UI{
		window: w,
	}

	// fontCollection, err := fonts.Prepare()
	// if err != nil {
	// 	return nil, err
	// }

	// repo, err := repository.NewFilesystem()
	// if err != nil {
	// 	return nil, err
	// }

	// explorerController := explorer.NewExplorer(w)

	// u.repo = repo

	// u.workspacesView = workspaces.NewView()
	// u.workspacesState = state.NewWorkspaces(repo)
	// u.workspacesController = workspaces.NewController(u.workspacesView, u.workspacesState, repo)
	// if err := u.workspacesController.LoadData(); err != nil {
	// 	return nil, err
	// }

	// grpcService := grpc.NewService(u.requestsState, u.environmentsState, u.protoFilesState)
	// restService := rest.New(u.requestsState, u.environmentsState)

	theme := material.NewTheme()
	// theme.Shaper = text.NewShaper(text.WithCollection(fontCollection))
	// // lest assume is dark theme, we will switch it later
	u.Theme = hakoniwatheme.New(theme, true)
	// // console need to be initialized before other pages as its listening for logs
	// u.consolePage = console.New()

	u.sidebar = NewSidebar(u.Theme)

	u.footer = NewFooter(u.Theme)

	u.logView = logs.NewView(w, u.Theme)
	u.logController = logs.NewController(u.logView)

	// u.header = NewHeader(u.environmentsState, u.workspacesState, u.Theme)
	// u.header.LoadWorkspaces(u.workspacesState.GetWorkspaces())

	// u.header.OnSelectedEnvChanged = func(env *domain.Environment) {
	// 	preferences, err := u.repo.ReadPreferencesData()
	// 	if err != nil {
	// 		fmt.Println("failed to read preferences: ", err)
	// 		return
	// 	}

	// 	preferences.Spec.SelectedEnvironment.ID = env.MetaData.ID
	// 	preferences.Spec.SelectedEnvironment.Name = env.MetaData.Name
	// 	if err := repo.UpdatePreferences(preferences); err != nil {
	// 		fmt.Println("failed to update preferences: ", err)
	// 	}

	// 	u.environmentsState.SetActiveEnvironment(env)
	// }

	return u, u.load()
}

func (u *UI) load() error {
	// preferences, err := u.repo.ReadPreferencesData()
	// if err != nil {
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		preferences = domain.NewPreferences()
	// 		if err := u.repo.UpdatePreferences(preferences); err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		return err
	// 	}
	// }

	// config, err := u.repo.GetConfig()
	// if err != nil {
	// 	return err
	// }

	// u.header.SetTheme(preferences.Spec.DarkMode)

	// if err := u.environmentsController.LoadData(); err != nil {
	// 	return err
	// }

	// u.header.LoadEnvs(u.environmentsState.GetEnvironments())

	// if selectedEnv := u.environmentsState.GetEnvironment(preferences.Spec.SelectedEnvironment.ID); selectedEnv != nil {
	// 	u.environmentsState.SetActiveEnvironment(selectedEnv)
	// 	u.header.SetSelectedEnvironment(u.environmentsState.GetActiveEnvironment())
	// }

	// if selectedWs := u.workspacesState.GetWorkspace(config.Spec.ActiveWorkspace.ID); selectedWs != nil {
	// 	u.workspacesState.SetActiveWorkspace(selectedWs)
	// 	u.header.SetSelectedWorkspace(u.workspacesState.GetActiveWorkspace())
	// }

	// return u.requestsController.LoadData()

	return nil
}

func (u *UI) SetChannels(logCh chan string) {
	u.logController.SetChannels(u.window, logCh)
}

func (u *UI) Run() error {
	// ops are the operations from the UI
	var ops op.Ops

	for {
		switch e := u.window.Event().(type) {
		// this is sent when the application should re-render.
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// render and handle UI.
			u.Layout(gtx)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)
		// this is sent when the application is closed.
		case app.DestroyEvent:
			return e.Err
		}
	}
}

// Layout displays the main program layout.
func (u *UI) Layout(gtx layout.Context) layout.Dimensions {
	// set the background color
	macro := op.Record(gtx.Ops)
	rect := image.Rectangle{
		Max: image.Point{
			X: gtx.Constraints.Max.X,
			Y: gtx.Constraints.Max.Y,
		},
	}
	paint.FillShape(gtx.Ops, u.Theme.Palette.Bg, clip.Rect(rect).Op())
	background := macro.Stop()

	background.Add(gtx.Ops)
	layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return layout.Flex{Axis: layout.Vertical, Spacing: 0}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: 0}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return u.sidebar.Layout(gtx, u.Theme)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							switch u.sidebar.SelectedIndex() {
							case 0:
								return u.logView.Layout(gtx, u.Theme)
								// case 4:
								//	return u.consolePage.Layout(gtx, u.Theme)
							}
							return layout.Dimensions{}
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// return layout.Dimensions{}
					return u.footer.Layout(gtx, u.Theme)
				}),
			)
		}),
		// layout.Expanded(func(gtx layout.Context) layout.Dimensions {
		// 	if modal.Visible() {
		// 		macro := op.Record(gtx.Ops)
		// 		dims := modal.Layout(gtx, u.Theme.Theme)
		// 		op.Defer(gtx.Ops, macro.Stop())
		// 		return dims
		// 	}

		// 	return notify.NotificationController.Layout(gtx, u.Theme)
		// }),
	)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
