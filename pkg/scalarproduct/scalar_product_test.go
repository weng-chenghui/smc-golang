package scalarproduct_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	twos "github.com/weng-chenghui/smc-golang/pkg/twoscomplement"

	. "github.com/weng-chenghui/smc-golang/pkg/scalarproduct"
)

func TestScalarProduct(t *testing.T) {
	testcases := []struct {
		name                string
		input_Xa            Vec
		input_Xa_int32      int32
		input_Xb            Vec
		input_Xb_int32      int32
		commodity_Ra        Vec
		commodity_Rb        Vec
		commodity_ra        int32
		input_yb            int32
		shouldBeAliceResult int32
		shouldBeBobResult   int32
	}{
		{
			name:                "Xa = 3 Vec{0,0,0,1,1}; Xb = 2 Vec{0,0,0,1,0}; Ra = 9 Vec{0,1,0,0,1}, Rb = 8 Vec{0,1,0,0,0}; ra = 13; yb = 66; execpted = (3, 2)",
			input_Xa:            Vec{0, 0, 0, 1, 1},
			input_Xa_int32:      3,
			input_Xb:            Vec{0, 0, 0, 1, 0},
			input_Xb_int32:      2,
			commodity_Ra:        Vec{0, 1, 0, 0, 1},
			commodity_Rb:        Vec{0, 1, 0, 0, 0},
			commodity_ra:        13,
			input_yb:            66,
			shouldBeAliceResult: -65,
			shouldBeBobResult:   66,
		},
		{
			name:                "Xa = 3 Vec{0,0,0,1,1}; Xb = 2 Vec{0,0,0,1,0}; Ra = 9 Vec{0,1,0,0,1}, Rb = 8 Vec{0,1,0,0,0}; ra = 13; yb = 2; execpted = (3, 2)",
			input_Xa:            Vec{0, 0, 0, 1, 1},
			input_Xa_int32:      3,
			input_Xb:            Vec{0, 0, 0, 1, 0},
			input_Xb_int32:      2,
			commodity_Ra:        Vec{0, 1, 0, 0, 1},
			commodity_Rb:        Vec{0, 1, 0, 0, 0},
			commodity_ra:        13,
			input_yb:            2,
			shouldBeAliceResult: -1,
			shouldBeBobResult:   2,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			if twos.ToInt32(tc.input_Xa) != tc.input_Xa_int32 {
				tt.Fatalf("wrong input_Xa: %s", cmp.Diff(twos.ToInt32(tc.input_Xa), tc.input_Xa_int32))
			}
			if twos.ToInt32(tc.input_Xb) != tc.input_Xb_int32 {
				tt.Fatalf("wrong input_Xb: %s", cmp.Diff(twos.ToInt32(tc.input_Xb), tc.input_Xb_int32))
			}

			gotAlice, gotBob := Run(tc.input_Xa,
				tc.input_Xb,
				tc.commodity_Ra,
				tc.commodity_Rb,
				tc.commodity_ra,
				tc.input_yb,
			)

			if !cmp.Equal(gotAlice, tc.shouldBeAliceResult) {
				tt.Fatalf("wrong Alice: %s", cmp.Diff(gotAlice, tc.shouldBeAliceResult))
			}
			if !cmp.Equal(gotBob, tc.shouldBeBobResult) {
				tt.Fatalf("wrong Bob: %s", cmp.Diff(gotBob, tc.shouldBeBobResult))
			}

			// In the original paper:
			//
			// 1. Values of Xa, Xb, ra are int greater than or equal to zero.
			// 2. AliceGot: Xa . Xb + v, BobGot: v, AliceGot - BobGot = Xa . Xb
			//    while v = v' - rb, and v' is the generated one like yb` here,
			//    but the sign is changed.
			//
			// Therefore, in following master thesises, they all become:
			//
			// AliceGot + BobGot = Xa . Xb
			//
			// NOT
			//
			// AliceGot - BobGot = Xa . Xb
			//
			shouldBeResult := int32(tc.input_Xa.Dot(tc.input_Xb))
			got := gotAlice + gotBob
			if !cmp.Equal(got, shouldBeResult) {
				tt.Fatalf("wrong result: %s", cmp.Diff(got, shouldBeResult))
			}
		})
	}
}
