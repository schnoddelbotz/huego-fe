package gui

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/amimof/huego"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

type App struct {
	w    *app.Window
	ui   *UI
	ctrl *hueController.Controller

	selectedLight *huego.Light
	briChan       chan uint8
	pwrChan       chan uint8
	loggedIn      bool
}

func newApp(w *app.Window, c *hueController.Controller) *App {
	a := &App{
		w:       w,
		ctrl:    c,
		briChan: make(chan uint8, 100),
		pwrChan: make(chan uint8, 100),
		ui: &UI{
			buttonOn:     new(widget.Clickable),
			buttonOff:    new(widget.Clickable),
			buttonToggle: new(widget.Clickable),
			reachableIB:  new(widget.Clickable),
			float:        new(widget.Float),
			list: &layout.List{
				Axis: layout.Vertical,
			},
		},
	}
	a.ui.reachableIconMap = make(map[bool]*widget.Icon)
	a.ui.reachableIconMap[true], _ = widget.NewIcon(icons.DeviceSignalWiFi4Bar)
	a.ui.reachableIconMap[false], _ = widget.NewIcon(icons.DeviceSignalWiFiOff)
	// ^ tbdL override icon color?
	return a
}
