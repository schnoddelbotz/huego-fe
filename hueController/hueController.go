package hueController

import (
	"fmt"

	"github.com/amimof/huego"
)

type controller struct {
	bridge     *huego.Bridge
	bridgeIP   string
	bridgeUser string
}

func New(ip string, user string) *controller {
	return &controller{
		bridge:     huego.New(ip, user),
		bridgeIP:   ip,
		bridgeUser: user,
	}
}

func Login() (string, string, error) {
	bridge, err := huego.Discover()
	if err != nil {
		return "", "", err
	}
	user, err := bridge.CreateUser("huego-fe")
	if err != nil {
		return "", "", err
	}
	return bridge.Host, user, nil
}

func (ctrl *controller) IsLoggedIn() bool {
	return ctrl.bridgeIP != "" && ctrl.bridgeUser != ""
}

func (ctrl *controller) List() {
	l, err := ctrl.bridge.GetLights()
	if err != nil {
		panic(err)
	}
	listFormat := "%12s: %v\n"
	for n, light := range l {
		if n > 0 {
			println()
		}
		fmt.Printf("%d: %s [%s]\n", n, light.Name, light.ModelID)
		fmt.Printf(listFormat, "PoweredOn", light.State.On)
		fmt.Printf(listFormat, "Reachable", light.State.Reachable)
		fmt.Printf(listFormat, "ColorMode", light.State.ColorMode)
	}
}

func (ctrl *controller) PowerOff(lightId int) error {
	light, err := ctrl.bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.Off()
}

func (ctrl *controller) PowerOn(lightId int) error {
	light, err := ctrl.bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.On()
}

func (ctrl *controller) SetBrightness(lightId int, brightness uint8) error {
	light, err := ctrl.bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.Bri(brightness)
}

// SetColor / Temp ...
