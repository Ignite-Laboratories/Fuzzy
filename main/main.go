package main

import (
	"fmt"
	"github.com/ignite-laboratories/arwen/alpha"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
)

var framerate = 60.0 //Hz

func main() {
	core.Verbose = true

	source := []byte{50, 60, 70, 75, 70, 60, 50, 40, 35, 40}
	approximation := []byte{50, 75, 50, 40}
	delta := alpha.CreateDeltaWave(source, approximation)
	unsigned := alpha.UnsignDeltaWave(delta)
	fmt.Println(delta)
	fmt.Println(unsigned)

	view := viewport.NewStackedByteWave(core.Impulse, false, when.Frequency(&framerate), "Stacked Byte Waves", nil, nil, std.RGBFromHex(0x444444, 0xff))

	view.AddBytes(std.RGBA[byte]{R: 255}, source)
	view.AddBytes(std.RGBA[byte]{G: 255}, approximation)
	view.AddBytes(std.RGBA[byte]{B: 255}, unsigned)

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
