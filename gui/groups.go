package gui

import "sort"

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

func (a *App) cycleGroup(direction direction) error {
	groups, err := a.getSortedGroupIDs()
	if err != nil {
		return err
	}
	currentID := a.selectedGroup.ID
	cycleToID := a.selectedGroup.ID
	switch direction {
	case directionUp:
		cycleToID = idHigherThan(a.selectedGroup.ID, groups)
	case directionDown:
		cycleToID = idLowerThan(a.selectedGroup.ID, groups)
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
	groups, err := a.ctrl.GroupsFiltered(a.groupFilter)
	if err != nil {
		return ids, err
	}
	for _, l := range groups {
		ids = append(ids, l.ID)
	}
	sort.Ints(ids)
	return ids, nil
}
