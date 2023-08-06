package hue

import (
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
