package converter

import "bytes"

func Convert(n int) string {
	var buf bytes.Buffer

	buf.Write(DetermineXs(n))
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

func DetermineXs(n int) []byte {
	if n >= 10 {
		return []byte("X")
	}
	return []byte("")
}