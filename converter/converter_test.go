package converter

import (
"testing"
"bytes"
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
	}

	for _, c := range cases {
		res := Convert(c.in)

		if res != c.exp {
			t.Errorf("Convert(%d) Expected=%s Result=%s", c.in, c.exp, res)
		}
	}
}

func TestBuildTens(t *testing.T) {
	cases := []struct {
		in  int
		exp []byte
	}{
		{1, []byte("")},
		{10, []byte("X")},
		{20, []byte("XX")},
	}

	for _, c := range cases {
		res := BuildTens(c.in)

		if !bytes.Equal(res, c.exp) {
			t.Errorf("BuildTens(%d) Expected=%d Result=%d", c.in, c.exp, res)
		}
	}
}