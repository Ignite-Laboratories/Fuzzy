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
	l, r := phrase.BreakMeasuresApart(2)

	h1 := sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Stacked Byte Waves", nil, nil)
	view := viewport.NewStackedByteWave(h1, backgroundColor)
	view.AddBytes(colors.Red, source)
	view.AddBytes(colors.Green, l.AsBytes())
	view.AddBytes(colors.Blue, r.AsBytes())

	recombined := tiny.RecombinePhrases(l, r)
	h2 := sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Recombined", nil, nil)
	viewport.NewBasicByteWave(h2, backgroundColor, colors.Red, recombined.AsBytes())

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}

// Massage performs several operations on the provided data.
//
// If invert is true, every byte is XORd against 11111111.
//
// If derive is true, the derivative of the data is used.
//
// The provided mask is then projected against the data, and the difference from the
// projection to the original data is returned. As this yields a signed value, the data
// is "biased" upward by adding 128 to every value.
func Massage(invert bool, derive bool, mask byte, data []byte) []byte {

}
