package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/color"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
	"github.com/ignite-laboratories/support/ipsum"
	"github.com/ignite-laboratories/tiny"
)

var framerate = 60.0 //Hz
var backgroundColor = color.Grey44

func main() {
	core.Verbose = true

	source := ipsum.GenerateBytes()
	phrase := tiny.NewPhrase(source...)
	l, r := phrase.BreakMeasuresApart(2)

	view := viewport.NewStackedByteWave(core.Impulse, false, when.Frequency(&framerate), "Stacked Byte Waves", nil, nil, backgroundColor)
	view.AddBytes(color.Red, source)
	view.AddBytes(color.Green, l.AsBytes())
	view.AddBytes(color.Blue, r.AsBytes())

	recombined := tiny.RecombinePhrases(l, r)
	viewport.NewBasicByteWave(core.Impulse, false, when.Frequency(&framerate), "Recombined", nil, nil, backgroundColor, color.Red, recombined.AsBytes())

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
