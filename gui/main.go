package gui

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/amimof/huego"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var (
	button  = new(widget.Clickable)
	flatBtn = new(widget.Clickable)
	list    = &layout.List{
		Axis: layout.Vertical,
	}
	green         = true
	topLabel      = "huego-fe"
	float         = new(widget.Float)
	selectedLight huego.Light
)

func Main(ctrl *hueController.Controller, appVersion string, selectLight int) {
	topLabel = "huego-fe " + appVersion

	if ctrl.IsLoggedIn() {
		light, err := ctrl.LightById(selectLight) // FEELS BUGGY m(
		if err != nil {
			log.Fatal(err)
		}
		selectedLight = *light
		topLabel = selectedLight.Name
	} else {
		topLabel = "NOT PAIRED YET"
	}

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(400), unit.Dp(250)), app.Title("huego-fe - Hue Control UI"))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case key.Event:
				switch e.Name {
				case key.NameEscape:
					os.Exit(0)
				case key.NameLeftArrow:
					log.Printf("Left -- dec bri ; pressed with: %v", e.Modifiers)
					if e.Modifiers.Contain(key.ModShift) {
						log.Printf(" left with Modshift --> step = step+10 ?")
					}
				case key.NameRightArrow:
					log.Printf("Right -- incr bri")
				case key.NameUpArrow:
					log.Printf("Up - select next/higher-id lamp")
				case key.NameDownArrow:
					log.Printf("Down - select prev/lower-id lamp")
				case key.NameEnter:
					log.Printf("ENTER! Toggle and Quit!")
				}
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				for flatBtn.Clicked() {
					w.ReadClipboard()
				}
				kitchen(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}
