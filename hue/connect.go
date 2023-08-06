package hue

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
)

var (
	HueUsername string
	HueBridgeIp string
)

func Connect() *huego.Bridge {
	bridge := huego.New(HueBridgeIp, HueUsername)
	return bridge
}

func LogLights(bridge *huego.Bridge) {
	lights, err := bridge.GetLights()

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(lights); i++ {
		lights[i].On()
		fmt.Println("Turned Off Light: ", lights[i])
	}
}
