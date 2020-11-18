package gui

import (
	"gioui.org/font/gofont"
	"gioui.org/widget/material"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var (
	lineEditor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	button            = new(widget.Clickable)
	greenButton       = new(widget.Clickable)
	iconTextButton    = new(widget.Clickable)
	iconButton        = new(widget.Clickable)
	flatBtn           = new(widget.Clickable)
	disableBtn        = new(widget.Clickable)
	radioButtonsGroup = new(widget.Enum)
	list              = &layout.List{
		Axis: layout.Vertical,
	}
	progress            = 0
	progressIncrementer chan int
	green               = true
	topLabel            = "gcl "
	icon                *widget.Icon
	checkbox            = new(widget.Bool)
	swtch               = new(widget.Bool)
	transformTime       time.Time
	float               = new(widget.Float)
)

func Main(ctrl *hueController.Controller, appVersion string) {
	topLabel = "gcl " + appVersion
	ic, err := widget.NewIcon(icons.ContentAdd)
	if err != nil {
		log.Fatal(err)
	}
	icon = ic
	progressIncrementer = make(chan int)

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			progressIncrementer <- 1
		}
	}()

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(1024), unit.Dp(768)), app.Title("gcl - Google Cloud Logging UI"))
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
				case "P":
					log.Printf("P pressed with .... %v", e.Modifiers)
					// nice ... but remember that console version should have SAME kbd mapping.
					// todo:
					// - extract code.
					// - key up/dn: select log 1..n / cycle focus
					// - key up/dn + shift: move focused log to top/bottom
					// - 1..n: select log / focus
					// - 1 + modifiers:
					//     - shift:  force reload now
					//     - ctrl: toggle polling
					// - Xn - eXclusively show N
					// - q - edit query for selected log
					// - s - split screen view
					// - t - tab view
					if e.Modifiers.Contain(key.ModShortcut) {
						// ModShortcut is the platform's shortcut modifier, usually the Ctrl
						// key. On Apple platforms it is the Cmd key.
						log.Printf(" P mit ModShortcut")
					}
					if e.Modifiers.Contain(key.ModAlt) {
						log.Printf(" P mit ALT")
					}
					// ...
				}
			case system.ClipboardEvent:
				lineEditor.SetText(e.Text)
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				for iconButton.Clicked() {
					w.WriteClipboard(lineEditor.Text())
				}
				for flatBtn.Clicked() {
					w.ReadClipboard()
				}
				//if *disable {
				//	gtx = gtx.Disabled()
				//}
				if checkbox.Changed() {
					if checkbox.Value {
						transformTime = e.Now
					} else {
						transformTime = time.Time{}
					}
				}

				//transformedKitchen(gtx, th)
				kitchen(gtx, th)
				e.Frame(gtx.Ops)
			}
		case p := <-progressIncrementer:
			progress += p
			if progress > 100 {
				progress = 0
			}
			w.Invalidate()
		}
	}
}
