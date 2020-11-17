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

func Login() {
	bridge, _ := huego.Discover()
	user, _ := bridge.CreateUser("huego-fe") // Link button needs to be pressed BEFORE -- fixme, notify/wait user
	bridge = bridge.Login(user)
	fmt.Printf("bridge: %v", bridge)
	fmt.Printf("user  : %v", user)
}

func SaveLoginToConfigFile() error {
	return nil
}

func (ctrl *controller) List() {
	bridge := huego.New(ctrl.bridgeIP, ctrl.bridgeUser)
	l, err := bridge.GetLights()
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
	light, _ := ctrl.bridge.GetLight(lightId)
	return light.On()
}

func (ctrl *controller) SetBrightness(lightId int, brightness uint8) error {
	light, _ := ctrl.bridge.GetLight(lightId)
	return light.Bri(brightness)
}

// SetColor / Temp ...
