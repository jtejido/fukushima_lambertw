package fukushima_lambertw_test

import (
	"github.com/jtejido/fukushima_lambertw"
	"math"
	"strconv"
	"testing"
)

func TestLambertW(t *testing.T) {
	cases := []struct {
		k                 int
		x, expectedValues float64
	}{
		{2, 2.2, math.NaN()},
		{0, 0, 0},
		{0, -3, math.NaN()},
		{0, math.Inf(1), math.Inf(1)},
		{-1, 0, math.Inf(-1)},
		{-1, -3, math.NaN()},
		{-1, 3, math.NaN()},
		{0, -0.33, -0.6032666497551331},
		{0, 0.1, 0.09127652716086226},
		{0, 4.5, 1.2672378143074348},
		{0, 9.9, 1.7391425517333516},
		{0, 10.89, 1.8000374607381258},
		{0, 100.12, 3.386555992882349},
		{-1, -0.36787944117144233, -1},
		{-1, -0.33, -1.541268224332639},
		{-1, -0.1, -3.577152063957297},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := fukushima_lambertw.LambertW(c.k, c.x)

			if math.Abs(res-c.expectedValues) > 1e-10 {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expectedValues, res)
			}

		})
	}
}
