// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"flag"
	"fmt"
	"go_training/ch03/ex03/point"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var width = 600          // = flag.Int("w", 600, "canvas width")
var height = 320         // = flag.Int("h", 320, "canvas height")
var color = "ffffff"     // = flag.String("vc", "#ffffff", "vertex color")
var cells int            // number of grid cells
var xyrange float64      // axis ranges (-xyrange..+xyrange)
var xyscale float64      // axis ranges (-xyrange..+xyrange)
var zscale float64       // pixels per z unit
var angle float64        // angle of x, y axes (=30°)4
var sin30, cos30 float64 // sin(30°), cos(30°)

func main() {
	flag.Parse()

	handler := func(w http.ResponseWriter, r *http.Request) {
		if param := r.FormValue("width"); len(param) > 0 {
			if paramNum, err := strconv.Atoi(param); err != nil {
				log.Print(err)
			} else {
				width = paramNum
			}
		}
		if param := r.FormValue("height"); len(param) > 0 {
			if paramNum, err := strconv.Atoi(param); err != nil {
				log.Print(err)
			} else {
				height = paramNum
			}
		}
		if param := r.FormValue("color"); len(param) > 0 {
			color = param
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w, width, height, color)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer, width, height int, vertexColor string) {
	cells = 100
	xyrange = 30.0
	xyscale = float64(width/2) / xyrange
	zscale = float64(height) * 0.4
	angle = math.Pi / 6
	sin30, cos30 = math.Sin(angle), math.Cos(angle)

	//fmt.Printf("%v %v %v %v %v %v %v %v %v", width, height, cells, xyrange, xyscale, zscale, angle, sin30, cos30)

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

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for _, p := range polygons {
		ax, ay := convert2dpoint(p.Points[0])
		bx, by := convert2dpoint(p.Points[1])
		cx, cy := convert2dpoint(p.Points[2])
		dx, dy := convert2dpoint(p.Points[3])

		fill := color
		if p.MaxZ() > maxZ-0.01 {
			fill = "ff0000"
		}
		if p.MinZ() < minZ+0.01 {
			fill = "0000ff"
		}

		fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#%s'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy, fill)
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (point.Point3d, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

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
	sx := float64(width/2) + (p.X-p.Y)*cos30*xyscale
	sy := float64(height/2) + (p.X+p.Y)*sin30*xyscale - (p.Z * zscale)
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// func c2r(c color.RGBA) string {
// 	return fmt.Sprintf("%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
// }

//!-
