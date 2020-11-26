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
	"gioui.org/widget/material"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

const (
	ctFloatMaxVal        float32 = 500.0
	briFloatMaxVal       float32 = 255.0
	floatDefaultStepSize float32 = 20.0
	floatCtrlStepSize    float32 = 10.0
)

// Main is called by cmd/root.go if huego-fe is invoked without command line arguments
func Main(ctrl *huecontroller.Controller, selectLight int, selectGroup int, ctrlSingle bool, lightFilter, groupFilter string) {
	a := newApp(nil, ctrl, lightFilter, groupFilter)
	if ctrl.IsLoggedIn() {
		err := a.selectInitialGuiLightAndGroup(selectLight, selectGroup, ctrlSingle)
		if err != nil {
			log.Fatalf("unable to select light %d", selectLight)
		}
		// todo: feedback via gui
		a.loggedIn = true
	} else {
		go a.login(selectLight, selectGroup, ctrlSingle)
	}

	go a.handleControlCommands()

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(400), unit.Dp(230)), app.Title("huego-fe - Hue Control UI"))
		a.w = w
		if err := a.loop(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
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
					if e.Modifiers.Contain(key.ModShift) {
						a.ui.ctFloat.Value = getSliderValueFor(directionUp, a.ui.ctFloat.Value, e.Modifiers, 1.0, ctFloatMaxVal)
						a.ctrlChan <- controlCommand{command: SetColorTemperature, targetValue: uint16(a.ui.ctFloat.Value)}
					} else {
						a.ui.briFloat.Value = getSliderValueFor(directionUp, a.ui.briFloat.Value, e.Modifiers, 1.0, briFloatMaxVal)
						a.ctrlChan <- controlCommand{command: SetBrightness, targetValue: uint16(a.ui.briFloat.Value)}
					}
				case key.NameLeftArrow:
					if e.Modifiers.Contain(key.ModShift) {
						a.ui.ctFloat.Value = getSliderValueFor(directionDown, a.ui.ctFloat.Value, e.Modifiers, 1.0, ctFloatMaxVal)
						a.ctrlChan <- controlCommand{command: SetColorTemperature, targetValue: uint16(a.ui.ctFloat.Value)}
					} else {
						a.ui.briFloat.Value = getSliderValueFor(directionDown, a.ui.briFloat.Value, e.Modifiers, 1.0, briFloatMaxVal)
						a.ctrlChan <- controlCommand{command: SetBrightness, targetValue: uint16(a.ui.briFloat.Value)}
					}

				case key.NameUpArrow:
					if a.ui.controlOneLight {
						a.cycleLight(directionUp)
					} else {
						a.cycleGroup(directionUp)
					}
				case key.NameDownArrow:
					if a.ui.controlOneLight {
						a.cycleLight(directionDown)
					} else {
						a.cycleGroup(directionDown)
					}

				case key.NamePageUp:
					fallthrough
				case key.NameHome:
					a.ctrlChan <- controlCommand{command: PowerOn}

				case key.NameTab:
					a.ui.controlOneLight = !a.ui.controlOneLight

				case key.NamePageDown:
					fallthrough
				case key.NameEnd:
					a.ctrlChan <- controlCommand{command: PowerOff}

				// TODO: Cleanup (key bindings) -- Confusing!
				case key.NameReturn:
					fallthrough
				case key.NameEnter:
					a.ctrlChan <- controlCommand{command: PowerToggle}

				case "Space": // Mac (+Win?)
					fallthrough
				case " ": // Linux
					//log.Printf("Space pressed - toggling state and saying bye")
					a.ctrlChan <- controlCommand{command: PowerToggle}
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
					for a.ui.briFloat.Changed() {
						// log.Printf("user moved slider using mouse to: %f", briFloat.Value)
						a.ctrlChan <- controlCommand{
							command:     SetBrightness,
							targetValue: uint16(a.ui.briFloat.Value),
						}
					}
					for a.ui.ctFloat.Changed() {
						a.ctrlChan <- controlCommand{
							command:     SetColorTemperature,
							targetValue: uint16(a.ui.briFloat.Value),
						}
					}
					a.layoutControlPanel(gtx, th)
				} else {
					a.layoutPairingScreen(gtx, th)
				}
				e.Frame(gtx.Ops)
			}
		}
	}
}

func (a *App) selectInitialGuiLightAndGroup(selectLight int, selectGroup int, ctrlSingle bool) error {
	var err error
	if ctrlSingle {
		a.ui.controlOneLight = true
		err = a.selectLightByID(selectLight, true)
		if selectGroup > 0 {
			a.selectGroupByID(selectGroup, false)
		}
	} else {
		err = a.selectGroupByID(selectGroup, true)
		if selectLight > 0 {
			a.selectLightByID(selectLight, false)
		}
	}
	return err
}

func (a *App) login(selectLight int, selectGroup int, ctrlSingle bool) {
	// TODO: This has zero GUI feedback beyond "please press..." (and dies only via console msg...)
	log.Printf("trying to log in ...")
	for a.loggedIn == false {
		log.Printf("retrying login ... ")
		// bad. copy-paste from cmd/login.go. fixme.
		err := a.ctrl.Login()
		if err == nil {
			perr := a.ctrl.SavePrefs()
			if perr != nil {
				log.Fatalf("pairing success, but unable to save prefs! Error: %s", err)
			}
			lights, err := a.getSortedLampIDs()
			if err != nil {
				log.Fatalf("error during initial lamp listing: %s", err)
			}
			if len(lights) == 0 {
				log.Fatalf("no lights on Hue found?!")
			}
			err = a.selectLightByID(lights[0], true)
			if err != nil {
				log.Fatalf("unable to select light: %s", err)
			}
			log.Printf("login succes! initializing GUI with light / group: %d / %d", selectLight, selectGroup)
			err = a.selectInitialGuiLightAndGroup(selectLight, selectGroup, ctrlSingle)
			if err != nil {
				log.Fatalf("unable to select light %d", selectLight)
			}
			log.Print("all good - enabling control UI ...!")
			a.loggedIn = true
			return
		}
		log.Printf("still no pairing success, sleeping 2 seconds ...")
		time.Sleep(2 * time.Second)
	}
}

func getSliderValueFor(action direction, current float32, modifiers key.Modifiers, min, max float32) float32 {
	change := floatDefaultStepSize // 20?
	if modifiers.Contain(key.ModCtrl) {
		change = floatCtrlStepSize // 10?
	} else if modifiers.Contain(key.ModAlt) {
		change = max // to jump min/max
	}
	var newValue float32
	if action == directionUp {
		newValue = current + change
		if newValue > max {
			newValue = max
		}
	} else {
		newValue = current - change
		if newValue < min {
			newValue = min
		}
	}
	return newValue
}
