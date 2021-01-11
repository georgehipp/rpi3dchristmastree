package ledtopins

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func Pins() ([]*gpio.LedDriver, [][]*gpio.LedDriver, [][]*gpio.LedDriver, *raspi.Adaptor) {

	raspiAdaptor := raspi.NewAdaptor()

	fullTree := map[string]*gpio.LedDriver{
		"led7":  gpio.NewLedDriver(raspiAdaptor, "7"),
		"led8":  gpio.NewLedDriver(raspiAdaptor, "8"),
		"led10": gpio.NewLedDriver(raspiAdaptor, "10"),
		"led11": gpio.NewLedDriver(raspiAdaptor, "11"),
		"led12": gpio.NewLedDriver(raspiAdaptor, "12"),
		"led13": gpio.NewLedDriver(raspiAdaptor, "13"),
		"led15": gpio.NewLedDriver(raspiAdaptor, "15"),
		"led16": gpio.NewLedDriver(raspiAdaptor, "16"),
		"led18": gpio.NewLedDriver(raspiAdaptor, "18"),
		"led19": gpio.NewLedDriver(raspiAdaptor, "19"),
		"led21": gpio.NewLedDriver(raspiAdaptor, "21"),
		"led22": gpio.NewLedDriver(raspiAdaptor, "22"),
		"led23": gpio.NewLedDriver(raspiAdaptor, "23"),
		"led24": gpio.NewLedDriver(raspiAdaptor, "24"),
		"led26": gpio.NewLedDriver(raspiAdaptor, "26"),
		"led29": gpio.NewLedDriver(raspiAdaptor, "29"),
		"led31": gpio.NewLedDriver(raspiAdaptor, "31"),
		"led32": gpio.NewLedDriver(raspiAdaptor, "32"),
		"led33": gpio.NewLedDriver(raspiAdaptor, "33"),
		"led35": gpio.NewLedDriver(raspiAdaptor, "35"),
		"led36": gpio.NewLedDriver(raspiAdaptor, "36"),
		"led37": gpio.NewLedDriver(raspiAdaptor, "37"),
		"led38": gpio.NewLedDriver(raspiAdaptor, "38"),
		"led40": gpio.NewLedDriver(raspiAdaptor, "40"),
	}

	tree := []*gpio.LedDriver{fullTree["led13"], fullTree["led11"], fullTree["led7"], fullTree["led12"], fullTree["led10"], fullTree["led8"], fullTree["led23"], fullTree["led29"], fullTree["led24"], fullTree["led32"], fullTree["led31"], fullTree["led26"], fullTree["led35"], fullTree["led36"], fullTree["led37"], fullTree["led33"], fullTree["led38"], fullTree["led40"], fullTree["led22"], fullTree["led21"], fullTree["led15"], fullTree["led18"], fullTree["led16"], fullTree["led19"]}

	branch1 := []*gpio.LedDriver{fullTree["led13"], fullTree["led11"], fullTree["led7"], fullTree["led12"], fullTree["led10"], fullTree["led8"]}
	branch2 := []*gpio.LedDriver{fullTree["led23"], fullTree["led29"], fullTree["led24"], fullTree["led32"], fullTree["led31"], fullTree["led26"]}
	branch3 := []*gpio.LedDriver{fullTree["led35"], fullTree["led36"], fullTree["led37"], fullTree["led33"], fullTree["led38"], fullTree["led40"]}
	branch4 := []*gpio.LedDriver{fullTree["led22"], fullTree["led21"], fullTree["led15"], fullTree["led18"], fullTree["led16"], fullTree["led19"]}
	branches := [][]*gpio.LedDriver{branch1, branch2, branch3, branch4}

	side1 := []*gpio.LedDriver{fullTree["led24"], fullTree["led31"], fullTree["led26"]}
	side2 := []*gpio.LedDriver{fullTree["led23"], fullTree["led29"], fullTree["led32"]}
	side3 := []*gpio.LedDriver{fullTree["led35"], fullTree["led37"], fullTree["led40"]}
	side4 := []*gpio.LedDriver{fullTree["led36"], fullTree["led33"], fullTree["led38"]}
	side5 := []*gpio.LedDriver{fullTree["led22"], fullTree["led18"], fullTree["led16"]}
	side6 := []*gpio.LedDriver{fullTree["led21"], fullTree["led15"], fullTree["led19"]}
	side7 := []*gpio.LedDriver{fullTree["led11"], fullTree["led7"], fullTree["led8"]}
	side8 := []*gpio.LedDriver{fullTree["led13"], fullTree["led12"], fullTree["led10"]}
	sides := [][]*gpio.LedDriver{side1, side2, side3, side4, side5, side6, side7, side8}

	return tree, branches, sides, raspiAdaptor

}
