package maze

import (
	"fmt"
	"path/filepath"

	"gioui.org/app"

	"github.com/taniho0707/HagoniwaMouse/internal/mazedata"
	udp_domain "github.com/taniho0707/HagoniwaMouse/server/domain"
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

// return: needUpdate bool, udpData udp_server.UdpCommand
func (c *Controller) onGetNewUdpData(udpData udp_domain.UdpCommand) (bool, udp_domain.UdpCommand) {
	switch udpData.Code {
	case udp_domain.CommandSetMousePosition:
		c.view.MazeWidget.SetMousePos(widgets.Position{X: udpData.MousePositionX, Y: udpData.MousePositionY})
		c.view.MazeWidget.SetMouseAngle(udpData.MousePositionAngle)
		return true, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandGetMouseWallsensorValue:
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandResultSuccess} // FIXME: return simulated value
	case udp_domain.CommandGetMouseImuValue:
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandResultSuccess} // FIXME: return simulated value
	case udp_domain.CommandGetMouseBatteryValue:
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandResultSuccess} // FIXME: return simulated value
	case udp_domain.CommandGetMouseEncoderValue:
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandResultSuccess} // FIXME: return simulated value
	case udp_domain.CommandSetMaze:
		// c.view.MazeData = mazedata.NewMazeDataFromUdpCommand(udpData.MazeName) // FIXME: load maze data
		return true, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandSetMouseModel:
		// TODO: set mouse model
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandSetMouseWallsensorType:
		// TODO: set wallsensor type
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandSetMouseWallsensorNum:
		// TODO: set wallsensor num
		return false, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandSetMazeCellMarker:
		// TODO: set maze cell marker
		return true, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandSetPathByCell:
		// TODO: set path by cell
		return true, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandSetPathByPosition:
		// TODO: set path by position
		return true, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	case udp_domain.CommandDeletePathAll:
		// TODO: delete all path
		return true, udp_domain.UdpCommand{Code: udp_domain.CommandInternalNoResponse}
	default:
		return false, udpData
	}
}

func (c *Controller) SetChannels(w *app.Window, udpReceiveCh chan udp_domain.UdpCommand, udpResponseCh chan udp_domain.UdpCommand) {
	go func() {
		for udpData := range udpReceiveCh {
			needUpdate, responseData := c.onGetNewUdpData(udpData)
			udpResponseCh <- responseData
			if needUpdate {
				w.Invalidate()
			}
		}
	}()
}
