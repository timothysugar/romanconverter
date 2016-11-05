package converter

import (
	"bytes"
	"math"
)

func Convert(n int) string {
	var buf bytes.Buffer

	buf.Write(BuildHundreds(n))
	buf.Write(BuildTens(n))
	buf.Write(BuildUnits(n))

	return buf.String()
}

func BuildUnits(n int) []byte {
	var u = n % 10

	switch {
	case u == 0:
		return []byte("")
	case u == 1:
		return []byte("I")
	case u == 2:
		return []byte("II")
	case u == 3:
		return []byte("III")
	case u == 4:
		return []byte("IV")
	case u == 5:
		return []byte("V")
	case u == 6:
		return []byte("VI")
	case u == 7:
		return []byte("VII")
	case u == 8:
		return []byte("VIII")
	}
	return []byte("IX")
}

func BuildTens(n int) []byte {
	var t = math.Floor(float64(n%100) / 10)

	switch {
	case t == 0:
		return []byte("")
	case t == 1:
		return []byte("X")
	case t == 2:
		return []byte("XX")
	case t == 3:
		return []byte("XXX")
	case t == 4:
		return []byte("XL")
	case t == 5:
		return []byte("L")
	case t == 6:
		return []byte("LX")
	case t == 7:
		return []byte("LXX")
	case t == 8:
		return []byte("LXXX")
	}

	return []byte("XC")
}

func BuildHundreds(n int) []byte {
	var h = math.Floor(float64(n%1000) / 100)

	switch {
	case h == 0:
		return []byte("")
	case h == 1:
		return []byte("C")
	case h == 2:
		return []byte("CC")
	case h == 3:
		return []byte("CCC")
	case h == 4:
		return []byte("CD")
	case h == 5:
		return []byte("D")
	case h == 6:
		return []byte("DC")
	case h == 7:
		return []byte("DCC")
	case h == 8:
		return []byte("DCCC")
	}

	return []byte("CM")
}
