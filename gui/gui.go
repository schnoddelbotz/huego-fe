package gui

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// UI holds huego-fe GUI instance UI components
type UI struct {
	buttonOn           *widget.Clickable
	buttonOff          *widget.Clickable
	buttonToggle       *widget.Clickable
	briFloat           *widget.Float // brightness slider
	ctFloat            *widget.Float // color temperature slider
	list               *layout.List
	controlOneLight    bool // .. or a group
	reachableIconMap   map[bool]*widget.Icon
	controlModeIconMap map[bool]*widget.Icon
}

var (
	btnColorMap = map[bool]color.NRGBA{
		true:  {A: 0xcc},
		false: {A: 0x55},
	}
)

type (
	// D is used as shorthand for layout.Dimensions in layout functions below
	D = layout.Dimensions
	// C is used as shorthand for layout.Context in layout functions below
	C = layout.Context
)

func (a *App) controlPanel(gtx layout.Context, th *material.Theme) layout.Dimensions {
	widgets := []layout.Widget{
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Start}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
						return a.ui.controlModeIconMap[a.ui.controlOneLight].Layout(gtx, unit.Dp(24))
					})
				}),
				layout.Flexed(1, func(gtx C) D {
					return layout.UniformInset(unit.Dp(10)).Layout(gtx,
						material.Label(th, unit.Dp(20), a.selectedLight.Name).Layout,
					)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
						return a.ui.reachableIconMap[a.selectedLight.State.Reachable].Layout(gtx, unit.Dp(24))
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				// todo: make slider gray/disabled if lamp is powerd off
				layout.Flexed(0.3, material.Label(th, unit.Dp(14), "  Brightness").Layout),
				layout.Flexed(1, material.Slider(th, a.ui.briFloat, 0, 255).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(4)).Layout(gtx,
						material.Body1(th, fmt.Sprintf("%.0f", a.ui.briFloat.Value)).Layout,
					)
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				// todo: make slider gray/disabled if lamp is powerd off
				layout.Flexed(0.3, material.Label(th, unit.Dp(14), "  ColorTemp").Layout),
				layout.Flexed(1, material.Slider(th, a.ui.ctFloat, 0, 500).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(4)).Layout(gtx,
						material.Body1(th, fmt.Sprintf("%.0f", a.ui.ctFloat.Value)).Layout,
					)
				}),
			)
		},
		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for a.ui.buttonOff.Clicked() {
							a.pwrChan <- powerOff
						}
						btn := material.Button(th, a.ui.buttonOff, "Off")
						btn.Background = btnColorMap[a.selectedLight.State.On]
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for a.ui.buttonOn.Clicked() {
							a.pwrChan <- powerOn
						}
						btn := material.Button(th, a.ui.buttonOn, "On")
						btn.Background = btnColorMap[!a.selectedLight.State.On]
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for a.ui.buttonToggle.Clicked() {
							a.pwrChan <- powerToggle
						}
						buttonText := "Toggle on"
						if a.selectedLight.State.On {
							buttonText = "Toggle off"
						}
						btn := material.Button(th, a.ui.buttonToggle, buttonText)
						btn.Background = btnColorMap[true]
						return btn.Layout(gtx)
					})
				}),
			)
		},
	}
	return a.ui.list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, widgets[i])
	})
}

func (a *App) pairingRequiredScreen(gtx layout.Context, th *material.Theme) layout.Dimensions {
	widgets := []layout.Widget{
		material.Label(th, unit.Dp(20), "Please press Hue's link button").Layout,
		func(gtx C) D {
			return layout.Flex{Alignment: layout.End, Spacing: 2}.Layout(gtx,
				layout.Flexed(0.2, material.Loader(th).Layout),
				layout.Flexed(1,
					material.Label(th, unit.Dp(14), "      press ESC to abort discovery!").Layout,
				),
			)
		},
	}
	return a.ui.list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(25)).Layout(gtx, widgets[i])
	})
}
