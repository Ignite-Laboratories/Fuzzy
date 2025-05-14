package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/colors"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
	"github.com/ignite-laboratories/support/ipsum"
	"github.com/ignite-laboratories/tiny"
)

var framerate = 60.0 //Hz
var backgroundColor = colors.Grey44

func main() {
	core.Verbose = true

	source := ipsum.GenerateBytes()
	phrase := tiny.NewPhrase(source...)
	leftBits, rightBits := phrase.BreakMeasurementsApart(2)

	h1 := sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Stacked Byte Waves", nil, nil)
	view := viewport.NewStackedByteWave(h1, backgroundColor)
	view.AddBytes(colors.SlateBlue, source)
	view.AddBytes(colors.LimeGreen, leftBits.AsBytes())
	view.AddBytes(colors.Fuchsia, rightBits.AsBytes())

	recombined := tiny.RecombinePhrases(leftBits, rightBits)
	h2 := sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Recombined", nil, nil)
	viewport.NewBasicByteWave(h2, backgroundColor, colors.Red, recombined.AsBytes())

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}

// Massage performs several operations on the provided data.
//
// If invert is true, every bit is XORd against 1.
//
// The provided mask is then projected against the data, and the difference from the
// projection to the original data is returned. As this yields a signed value, the data
// is "biased" upward by adding 128 to every value.
func Massage(invert bool, mask tiny.Measurement, data []byte) []byte {
	if len(data) == 0 {
		return data
	}

	if invert {
		for i := range data {
			data[i] ^= 0xFF
		}
	}

	return data
}
