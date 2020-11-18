package hueController

import (
	"errors"
	"fmt"

	"github.com/amimof/huego"
	"github.com/spf13/viper"
)

type Controller struct {
	bridge     *huego.Bridge
	bridgeIP   string
	bridgeUser string
}

func New(ip string, user string) *Controller {
	return &Controller{
		bridge:     huego.New(ip, user),
		bridgeIP:   ip,
		bridgeUser: user,
	}
}

func (ctrl *Controller) SavePrefs() error {
	const flagHueUser = "hue-user" // UGLY! copy pasta from cmd/root.go -- share elsewhere
	const flagHueIP = "hue-ip"
	viper.Set(flagHueUser, ctrl.bridgeUser)
	viper.Set(flagHueIP, ctrl.bridgeIP)
	return viper.SafeWriteConfig()
}

func (ctrl *Controller) Login() error {
	bridge, err := huego.Discover()
	if err != nil {
		return err
	}
	user, err := bridge.CreateUser("huego-fe")
	if err != nil {
		return err
	}
	ctrl.bridgeIP = bridge.Host
	ctrl.bridgeUser = user
	ctrl.bridge = huego.New(ctrl.bridgeIP, ctrl.bridgeUser)
	return nil
}

func (ctrl *Controller) IsLoggedIn() bool {
	return ctrl.bridgeIP != "" && ctrl.bridgeUser != ""
}

func (ctrl *Controller) IP() string {
	return ctrl.bridgeIP
}

func (ctrl *Controller) Lights() ([]huego.Light, error) {
	return ctrl.bridge.GetLights()
}

func (ctrl *Controller) LightById(id int) (*huego.Light, error) {
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

func (ctrl *Controller) List() {
	l, err := ctrl.bridge.GetLights()
	if err != nil {
		panic(err)
	}
	listFormat := "%12s: %v\n"
	for n, light := range l {
		if n > 0 {
			println()
		}
		fmt.Printf("%d: %s [%s]\n", light.ID, light.Name, light.ModelID)
		fmt.Printf(listFormat, "PoweredOn", light.State.On)
		fmt.Printf(listFormat, "Brightness", light.State.Bri)
		fmt.Printf(listFormat, "Reachable", light.State.Reachable)
		fmt.Printf(listFormat, "ColorMode", light.State.ColorMode)
	}
}

func (ctrl *Controller) PowerOff(lightId int) error {
	light, err := ctrl.bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.Off()
}

func (ctrl *Controller) PowerOn(lightId int) error {
	light, err := ctrl.bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.On()
}

func (ctrl *Controller) SetBrightness(lightId int, brightness uint8) error {
	light, err := ctrl.bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.Bri(brightness)
}

// SetColor / Temp ...
