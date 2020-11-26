package gui

import (
	"sort"
)

func (a *App) cycleLight(direction direction) error {
	lights, err := a.getSortedLampIDs()
	if err != nil {
		return err
	}
	currentID := a.selectedLight.ID
	cycleToID := a.selectedLight.ID
	switch direction {
	case directionUp:
		cycleToID = idHigherThan(a.selectedLight.ID, lights)
	case directionDown:
		cycleToID = idLowerThan(a.selectedLight.ID, lights)
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

// common util (used by groups, too)

func sliceIndex(haystack []int, needle int) int {
	for index, val := range haystack {
		if val == needle {
			return index
		}
	}
	return -1
}

func idHigherThan(currentID int, lights []int) int {
	currentLightIndex := sliceIndex(lights, currentID)
	if currentLightIndex+1 < len(lights) {
		return lights[currentLightIndex+1]
	}
	return currentID
}

func idLowerThan(currentID int, lights []int) int {
	currentLightIndex := sliceIndex(lights, currentID)
	if currentLightIndex > 0 {
		return lights[currentLightIndex-1]
	}
	return currentID
}
