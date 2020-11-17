package hue_cmd

import (
	"fmt"

	"github.com/amimof/huego"
)

func Login() {
	bridge, _ := huego.Discover()
	user, _ := bridge.CreateUser("huego-fe") // Link button needs to be pressed BEFORE -- fixme, notify/wait user
	bridge = bridge.Login(user)
	fmt.Printf("bridge: %v", bridge)
	fmt.Printf("user  : %v", user)
}

func List(ip string, user string) {
	bridge := huego.New(ip, user)
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

func Off(ip string, user string, lightId int) error {
	bridge := huego.New(ip, user)
	light, err := bridge.GetLight(lightId)
	if err != nil {
		return err
	}
	return light.Off()
}

func On(ip string, user string, lightId int) error {
	bridge := huego.New(ip, user)
	light, _ := bridge.GetLight(lightId)
	return light.On()
}

func Brightness(ip string, user string, lightId int, brightness uint8) error {
	bridge := huego.New(ip, user)
	light, _ := bridge.GetLight(lightId)
	return light.Bri(brightness)
}
