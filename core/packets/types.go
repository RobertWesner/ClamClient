package packets

import "math"

type Uint8angle uint8

func (a Uint8angle) To360() int {
	return int(math.Round(float64(a) * 360.0 / 256.0))
}
