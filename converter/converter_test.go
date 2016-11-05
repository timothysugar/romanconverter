package converter

import (
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		in  int
		exp string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{21, "XXI"},
		{29, "XXIX"},
		{30, "XXX"},
		{31, "XXXI"},
		{41, "XLI"},
		{51, "LI"},
		{61, "LXI"},
		{71, "LXXI"},
		{81, "LXXXI"},
		{91, "XCI"},
		{100, "C"},
		{101, "CI"},
		{999, "CMXCIX"},
	}

	for _, c := range cases {
		res := Convert(c.in)

		if res != c.exp {
			t.Errorf("Convert(%d) Expected=%s Result=%s", c.in, c.exp, res)
		}
	}
}
