package packets

import (
	"fmt"
)

type Packet13PlayerPositionAndLook struct {
	X        float64
	Y        float64
	Stance   float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p Packet13PlayerPositionAndLook) Id() uint8 {
	return 0x0D
}

func (p Packet13PlayerPositionAndLook) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.Write(p.X); err != nil {
		return nil, fmt.Errorf("013 write x: %w", err)
	}

	if err = writer.Write(p.Y); err != nil {
		return nil, fmt.Errorf("013 write y: %w", err)
	}

	if err = writer.Write(p.Stance); err != nil {
		return nil, fmt.Errorf("013 write stance: %w", err)
	}

	if err = writer.Write(p.Z); err != nil {
		return nil, fmt.Errorf("013 write z: %w", err)
	}

	if err = writer.Write(p.Yaw); err != nil {
		return nil, fmt.Errorf("013 write yaw: %w", err)
	}

	if err = writer.Write(p.Pitch); err != nil {
		return nil, fmt.Errorf("013 write patch: %w", err)
	}

	if err = writer.Write(p.OnGround); err != nil {
		return nil, fmt.Errorf("013 write onground: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet13PlayerPositionAndLook) Read(reader PacketReader) error {
	var err error

	if p.X, err = reader.Float64(); err != nil {
		return fmt.Errorf("013 read x: %w", err)
	}

	if p.Y, err = reader.Float64(); err != nil {
		return fmt.Errorf("013 read y: %w", err)
	}

	if p.Stance, err = reader.Float64(); err != nil {
		return fmt.Errorf("013 read stance: %w", err)
	}

	if p.Z, err = reader.Float64(); err != nil {
		return fmt.Errorf("013 read z: %w", err)
	}

	if p.Yaw, err = reader.Float32(); err != nil {
		return fmt.Errorf("013 read yaw: %w", err)
	}

	if p.Pitch, err = reader.Float32(); err != nil {
		return fmt.Errorf("013 read pitch: %w", err)
	}

	if p.OnGround, err = reader.Bool(); err != nil {
		return fmt.Errorf("013 read onground: %w", err)
	}

	return nil
}

func NewPacket13PlayerPositionAndLook(
	x float64,
	y float64,
	stance float64,
	z float64,
	yaw float32,
	pitch float32,
	onGround bool,
) Packet13PlayerPositionAndLook {
	return Packet13PlayerPositionAndLook{
		X:        x,
		Y:        y,
		Stance:   stance,
		Z:        z,
		Yaw:      yaw,
		Pitch:    pitch,
		OnGround: onGround,
	}
}
