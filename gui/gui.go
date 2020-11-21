package gui

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type UI struct {
	buttonOn     *widget.Clickable
	buttonOff    *widget.Clickable
	buttonToggle *widget.Clickable
	float        *widget.Float
	list         *layout.List
}

var (
	colorBtnDisabled = color.RGBA{A: 0x77}
	colorBtnEnabled  = color.RGBA{A: 0xcc}
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func (a *App) kitchen(gtx layout.Context, th *material.Theme) layout.Dimensions {
	widgets := []layout.Widget{
		material.H6(th, a.topLabel).Layout,
		func(gtx C) D {
			if !a.loggedIn {
				// ugliest way to do it here?
				return layout.Dimensions{}
			}
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				// todo: make slider gray/disabled if lamp is powerd off
				layout.Flexed(1, material.Slider(th, a.ui.float, 0, 255).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx,
						material.Body1(th, fmt.Sprintf("%.0f", a.ui.float.Value)).Layout,
					)
				}),
			)
		},
		func(gtx C) D {
			if !a.loggedIn {
				// ...again!
				return layout.Dimensions{}
			}
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for a.ui.buttonOff.Clicked() {
							a.pwrChan <- powerOff
						}
						btn := material.Button(th, a.ui.buttonOff, "Off")
						if a.powerState == powerOff {
							btn.Background = colorBtnDisabled
						} else if a.powerState == powerOn {
							btn.Background = colorBtnEnabled
						}
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for a.ui.buttonOn.Clicked() {
							a.pwrChan <- powerOn
						}
						btn := material.Button(th, a.ui.buttonOn, "On")
						if a.powerState == powerOff {
							btn.Background = colorBtnEnabled
						} else if a.powerState == powerOn {
							btn.Background = colorBtnDisabled
						}
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for a.ui.buttonToggle.Clicked() {
							a.pwrChan <- powerToggle
						}
						buttonText := "Toggle on"
						if a.powerState == powerOn {
							buttonText = "Toggle off"
						}
						btn := material.Button(th, a.ui.buttonToggle, buttonText)
						btn.Background = colorBtnEnabled
						return btn.Layout(gtx)
					})
				}),
			)
		},
	}
	return a.ui.list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}
