package huecontroller

import (
	"errors"

	"github.com/amimof/huego"
)

// Groups just wraps huego.bridge.GetLights()
func (ctrl *Controller) Groups() ([]huego.Group, error) {
	return ctrl.bridge.GetGroups()
}

// GroupsFiltered calls huego.bridge.GeGroups() and drops any groups contained in groupFilter, before returning.
func (ctrl *Controller) GroupsFiltered(groupFilter []int) ([]huego.Group, error) {
	var result []huego.Group
	groups, err := ctrl.bridge.GetGroups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if getSliceIndex(groupFilter, group.ID) > -1 {
			continue
		}
		result = append(result, group)
	}
	return result, nil
}

// GroupByID returns *huego.Group on success, raises error otherwise
func (ctrl *Controller) GroupByID(id int) (*huego.Group, error) {
	groups, err := ctrl.bridge.GetGroups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group.ID == id {
			return &group, nil
		}
	}
	return nil, errors.New("group not found - check hue-group setting in ~/.huego-fe.yml")
}

// GroupPowerOff powers the given groupID off
func (ctrl *Controller) GroupPowerOff(groupID int) error {
	group, err := ctrl.bridge.GetGroup(groupID)
	if err != nil {
		return err
	}
	return group.Off()
}

// GroupPowerOn powers the given groupID on
func (ctrl *Controller) GroupPowerOn(groupID int) error {
	group, err := ctrl.bridge.GetGroup(groupID)
	if err != nil {
		return err
	}
	return group.On()
}

// GroupSetBrightness controls brightness on given groupID
func (ctrl *Controller) GroupSetBrightness(groupID int, brightness uint8) error {
	group, err := ctrl.bridge.GetGroup(groupID)
	if err != nil {
		return err
	}
	return group.Bri(brightness)
}

// GroupSetColorTemperature controls color temperature on given groupID
func (ctrl *Controller) GroupSetColorTemperature(groupID int, colorTemperature uint16) error {
	group, err := ctrl.bridge.GetGroup(groupID)
	if err != nil {
		return err
	}
	return group.Ct(colorTemperature)
}
