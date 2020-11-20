package gui

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var (
	colorBtnDisabled = color.RGBA{A: 0x77}
	colorBtnEnabled  = color.RGBA{A: 0xcc}
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func kitchen(gtx layout.Context, th *material.Theme) layout.Dimensions {
	widgets := []layout.Widget{
		material.H6(th, topLabel).Layout,
		func(gtx C) D {
			if !loggedIn {
				// ugliest way to do it here?
				return layout.Dimensions{}
			}
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				// todo: make slider gray/disabled if lamp is powerd off
				layout.Flexed(1, material.Slider(th, float, 0, 255).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx,
						material.Body1(th, fmt.Sprintf("%.0f", float.Value)).Layout,
					)
				}),
			)
		},
		func(gtx C) D {
			if !loggedIn {
				// ...again!
				return layout.Dimensions{}
			}
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for buttonOff.Clicked() {
							pwrChan <- powerOff
						}
						btn := material.Button(th, buttonOff, "Off")
						if powerState == powerOff {
							btn.Background = colorBtnDisabled
						} else if powerState == powerOn {
							btn.Background = colorBtnEnabled
						}
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for buttonOn.Clicked() {
							pwrChan <- powerOn
						}
						btn := material.Button(th, buttonOn, "On")
						if powerState == powerOff {
							btn.Background = colorBtnEnabled
						} else if powerState == powerOn {
							btn.Background = colorBtnDisabled
						}
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for buttonToggle.Clicked() {
							pwrChan <- powerToggle
						}
						buttonText := "Toggle on"
						if powerState == powerOn {
							buttonText = "Toggle off"
						}
						btn := material.Button(th, buttonToggle, buttonText)
						btn.Background = colorBtnEnabled
						return btn.Layout(gtx)
					})
				}),
			)
		},
	}
	return list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}
