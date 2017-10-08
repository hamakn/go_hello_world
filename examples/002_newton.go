package examples

import "math"

func MySqrt(x float64) float64 {
	z := 1.0
	//for i := 0; i < 10; i++ {
	//    z = z - (z * z - x) / (2 * x)
	//}
	z_old := 0.0
	for {
		z = z - (z*z-x)/(2*x)
		diff := math.Abs(z - z_old)
		if diff < 0.00000001 {
			break
		}
		z_old = z
	}
	return z
}
