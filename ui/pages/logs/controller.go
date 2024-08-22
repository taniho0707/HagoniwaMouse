package logs

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
	c.view.textBox.SetCode("")
}

func (c *Controller) onGetNewLog(str string) {
	c.view.textBox.AppendCodeTail(str)
}

func (c *Controller) SetChannels(logCh chan string) {
	go func() {
		for {
			select {
			case log := <-logCh:
				c.onGetNewLog(log)
			}
		}
	}()
}
