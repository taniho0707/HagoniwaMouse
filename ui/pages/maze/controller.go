package maze

import (
	"fmt"
	"path/filepath"

	"github.com/taniho0707/HagoniwaMouse/internal/mazedata"
	"github.com/taniho0707/HagoniwaMouse/ui/explorer"
)

type Controller struct {
	view *View

	// repo repository.Repository

	explorer *explorer.Explorer
}

func NewController(view *View, explorer *explorer.Explorer) *Controller {
	c := &Controller{
		view:     view,
		explorer: explorer,
	}

	view.SetOnMazeOpen(c.onMazeOpen)
	view.SetOnAllReset(c.onAllReset)
	view.SetOnZoomPlus(c.onZoomPlus)
	view.SetOnZoomMinus(c.onZoomMinus)
	view.SetOnZoomReset(c.onZoomReset)

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
		c.view.SetMazeData(mazeData)
	}, "maze")
}

func (c *Controller) onAllReset() {
	// c.view.AddLog(str)
}

func (c *Controller) onZoomPlus() {
	// c.view.AddLog(str)
}

func (c *Controller) onZoomMinus() {
	// c.view.AddLog(str)
}

func (c *Controller) onZoomReset() {
	// c.view.AddLog(str)
}

// func (c *Controller) SetChannels(w *app.Window, logCh chan string) {
// 	go func() {
// 		for log := range logCh {
// 			c.onGetNewLog(log)
// 			w.Invalidate()
// 		}
// 	}()
// }
