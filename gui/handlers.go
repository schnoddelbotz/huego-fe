package gui

// handle GUI single-light & group actions (on App.ctrlChan)

func (a *App) handleControlCommands() {
	for c := range a.ctrlChan {
		switch c.command {
		case PowerOff:
			fallthrough
		case PowerOn:
			fallthrough
		case PowerToggle:
			if a.ui.controlOneLight {
				a.handleLightPowerAction(c)
			} else {
				a.handleGroupPowerAction(c)
			}
		case SetBrightness:
			a.handleBrightnessAction(c)
		case SetColorTemperature:
			a.handleColorTempAction(c)
		}
	}
}

func (a *App) handleLightPowerAction(c controlCommand) {
	switch c.command {
	case PowerOff:
		a.selectedLight.Off()
	case PowerOn:
		a.selectedLight.On()
	case PowerToggle:
		if a.selectedLight.State.On {
			a.selectedLight.Off()
		} else {
			a.selectedLight.On()
		}
	}
	a.w.Invalidate()
}

func (a *App) handleBrightnessAction(c controlCommand) {
	if a.ui.controlOneLight {
		a.selectedLight.Bri(uint8(c.targetValue))
	} else {
		a.selectedGroup.Bri(uint8(c.targetValue))
	}
}

func (a *App) handleColorTempAction(c controlCommand) {
	if a.ui.controlOneLight {
		a.selectedLight.Ct(c.targetValue)
	} else {
		a.selectedGroup.Ct(c.targetValue)
	}
}
