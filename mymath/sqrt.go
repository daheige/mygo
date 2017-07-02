package mymath

func Sqrt(x float64) float64 {
	z := 0.0 //float64
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func GetUser(x string) string {
	return x
}
