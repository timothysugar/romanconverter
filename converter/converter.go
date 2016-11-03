package converter

func Convert(n int) string {
	if n < 10 {
		return BuildUnits(n)
	}
	return "X"
}

func BuildUnits(n int) string {
	switch {
	case n == 1:
		return "I"
	case n == 2:
		return "II"
	case n == 3:
		return "III"
	case n == 4:
		return "IV"
	case n == 5:
		return "V"
	case n == 6:
		return "VI"
	case n == 7:
		return "VII"
	case n == 8:
		return "VIII"
	}
	return "IX"
}