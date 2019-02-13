package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
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
	//!+http
	handler := func(w http.ResponseWriter, r *http.Request) {
		res := ""
		w.Header().Set("Content-Type", "image/svg+xml")
		res += "<svg xmlns='http://www.w3.org/2000/svg' " +
			"style='stroke: grey; fill: white; stroke-width: 0.7' " +
			"width='" + strconv.Itoa(width) + "' height='" + strconv.Itoa(height) + "'>"
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				res += "<polygon points='" + fmt.Sprintf("%g", ax) + ", " + fmt.Sprintf("%g", ay) + ", " + fmt.Sprintf("%g", bx) + ", " + fmt.Sprintf("%g", by) + ", " + fmt.Sprintf("%g", cx) + ", " + fmt.Sprintf("%g", cy) + ", " + fmt.Sprintf("%g", dx) + ", " + fmt.Sprintf("%g", dy) + "'/>\n"
			}
		}
		res += "</svg>"
		w.Write([]byte(res))
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
