package gui

import (
	"sort"
)

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

func (a *App) handleGroupPowerAction(c controlCommand) {
	switch c.command {
	case PowerOff:
		a.selectedGroup.Off()
	case PowerOn:
		a.selectedGroup.On()
	case PowerToggle:
		if a.selectedGroup.State.On {
			a.selectedGroup.Off()
		} else {
			a.selectedGroup.On()
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

// (single) lights

func (a *App) cycleLight(op int8) error {
	lights, err := a.getSortedLampIDs()
	if err != nil {
		return err
	}
	currentID := a.selectedLight.ID
	cycleToID := a.selectedLight.ID
	switch op {
	case cycleUp:
		cycleToID = getLightIDHigherThan(a.selectedLight.ID, lights)
	case cycleDown:
		cycleToID = getLightIDLowerThan(a.selectedLight.ID, lights)
	}
	if currentID == cycleToID {
		return nil
	}
	err = a.selectLightByID(cycleToID, true)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) selectLightByID(lightID int, updateView bool) error {
	newLight, err := a.ctrl.LightByID(lightID)
	if err != nil {
		return nil
	}
	a.selectedLight = newLight
	if !updateView {
		return nil
	}
	a.ui.briFloat.Value = float32(a.selectedLight.State.Bri)
	a.ui.ctFloat.Value = float32(a.selectedLight.State.Ct)
	if a.w != nil {
		a.w.Invalidate()
	}
	return nil
}

func (a *App) selectGroupByID(groupID int, updateView bool) error {
	newGroup, err := a.ctrl.GroupByID(groupID)
	if err != nil {
		return nil
	}
	a.selectedGroup = newGroup
	if !updateView {
		return nil
	}
	a.ui.briFloat.Value = float32(a.selectedGroup.State.Bri)
	a.ui.ctFloat.Value = float32(a.selectedGroup.State.Ct)
	if a.w != nil {
		a.w.Invalidate()
	}
	return nil
}

func (a *App) getSortedLampIDs() ([]int, error) {
	var ids []int
	lights, err := a.ctrl.LightsFiltered(a.lightFilter)
	if err != nil {
		return ids, err
	}
	for _, l := range lights {
		ids = append(ids, l.ID)
	}
	sort.Ints(ids)
	return ids, nil
}

// light groups

func (a *App) cycleGroup(op int8) error {
	groups, err := a.getSortedGroupIDs()
	if err != nil {
		return err
	}
	currentID := a.selectedGroup.ID
	cycleToID := a.selectedGroup.ID
	switch op {
	case cycleUp:
		cycleToID = getLightIDHigherThan(a.selectedGroup.ID, groups)
	case cycleDown:
		cycleToID = getLightIDLowerThan(a.selectedGroup.ID, groups)
	}
	if currentID == cycleToID {
		return nil
	}
	err = a.selectGroupByID(cycleToID, true)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) getSortedGroupIDs() ([]int, error) {
	var ids []int
	groups, err := a.ctrl.Groups()
	if err != nil {
		return ids, err
	}
	for _, l := range groups {
		ids = append(ids, l.ID)
	}
	sort.Ints(ids)
	return ids, nil
}

// common util

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
