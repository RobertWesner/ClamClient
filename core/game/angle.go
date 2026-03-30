package game

type Angle struct {
	Yaw   float64
	Pitch float64
}

func NewAngle(yaw, pitch float64) Angle {
	return Angle{
		Yaw:   yaw,
		Pitch: pitch,
	}
}

func (a Angle) Add(o Angle) Angle {
	return Angle{
		Yaw:   a.Yaw + o.Yaw,
		Pitch: a.Pitch + o.Pitch,
	}
}

func (a Angle) Scale(s float64) Angle {
	return Angle{
		Yaw:   a.Yaw * s,
		Pitch: a.Pitch * s,
	}
}
