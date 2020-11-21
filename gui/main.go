package gui

import (
	"log"
	"os"
	"sort"
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
	cycleLightUp int8 = iota
	cycleLightDown
)

type App struct {
	w    *app.Window
	ui   *UI
	ctrl *hueController.Controller

	selectedLight *huego.Light
	briChan       chan uint8
	pwrChan       chan uint8
	loggedIn      bool
	topLabel      string
	powerState    uint8
}

func newApp(w *app.Window, c *hueController.Controller) *App {
	a := &App{
		w:             w,
		ctrl:          c,
		selectedLight: nil,
		briChan:       make(chan uint8, 100),
		pwrChan:       make(chan uint8, 100),
		loggedIn:      false,
		topLabel:      "huego-fe",
		powerState:    0,
		ui: &UI{
			buttonOn:     new(widget.Clickable),
			buttonOff:    new(widget.Clickable),
			buttonToggle: new(widget.Clickable),
			float:        new(widget.Float),
			list: &layout.List{
				Axis: layout.Vertical,
			},
		},
	}
	return a
}

func Main(ctrl *hueController.Controller, appVersion string, selectLight int) {
	a := newApp(nil, ctrl)
	a.topLabel = "huego-fe " + appVersion

	if ctrl.IsLoggedIn() {
		a.loggedIn = true
		light, err := ctrl.LightById(selectLight) // FEELS BUGGY m(
		if err != nil {
			log.Fatal(err)
		}
		a.selectedLight = light
		a.topLabel = a.selectedLight.Name
		if a.selectedLight.State.Reachable {
			a.powerState = powerOff
			if a.selectedLight.State.On {
				a.powerState = powerOn
			}
		}
		a.ui.float.Value = float32(a.selectedLight.State.Bri)
	} else {
		a.topLabel = "Please press Hue's link button"
	}

	go a.handleBrightnessAction()
	go a.handlePowerActions()

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(400), unit.Dp(200)), app.Title("huego-fe - Hue Control UI"))
		a.w = w
		if err := a.loop(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func (a *App) cycleLight(op int8) error {
	lights, err := a.getSortedLampIDs()
	if err != nil {
		return err
	}
	currentID := a.selectedLight.ID
	cycleToID := a.selectedLight.ID
	switch op {
	case cycleLightUp:
		cycleToID = getLightIDHigherThan(a.selectedLight.ID, lights)
	case cycleLightDown:
		cycleToID = getLightIDLowerThan(a.selectedLight.ID, lights)
	}
	if currentID == cycleToID {
		return nil
	}
	newLight, err := a.ctrl.LightById(cycleToID)
	if err != nil {
		return nil
	}
	// extract: !
	a.selectedLight = newLight
	a.topLabel = a.selectedLight.Name
	a.ui.float.Value = float32(a.selectedLight.State.Bri)
	// FIXME: button state update (on/off status "display")
	return nil
}

func getSliceIndex(haystack []int, needle int) int {
	for index, val := range haystack {
		if val == needle {
			return index
		}
	}
	return -1
}

func getLightIDHigherThan(currentID int, lights []int) int {
	currentLightIndex := getSliceIndex(lights, currentID)
	if currentLightIndex+1 < len(lights) {
		return lights[currentLightIndex+1]
	}
	return currentID
}

func getLightIDLowerThan(currentID int, lights []int) int {
	currentLightIndex := getSliceIndex(lights, currentID)
	if currentLightIndex > 0 {
		return lights[currentLightIndex-1]
	}
	return currentID
}

func (a *App) getSortedLampIDs() ([]int, error) {
	var ids []int
	lights, err := a.ctrl.Lights()
	if err != nil {
		return ids, err
	}
	for _, l := range lights {
		ids = append(ids, l.ID)
	}
	sort.Ints(ids)
	return ids, nil
}

func (a *App) loop() error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		select {
		case e := <-a.w.Events():
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
				if !a.loggedIn {
					continue
				}
				switch e.Name {
				case key.NameRightArrow:
					a.ui.float.Value = getSliderValueFor(actionIncrease, a.ui.float.Value, e.Modifiers)
					a.briChan <- uint8(a.ui.float.Value)
				case key.NameLeftArrow:
					a.ui.float.Value = getSliderValueFor(actionDecrease, a.ui.float.Value, e.Modifiers)
					a.briChan <- uint8(a.ui.float.Value)

				case key.NameUpArrow:
					a.cycleLight(cycleLightUp)
				case key.NameDownArrow:
					a.cycleLight(cycleLightDown)

				case key.NamePageUp:
					fallthrough
				case key.NameHome:
					a.pwrChan <- powerOn

				case key.NamePageDown:
					fallthrough
				case key.NameEnd:
					a.pwrChan <- powerOff

				case key.NameReturn:
					fallthrough
				case key.NameEnter:
					a.pwrChan <- powerToggle

				case "Space": // Mac (+Win?)
					fallthrough
				case " ": // Linux
					//log.Printf("Space pressed - toggling state and saying bye")
					a.pwrChan <- powerToggle
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
					a.w.Invalidate()
				}

			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				if a.loggedIn {
					for a.ui.float.Changed() {
						// log.Printf("user moved slider using mouse to: %f", float.Value)
						a.briChan <- uint8(a.ui.float.Value)
					}
				}
				a.kitchen(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}

func (a *App) handlePowerActions() {
	for newState := range a.pwrChan {
		switch newState {
		case powerOff:
			a.powerState = powerOff
			a.selectedLight.Off()
		case powerOn:
			a.powerState = powerOn
			a.selectedLight.On()
		case powerToggle:
			if a.selectedLight.State.On {
				a.powerState = powerOff
				a.selectedLight.Off()
			} else {
				a.powerState = powerOn
				a.selectedLight.On()
			}
		}
	}
}

func (a *App) handleBrightnessAction() {
	for newBrightness := range a.briChan {
		// seems to be true? tweak brightness and it powers on by default...
		a.powerState = powerOn
		// put yet-ignored retval on some user feedback chan?
		//log.Printf("Setting brightness %d for %s", newBrightness, l.Name)
		a.selectedLight.Bri(newBrightness)
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
