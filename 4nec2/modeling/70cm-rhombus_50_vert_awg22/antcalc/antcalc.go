// antcalc project main.go
package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	version string = "0.00.002 alpha"
	//    E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
	//    Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622
	//    Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
	//    SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
	//    SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
	//    SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339
	//    Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
	//    Log2E  = 1 / Ln2
	//    Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
	//    Log10E = 1 / Ln10
	Pi     = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
	Varrad = Pi / 180.0                                                       // (*) degrees to radians, (/) radians to degrees
	C      = 299792458.0                                                      // Speed of light in m/s | // https://oeis.org/A003678
)

var (
	freq   float64
	factor float64
	alpha  float64
)

func Rhombus(freq float64, factor float64, alpha float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	wl := C / 1000000 / freq                    // Wavelenght in m
	u := factor * wl                            // Perimeter in m incl. extension factor
	l := u / 4.0                                // Sidelenght in m
	beta := 180.0 - alpha                       // Angle in degree
	e := 2.0 * l * math.Cos((alpha*Varrad)/2.0) // horizontal diagonal in m
	f := 2.0 * l * math.Sin((alpha*Varrad)/2.0) // vertical diagonal in m
	a := l * l * math.Sin(alpha*Varrad)         // surface area in m²
	ri := l / 2 * math.Sin(alpha*Varrad)        // apothem in m

	return wl, u, l, beta, e, f, a, ri
}

func init() {
	flag.Float64Var(&freq, "freq", 435.0, "Which frequency in MHz should be calculated?")
	flag.Float64Var(&factor, "factor", 1.094842, "Which extension/shortening factor should be used?")
	flag.Float64Var(&alpha, "alpha", 53.00, "Which alpha angle should be used?")

	flag.Parse()
}

func main() {

	wl, u, l, beta, e, f, a, ri := Rhombus(freq, factor, alpha)
	fmt.Printf("Frequency in MHz:              %.5f\n", freq)
	fmt.Printf("Wavelenght in m:               %.5f\n", wl)
	fmt.Printf("Speed of light in Mm/s:        %.6f\n", float64(C/1000000))
	fmt.Printf("Extension/shortening factor:   %.5f\n", factor)
	fmt.Printf("Perimeter in m incl. factor:   %.5f\n", u)
	fmt.Printf("Sidelenght in m:               %.5f\n", l)
	fmt.Printf("Alpha angle in degree:         %.5f\n", alpha)
	fmt.Printf("Beta angle in degree:          %.5f\n", beta)
	fmt.Printf("Horizontal diagonal in m:      %.5f\n", e)
	fmt.Printf("Vertical diagonal in m:        %.5f\n", f)
	fmt.Printf("Surface area in m²:            %.5f\n", a)
	fmt.Printf("Apothem in m:                  %.5f\n", ri)

}
