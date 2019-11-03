// Antcalc project main.go
package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	version string = "0.00.001 alpha"
	C              = 299.792458 // 'Speed of light in Mm/s
	//	E       float64 = 2.7182818284590452353602874713527                                // Euler's number
	Pi     float64 = 3.14159265358979323846264338327950288419716939937510582097494459 // 1 Million Digits of Pi: https://www.piday.org/million/
	Varrad float64 = Pi / 180.0                                                       // (*) degrees to radians, (/) radians to degrees
)

var (
	freq   float64
	factor float64
	alpha  float64
)

func Rhombus(freq float64, factor float64, alpha float64) (float64, float64, float64, float64, float64, float64) {
	wl := C / freq                              // 'Wavelenght in m
	u := factor * wl                            // 'Perimeter in m incl. extension factor
	l := u / 4.0                                // 'Sidelenght in m
	beta := 180.0 - alpha                       // 'Angle in degree
	e := 2.0 * l * math.Cos((alpha*Varrad)/2.0) // 'horizontal diagonal
	f := 2.0 * l * math.Sin((alpha*Varrad)/2.0) // 'vertical diagonal
	return wl, u, l, beta, e, f
}

func init() {
	flag.Float64Var(&freq, "freq", 435.0, "Which frequency in MHz should be calculated?")
	flag.Float64Var(&factor, "factor", 1.094842, "Which extension/shortening factor should be used?")
	flag.Float64Var(&alpha, "alpha", 53.00, "Which alpha angle should be used?")

	flag.Parse()
}

func main() {

	wl, u, l, beta, e, f := Rhombus(freq, factor, alpha)
	fmt.Printf("Frequency in MHz:              %.5f\n", freq)
	fmt.Printf("Wavelenght in m:               %.5f\n", wl)
	fmt.Printf("Speed of light in Mm/s:        %.6f\n", C)
	fmt.Printf("Extension/shortening factor:   %.5f\n", factor)
	fmt.Printf("Perimeter in m incl. factor:   %.5f\n", u)
	fmt.Printf("Sidelenght in m:               %.5f\n", l)
	fmt.Printf("Alpha angle in degree:         %.5f\n", alpha)
	fmt.Printf("Beta angle in degree:          %.5f\n", beta)
	fmt.Printf("Horizontal diagonal in m:      %.5f\n", e)
	fmt.Printf("Vertical diagonal in m:        %.5f\n", f)
}
