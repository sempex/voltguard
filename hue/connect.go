package hue

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
)

func Connect(HueUsername string, HueBridgeIp string) *huego.Bridge {
	bridge := huego.New(HueBridgeIp, HueUsername)
	return bridge
}

func LogLights(bridge *huego.Bridge) string {
	lights, err := bridge.GetLights()

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(lights); i++ {
		lights[i].Off()
		fmt.Println("Turned Off Light: ", lights[i])
	}
	return "toggeld all lights :)"
}
