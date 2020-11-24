package gui

import (
	"image/color"
	"strconv"
	"strings"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/amimof/huego"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

// App holds a huego-fe GUI instance's state and provides channels for decoupled UI->hueCtrl communication.
type App struct {
	w    *app.Window
	ui   *UI
	ctrl *huecontroller.Controller

	selectedLight *huego.Light
	briChan       chan uint8
	ctChan        chan uint16
	pwrChan       chan uint8
	loggedIn      bool
	lightFilter   []int
}

func newApp(w *app.Window, c *huecontroller.Controller, lightFilter string) *App {
	a := &App{
		w:       w,
		ctrl:    c,
		briChan: make(chan uint8, 100),
		ctChan:  make(chan uint16, 100),
		pwrChan: make(chan uint8, 100),
		ui: &UI{
			buttonOn:     new(widget.Clickable),
			buttonOff:    new(widget.Clickable),
			buttonToggle: new(widget.Clickable),
			briFloat:     new(widget.Float),
			ctFloat:      new(widget.Float),
			list: &layout.List{
				Axis: layout.Vertical,
			},
		},
	}
	a.ui.reachableIconMap = make(map[bool]*widget.Icon)
	a.ui.reachableIconMap[true], _ = widget.NewIcon(icons.DeviceSignalWiFi4Bar)
	a.ui.reachableIconMap[false], _ = widget.NewIcon(icons.DeviceSignalWiFiOff)
	a.ui.controlModeIconMap = make(map[bool]*widget.Icon)
	a.ui.controlModeIconMap[true], _ = widget.NewIcon(icons.ActionLightbulbOutline)
	a.ui.controlModeIconMap[false], _ = widget.NewIcon(icons.ActionGroupWork)
	a.ui.controlModeIconMap[true].Color = color.NRGBA{A: 100}
	a.ui.controlModeIconMap[false].Color = color.NRGBA{A: 100}
	// ^ tbdL override icon color?
	a.ui.reachableIconMap[true].Color = color.NRGBA{A: 100}
	a.ui.reachableIconMap[false].Color = color.NRGBA{R: 225, A: 0xcc}
	if lightFilter != "" {
		ids := strings.Split(lightFilter, ",")
		for _, sid := range ids {
			id, err := strconv.Atoi(sid)
			if err == nil {
				a.lightFilter = append(a.lightFilter, id)
			}
		}
	}
	return a
}
