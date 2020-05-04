package point

import "math"

// Polygon ポリゴン構造体
type Polygon struct {
	Points []Point3d
}

// MaxZ Z値の最大値を取得
func (polygon Polygon) MaxZ() float64 {
	maxZ := 0.0
	for _, p := range polygon.Points {
		maxZ = math.Max(maxZ, p.Z)
	}
	return maxZ
}

// MinZ Z値の最小値を取得
func (polygon Polygon) MinZ() float64 {
	minZ := 0.0
	for _, p := range polygon.Points {
		minZ = math.Max(minZ, p.Z)
	}
	return minZ
}
