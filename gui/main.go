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

const (
	actionDecrease               = 0
	actionIncrease               = 1
	floatMinVal          float32 = 1.0
	floatMaxVal          float32 = 255.0
	floatDefaultStepSize float32 = 20.0
	floatCtrlStepSize    float32 = 10.0
	floatShiftStepSize   float32 = 1.0
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
	briChan       chan uint8
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
		float.Value = float32(selectedLight.State.Bri)
	} else {
		topLabel = "Not paired yet & UI cannot yet"
	}

	briChan = make(chan uint8, 100) // hack. make general cmd chan??
	go func(l huego.Light) {
		for newBrightness := range briChan {
			log.Printf("Setting %d for...", newBrightness)
			selectedLight.Bri(newBrightness)
		}
	}(selectedLight)

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(400), unit.Dp(200)), app.Title("huego-fe - Hue Control UI"))
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
				case key.NameRightArrow:
					// log.Printf("right with [modifiers=%s]. Was: %v", e.Modifiers, float.Value)
					float.Value = getSliderValueFor(actionIncrease, float.Value, e.Modifiers)
					briChan <- uint8(float.Value)
					w.Invalidate()
				case key.NameLeftArrow:
					// log.Printf("left  with [modifiers=%s]. Was: %v", e.Modifiers, float.Value)
					float.Value = getSliderValueFor(actionDecrease, float.Value, e.Modifiers)
					briChan <- uint8(float.Value)
					w.Invalidate()
				case key.NameUpArrow:
					log.Printf("Up - select next/higher-id lamp")
				case key.NameDownArrow:
					log.Printf("Down - select prev/lower-id lamp")
				case key.NameReturn:
					log.Printf("ENTER! Toggle and Quit! (just cant toggle yet...)")
					os.Exit(0)
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

func getSliderValueFor(action int, current float32, modifiers key.Modifiers) float32 {
	change := floatDefaultStepSize // 20?
	if modifiers.Contain(key.ModShift) {
		change = floatShiftStepSize // 1?
	} else if modifiers.Contain(key.ModCtrl) {
		change = floatCtrlStepSize // 10?
	} else if modifiers.Contain(key.ModAlt) {
		change = floatMaxVal // to jump min/max
	}
	newValue := current
	if action == actionIncrease {
		newValue = current + change
		if newValue > floatMaxVal {
			newValue = floatMaxVal
		}
	} else {
		newValue = current - change
		if newValue < floatMinVal {
			newValue = floatMinVal
		}
	}
	return newValue
}
