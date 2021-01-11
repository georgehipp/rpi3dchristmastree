package christmastree

import (
	"flag"
	"os"
	"time"

	ledtopins "rpi3dchristmastree/internal"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

func ledControl(leds []*gpio.LedDriver, action string) {
	switch do := action; do {
	case "on":
		for _, pin := range leds {
			pin.On()
		}
	case "off":
		for _, pin := range leds {
			pin.Off()
		}
	case "toggle":
		for _, pin := range leds {
			pin.Toggle()
		}
	}
}

func main() {
	blink := flag.Bool("blink", false, "Enable blink effect for full tree. Use speed to adjust blink speed.\n(default false)")
	speed := flag.Int("speed", 500, "Delay for roation steps in milliseconds.\n")
	star := flag.Bool("star", true, "Tree star light on.\n")
	style := flag.String("style", "tree", "Use tree, branch, side or off to select lighting style.\n\nExamples \nRotate through branches at 25ms per step with star blinking:\n./christmastree -star=true -style=branch -speed=25\n\nTurn Off Tree:\n./christmastree -style=off\n\nBlink whole tree with blinking start:\n./christmastree -star=true -style=tree -blink=true -speed=250\n")
	flag.Parse()

	tree, branches, sides, raspiAdaptor := ledtopins.Pins()

	speedConv := time.Duration(*speed)
	starLight := gpio.NewLedDriver(raspiAdaptor, "3")

	work := func() {
		if *star == false {
			starLight.Off()
		} else {
			starLight.On()
		}

		switch lightStyle := *style; lightStyle {
		case "tree":
			if *blink == true {
				for {
					ledControl(tree, "toggle")
					time.Sleep(speedConv * time.Millisecond)
				}
			} else {
				ledControl(tree, "on")
			}
		case "branch":
			for {
				for _, branch := range branches {
					ledControl(branch, "toggle")
					time.Sleep(speedConv * time.Millisecond)
				}
			}
		case "side":
			for {
				for _, side := range sides {
					ledControl(side, "toggle")
					time.Sleep(speedConv * time.Millisecond)
				}
			}
		case "off":
			ledControl(tree, "off")
			starLight.Off()
			os.Exit(1)
		}
	}

	robot := gobot.NewRobot("christmastree",
		work,
	)

	robot.Start()

}
