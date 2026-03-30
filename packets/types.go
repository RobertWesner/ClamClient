package packets

import "math"

type Uin8angle uint8

func (a Uin8angle) To360() int {
	return int(math.Round(float64(a) * 360.0 / 256.0))
}
