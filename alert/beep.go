package beep

import (
	"os/exec"

	"github.com/gen2brain/dlgs"
)

var on = true

func Beep() {
	if on {
		exec.Command("say", "time").Run()

		dlgs.List("Time is up!", "text goes here", []string{
			"5min extension",
			"15min extension",
			"30min extension",
		})
	}
}

func TurnOn() {
	on = true
}

func TurnOff() {
	on = false
}
