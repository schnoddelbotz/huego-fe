package gui

import "sort"

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
		a.powerState = powerOn
		a.selectedLight.Bri(newBrightness)
	}
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
	err = a.selectLightByID(cycleToID)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) selectLightByID(lightID int) error {
	newLight, err := a.ctrl.LightById(lightID)
	if err != nil {
		return nil
	}
	a.selectedLight = newLight
	a.topLabel = a.selectedLight.Name
	a.ui.float.Value = float32(a.selectedLight.State.Bri)
	// FIXME...
	return nil
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
