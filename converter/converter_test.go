package converter

import "testing"

func TestPrintI(t *testing.T) {
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
	}

	for _, c := range cases {
		res := Convert(c.in)

		if res != c.exp {
			t.Errorf("Convert(%d) Expected=%s Result=%s", c.in, c.exp, res)
		}
	}
}
