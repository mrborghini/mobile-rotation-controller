package types

type RotationVector3 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewRotationVector3(RotationX float64, rotationY float64, rotationZ float64, rotationW float64) *RotationVector3 {
	r := RotationVector3{
		X: RotationX,
		Y: rotationY,
		Z: rotationZ,
		W: rotationW,
	}

	return &r
}
