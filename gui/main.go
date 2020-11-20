package gui

import (
	"log"
	"os"
	"time"

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
	powerOff             uint8   = iota
	powerOn
	powerToggle
	powerUnknown
)

var (
	selectedLight *huego.Light
	briChan       chan uint8
	pwrChan       chan uint8
	loggedIn      = false
	topLabel      = "huego-fe"
	powerState    = powerUnknown
	buttonOn      = new(widget.Clickable)
	buttonOff     = new(widget.Clickable)
	buttonToggle  = new(widget.Clickable)
	float         = new(widget.Float)
	list          = &layout.List{
		Axis: layout.Vertical,
	}
)

func Main(ctrl *hueController.Controller, appVersion string, selectLight int) {
	topLabel = "huego-fe " + appVersion

	if ctrl.IsLoggedIn() {
		loggedIn = true
		light, err := ctrl.LightById(selectLight) // FEELS BUGGY m(
		if err != nil {
			log.Fatal(err)
		}
		selectedLight = light
		topLabel = selectedLight.Name
		if selectedLight.State.Reachable {
			powerState = powerOff
			if selectedLight.State.On {
				powerState = powerOn
			}
		}
		float.Value = float32(selectedLight.State.Bri)
	} else {
		topLabel = "Not paired yet & UI cannot yet"
	}

	briChan = make(chan uint8, 100) // hack. make general cmd chan??
	go handleBrightnessAction(selectedLight)

	pwrChan = make(chan uint8, 100) // hack. make general cmd chan??
	go handlePowerActions(selectedLight)

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
				// log.Printf("HIT %+v", e.State)
				// Linux gets state 0+1 (pressed+released) while Mac seems to see 0 only...
				// Only process one of the events...
				if e.State != 0 {
					// log.Print("ignoring key event, waiting for release event...")
					continue
				}
				if e.Name == key.NameEscape {
					// always permit Escape, even if not logged in
					os.Exit(0)
				}
				// While unpaired, stuff below will do no good... so:
				if !loggedIn {
					continue
				}
				switch e.Name {
				case key.NameRightArrow:
					float.Value = getSliderValueFor(actionIncrease, float.Value, e.Modifiers)
					briChan <- uint8(float.Value)
				case key.NameLeftArrow:
					float.Value = getSliderValueFor(actionDecrease, float.Value, e.Modifiers)
					briChan <- uint8(float.Value)

				case key.NameUpArrow:
					log.Printf("TODO Up - select next/higher-id lamp")
				case key.NameDownArrow:
					log.Printf("TODO Down - select prev/lower-id lamp")

				case key.NamePageUp:
					fallthrough
				case key.NameHome:
					pwrChan <- powerOn

				case key.NamePageDown:
					fallthrough
				case key.NameEnd:
					pwrChan <- powerOff

				case key.NameReturn:
					fallthrough
				case key.NameEnter:
					pwrChan <- powerToggle

				case "Space": // Mac (+Win?)
					fallthrough
				case " ": // Linux
					//log.Printf("Space pressed - toggling state and saying bye")
					pwrChan <- powerToggle
					go func() {
						// how to wait/ensure command was sent (+successfully?) - wait on feedback chan?
						time.Sleep(250 * time.Millisecond)
						os.Exit(0)
					}()
					//default:
					//	log.Printf("IGNORED: Key '%s'", e.Name) -- also exit() here?
				}

				if e.State == 0 {
					// invalidate after any keypress. not only too much as fired for ignored keys...
					w.Invalidate()
				}

			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				if loggedIn {
					for float.Changed() {
						// log.Printf("user moved slider using mouse to: %f", float.Value)
						briChan <- uint8(float.Value)
					}
				}
				kitchen(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}

func handlePowerActions(light *huego.Light) {
	for newState := range pwrChan {
		switch newState {
		case powerOff:
			powerState = powerOff
			light.Off()
		case powerOn:
			powerState = powerOn
			light.On()
		case powerToggle:
			if light.State.On {
				powerState = powerOff
				light.Off()
			} else {
				powerState = powerOn
				light.On()
			}
		}
	}
}

func handleBrightnessAction(light *huego.Light) {
	for newBrightness := range briChan {
		// seems to be true? tweak brightness and it powers on by default...
		powerState = powerOn
		// put yet-ignored retval on some user feedback chan?
		//log.Printf("Setting brightness %d for %s", newBrightness, l.Name)
		light.Bri(newBrightness)
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
