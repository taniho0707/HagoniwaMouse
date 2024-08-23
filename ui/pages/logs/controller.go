package logs

import "gioui.org/app"

type Controller struct {
	view *View

	// repo repository.Repository
}

func NewController(view *View) *Controller {
	c := &Controller{
		view: view,
	}

	view.SetOnClearLog(c.onClearlog)
	view.SetOnGetNewLog(c.onGetNewLog)

	return c
}

func (c *Controller) onClearlog() {
	c.view.text = []string{}
}

func (c *Controller) onGetNewLog(str string) {
	c.view.AddLog(str)
}

func (c *Controller) SetChannels(w *app.Window, logCh chan string) {
	go func() {
		for log := range logCh {
			c.onGetNewLog(log)
			w.Invalidate()
		}
	}()
}
