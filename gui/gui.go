package gui

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

type iconAndTextButton struct {
	theme  *material.Theme
	button *widget.Clickable
	icon   *widget.Icon
	word   string
}

func (b iconAndTextButton) Layout(gtx layout.Context) layout.Dimensions {
	return material.ButtonLayout(b.theme, b.button).Layout(gtx, func(gtx C) D {
		return layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx C) D {
			iconAndLabel := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
			textIconSpacer := unit.Dp(5)

			layIcon := layout.Rigid(func(gtx C) D {
				return layout.Inset{Right: textIconSpacer}.Layout(gtx, func(gtx C) D {
					var d D
					if icon != nil {
						size := gtx.Px(unit.Dp(56)) - 2*gtx.Px(unit.Dp(16))
						b.icon.Layout(gtx, unit.Px(float32(size)))
						d = layout.Dimensions{
							Size: image.Point{X: size, Y: size},
						}
					}
					return d
				})
			})

			layLabel := layout.Rigid(func(gtx C) D {
				return layout.Inset{Left: textIconSpacer}.Layout(gtx, func(gtx C) D {
					l := material.Body1(b.theme, b.word)
					l.Color = b.theme.Color.InvText
					return l.Layout(gtx)
				})
			})

			return iconAndLabel.Layout(gtx, layIcon, layLabel)
		})
	})
}

func kitchen(gtx layout.Context, th *material.Theme) layout.Dimensions {
	for _, e := range lineEditor.Events() {
		if e, ok := e.(widget.SubmitEvent); ok {
			topLabel = e.Text
			lineEditor.SetText("")
		}
	}
	widgets := []layout.Widget{
		material.H5(th, topLabel).Layout,
		func(gtx C) D {
			e := material.Editor(th, lineEditor, "Hint")
			e.Font.Style = text.Italic
			border := widget.Border{Color: color.RGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
			})
		},
		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, material.IconButton(th, iconButton, icon).Layout)
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, iconAndTextButton{theme: th, icon: icon, word: "Icon", button: iconTextButton}.Layout)
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for button.Clicked() {
							green = !green
						}
						return material.Button(th, button, "Click me!").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						l := "Green"
						if !green {
							l = "Blue"
						}
						btn := material.Button(th, greenButton, l)
						if green {
							btn.Background = color.RGBA{A: 0xff, R: 0x9e, G: 0x9d, B: 0x24}
						}
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						return material.Clickable(gtx, flatBtn, func(gtx C) D {
							return layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx C) D {
								flatBtnText := material.Body1(th, "Flat")
								if gtx.Queue == nil {
									flatBtnText.Color.A = 150
								}
								return layout.Center.Layout(gtx, flatBtnText.Layout)
							})
						})
					})
				}),
			)
		},
		material.ProgressBar(th, progress).Layout,
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(
					material.CheckBox(th, checkbox, "Transform").Layout,
				),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Left: unit.Dp(16)}.Layout(gtx,
						material.Switch(th, swtch).Layout,
					)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Left: unit.Dp(16)}.Layout(gtx, func(gtx C) D {
						text := "enabled"
						if !swtch.Value {
							text = "disabled"
							gtx = gtx.Disabled()
						}
						btn := material.Button(th, disableBtn, text)
						return btn.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Left: unit.Dp(16)}.Layout(gtx, func(gtx C) D {
						if !swtch.Value {
							return D{}
						}
						return material.Loader(th).Layout(gtx)
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(material.RadioButton(th, radioButtonsGroup, "r1", "RadioButton1").Layout),
				layout.Rigid(material.RadioButton(th, radioButtonsGroup, "r2", "RadioButton2").Layout),
				layout.Rigid(material.RadioButton(th, radioButtonsGroup, "r3", "RadioButton3").Layout),
			)
		},
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, material.Slider(th, float, 0, 2*math.Pi).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx,
						material.Body1(th, fmt.Sprintf("%.2f", float.Value)).Layout,
					)
				}),
			)
		},
	}

	return list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}
