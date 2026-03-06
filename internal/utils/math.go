package utils

func AbsFloat(f float64) float64 {
	if f < 0 {
		return -f
	} else {
		return f
	}
}
