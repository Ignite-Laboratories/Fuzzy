package fuzzy

import "math"

func CreateDeltaWave(source []byte, approximation []byte) []int {
	if len(approximation) > len(source) {
		panic("approximation is more resolute than source data")
	}

	output := make([]int, len(source))
	run := len(source) / (len(approximation) - 1)
	for ai, a := range approximation {
		// If we're the end of the approximation data...
		if ai == len(approximation)-1 {
			// ...calculate the final point's delta
			output[ai*run] = int(a) - int(source[ai*run])
			break
		}

		// Otherwise, interpolate by phasing across the missing steps...
		rise := int(approximation[ai+1]) - int(a)
		slope := float64(rise) / float64(run)
		for phase := 0; phase < run; phase++ {
			si := (ai * run) + phase
			interpolated := int(math.Floor((float64(phase) * slope) + float64(a)))
			output[si] = interpolated - int(source[si])
		}
	}

	return output
}

func UnsignDeltaWave(source []int) []byte {
	var lowest int
	out := make([]byte, len(source))

	for _, v := range source {
		if v < lowest {
			lowest = v
		}
	}

	for i, v := range source {
		result := v + int(math.Abs(float64(lowest)))
		if result > 255 {
			panic("delta wave overflow")
		}

		out[i] = byte(result)
	}
	return out
}
