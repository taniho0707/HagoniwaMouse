package maze

import (
	"fmt"
	"path/filepath"

	"github.com/taniho0707/HagoniwaMouse/internal/mazedata"
	"github.com/taniho0707/HagoniwaMouse/ui/explorer"
	"github.com/taniho0707/HagoniwaMouse/ui/widgets"
)

type Controller struct {
	view *View

	zoomrate       widgets.ZoomRate
	zoomcentermode widgets.ZoomCenterMode

	// repo repository.Repository

	explorer *explorer.Explorer
}

func NewController(view *View, explorer *explorer.Explorer) *Controller {
	c := &Controller{
		view:           view,
		explorer:       explorer,
		zoomrate:       widgets.Zoom32,
		zoomcentermode: widgets.ZoomCenterMaze,
	}

	view.SetOnMazeOpen(c.onMazeOpen)
	view.SetOnAllReset(c.onAllReset)
	view.SetOnZoomPlus(c.onZoomPlus)
	view.SetOnZoomMinus(c.onZoomMinus)
	view.SetOnZoomReset(c.onZoomReset)
	view.SetOnMouseCenter(c.onSetMouseCenter)
	view.SetOnMazeCenter(c.onSetMazeCenter)

	return c
}

func (c *Controller) onMazeOpen() {
	c.explorer.ChoseFile(func(result explorer.Result) {
		if result.Error != nil {
			fmt.Println("failed to get file", result.Error)
			return
		}

		fileName := filepath.Base(result.Filepath)
		// fileDir := filepath.Dir(result.Filepath)

		c.view.mazeFileName = fileName

		mazeData, err := mazedata.NewMazeDataFromFile(result.Data)
		if err != nil {
			fmt.Println("invalid maze data", err)
			return
		}
		c.view.MazeData = mazeData
	}, "maze")
}

func (c *Controller) onAllReset() {
	c.view.MazeWidget.SetZoom(widgets.Zoom32)
	c.view.MazeWidget.SetZoomCenter(widgets.ZoomCenterMaze)
	c.zoomrate = widgets.Zoom32
	c.zoomcentermode = widgets.ZoomCenterMaze
}

func (c *Controller) onZoomMinus() {
	switch c.zoomrate {
	case widgets.Zoom32:
		c.zoomrate = widgets.Zoom32
	case widgets.Zoom16:
		c.zoomrate = widgets.Zoom32
	case widgets.Zoom8:
		c.zoomrate = widgets.Zoom16
	case widgets.Zoom4:
		c.zoomrate = widgets.Zoom8
	case widgets.Zoom2:
		c.zoomrate = widgets.Zoom4
	}
	c.view.MazeWidget.SetZoom(c.zoomrate)
}

func (c *Controller) onZoomPlus() {
	switch c.zoomrate {
	case widgets.Zoom32:
		c.zoomrate = widgets.Zoom16
	case widgets.Zoom16:
		c.zoomrate = widgets.Zoom8
	case widgets.Zoom8:
		c.zoomrate = widgets.Zoom4
	case widgets.Zoom4:
		c.zoomrate = widgets.Zoom2
	case widgets.Zoom2:
		c.zoomrate = widgets.Zoom2
	}
	c.view.MazeWidget.SetZoom(c.zoomrate)
}

func (c *Controller) onZoomReset() {
	c.zoomrate = widgets.Zoom32
	c.view.MazeWidget.SetZoom(widgets.Zoom32)
}

func (c *Controller) onSetMouseCenter() {
	c.zoomcentermode = widgets.ZoomCenterMouse
	c.view.MazeWidget.SetZoomCenter(widgets.ZoomCenterMouse)
}

func (c *Controller) onSetMazeCenter() {
	c.zoomcentermode = widgets.ZoomCenterMaze
	c.view.MazeWidget.SetZoomCenter(widgets.ZoomCenterMaze)
}

// func (c *Controller) SetChannels(w *app.Window, logCh chan string) {
// 	go func() {
// 		for log := range logCh {
// 			c.onGetNewLog(log)
// 			w.Invalidate()
// 		}
// 	}()
// }
