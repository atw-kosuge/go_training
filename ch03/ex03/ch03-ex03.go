// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"go_training/ch03/ex03/point"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	polygons := []point.Polygon{}
	maxZ := 0.0
	minZ := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			a, aerr := corner(i+1, j)
			b, berr := corner(i, j)
			c, cerr := corner(i, j+1)
			d, derr := corner(i+1, j+1)

			if aerr == nil && berr == nil && cerr == nil && derr == nil {
				polygon := point.Polygon{Points: []point.Point3d{a, b, c, d}}
				maxZ = math.Max(maxZ, polygon.MaxZ())
				minZ = math.Min(minZ, polygon.MinZ())

				// ポリゴンリストに追加
				polygons = append(polygons, polygon)
			}
		}
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for _, p := range polygons {
		ax, ay := convert2dpoint(p.Points[0])
		bx, by := convert2dpoint(p.Points[1])
		cx, cy := convert2dpoint(p.Points[2])
		dx, dy := convert2dpoint(p.Points[3])

		color := "#ffffff"
		if p.MaxZ() > maxZ-0.01 {
			color = "#ff0000"
		}
		if p.MinZ() < minZ+0.01 {
			color = "#0000ff"
		}

		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy, color)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (point.Point3d, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) {
		return point.Point3d{X: 0, Y: 0, Z: 0}, errors.New("f is NaN")
	}
	//fmt.Printf("%g\n", z)
	return point.Point3d{X: x, Y: y, Z: z}, nil
}

func convert2dpoint(p point.Point3d) (float64, float64) {
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (p.X-p.Y)*cos30*xyscale
	sy := height/2 + (p.X+p.Y)*sin30*xyscale - p.Z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func c2r(c color.RGBA) string {
	return fmt.Sprintf("%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
}

//!-
