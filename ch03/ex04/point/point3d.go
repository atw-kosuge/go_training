package point

import "fmt"

// Point3d 3次元位置
type Point3d struct {
	X float64
	Y float64
	Z float64
}

// ToString output string
func (p Point3d) ToString() string {
	return fmt.Sprintf("%g,%g,%g", p.X, p.Y, p.Z)
}
