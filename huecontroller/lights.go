package huecontroller

import (
	"errors"

	"github.com/amimof/huego"
)

// Lights just wraps huego.bridge.GetLights()
func (ctrl *Controller) Lights() ([]huego.Light, error) {
	return ctrl.bridge.GetLights()
}

// LightsFiltered calls huego.bridge.GetLights() and drops any lights contained in lightFilter, before returning.
func (ctrl *Controller) LightsFiltered(lightFilter []int) ([]huego.Light, error) {
	var result []huego.Light
	lights, err := ctrl.bridge.GetLights()
	if err != nil {
		return nil, err
	}
	for _, light := range lights {
		if getSliceIndex(lightFilter, light.ID) > -1 {
			continue
		}
		result = append(result, light)
	}
	return result, nil
}

// LightByID returns *huego.Light on success, raises error otherwise
func (ctrl *Controller) LightByID(id int) (*huego.Light, error) {
	lights, err := ctrl.bridge.GetLights()
	if err != nil {
		return nil, err
	}
	for _, light := range lights {
		if light.ID == id {
			return &light, nil
		}
	}
	return nil, errors.New("light not found - check hue-light setting in ~/.huego-fe.yml")
}

// PowerOff powers the given lightID off
func (ctrl *Controller) PowerOff(lightID int) error {
	light, err := ctrl.bridge.GetLight(lightID)
	if err != nil {
		return err
	}
	return light.Off()
}

// PowerOn powers the given lightID on
func (ctrl *Controller) PowerOn(lightID int) error {
	light, err := ctrl.bridge.GetLight(lightID)
	if err != nil {
		return err
	}
	return light.On()
}

// SetBrightness controls brightness on given lightID
func (ctrl *Controller) SetBrightness(lightID int, brightness uint8) error {
	light, err := ctrl.bridge.GetLight(lightID)
	if err != nil {
		return err
	}
	return light.Bri(brightness)
}

// SetColorTemperature controls color temperature on given lightID
func (ctrl *Controller) SetColorTemperature(lightID int, colorTemperature uint16) error {
	light, err := ctrl.bridge.GetLight(lightID)
	if err != nil {
		return err
	}
	return light.Ct(colorTemperature)
}
