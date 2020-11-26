package huecontroller

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/spf13/viper"
)

// Controller enables Hue bridge communication via huego
type Controller struct {
	bridge     *huego.Bridge
	bridgeIP   string
	bridgeUser string
}

// New returns a Controller, set up to chat with bridge at ip using user provided. If both none, login should occur.
func New(ip string, user string) *Controller {
	return &Controller{
		bridge:     huego.New(ip, user),
		bridgeIP:   ip,
		bridgeUser: user,
	}
}

// SavePrefs calls viper.SafeWriteConfig() to write (successful) login data to ~/.huego-fe.yml for later use.
func (ctrl *Controller) SavePrefs() error {
	const flagHueUser = "hue-user" // UGLY! copy pasta from cmd/root.go -- share elsewhere
	const flagHueIP = "hue-ip"
	viper.Set(flagHueUser, ctrl.bridgeUser)
	viper.Set(flagHueIP, ctrl.bridgeIP)
	return viper.SafeWriteConfig()
}

// Login runs discovery, tries to link, and enables later ctrl.bridge usage upon success
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

// IsLoggedIn only checks instances variables to determine login state.
func (ctrl *Controller) IsLoggedIn() bool {
	return ctrl.bridgeIP != "" && ctrl.bridgeUser != ""
}

// IP returns Hue IP used by Controller instance
func (ctrl *Controller) IP() string {
	return ctrl.bridgeIP
}

// List is used by CLI to dump lights/groups to console
func (ctrl *Controller) List() {
	l, err := ctrl.bridge.GetLights()
	if err != nil {
		panic(err)
	}
	listFormat := "%18s: %v\n"
	println("### LIGHTS ###")
	for n, light := range l {
		if n > 0 {
			println()
		}
		fmt.Printf(listFormat, "ID", light.ID)
		fmt.Printf(listFormat, "Name", light.Name)
		fmt.Printf(listFormat, "ModelID", light.ModelID)
		fmt.Printf(listFormat, "ManufacturerName", light.ManufacturerName)
		fmt.Printf(listFormat, "Type", light.Type)
		fmt.Printf(listFormat, "ColorMode", light.State.ColorMode) // ct ColorTemp xy Color
		fmt.Printf(listFormat, "Reachable", light.State.Reachable)
		fmt.Printf(listFormat, "PoweredOn", light.State.On)
		fmt.Printf(listFormat, "ColorTemp", light.State.Ct)
		fmt.Printf(listFormat, "Brightness", light.State.Bri)
		if light.State.ColorMode == "xy" {
			fmt.Printf(listFormat, "Xy", light.State.Xy)
			fmt.Printf(listFormat, "Hue", light.State.Hue)
			fmt.Printf(listFormat, "Saturation", light.State.Sat)
		}
		fmt.Printf(listFormat, "Alert", light.State.Alert)
		fmt.Printf(listFormat, "Scene", light.State.Scene)
		fmt.Printf(listFormat, "Effect", light.State.Effect)
	}
	g, err := ctrl.bridge.GetGroups()
	if err != nil {
		panic(err)
	}
	println("\n### GROUPS ###\n")
	for x, group := range g {
		if x > 0 {
			println()
		}
		fmt.Printf(listFormat, "ID", group.ID)
		fmt.Printf(listFormat, "Name", group.Name)
		fmt.Printf(listFormat, "Lights", group.Lights)
		fmt.Printf(listFormat, "PoweredOn", group.State.On)
		fmt.Printf(listFormat, "Scene", group.State.Scene)
		fmt.Printf(listFormat, "Class", group.Class)
		fmt.Printf(listFormat, "Effect", group.State.Effect)
		fmt.Printf(listFormat, "ColorMode", group.State.ColorMode)
		fmt.Printf(listFormat, "Brightness", group.State.Bri)
		if group.State.ColorMode == "xy" {
			fmt.Printf(listFormat, "Hue", group.State.Hue)
			fmt.Printf(listFormat, "Xy", group.State.Xy)
			fmt.Printf(listFormat, "Saturation", group.State.Sat)
		}
	}
}

// dup! -> util!
func getSliceIndex(haystack []int, needle int) int {
	for index, val := range haystack {
		if val == needle {
			return index
		}
	}
	return -1
}
