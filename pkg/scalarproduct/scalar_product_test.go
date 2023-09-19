package scalarproduct_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	. "github.com/weng-chenghui/smc-golang/pkg/scalarproduct"
)

func TestScalarProduct(t *testing.T) {
	testcases := []struct {
		name                string
		input_Xa            Vec
		input_Xb            Vec
		commodity_Ra        Vec
		commodity_Rb        Vec
		commodity_ra        int32
		input_yb            int32
		shouldBeAliceResult int32
		shouldBeBobResult   int32
	}{
		{
			name:                "Xa = Vec{3}; Xb = Vec{2}; Ra = Vec{9}, Rb = Vec{8}; ra = 13; yb = 66; execpted = (3, 2)",
			input_Xa:            Vec{3},
			input_Xb:            Vec{2},
			commodity_Ra:        Vec{9},
			commodity_Rb:        Vec{8},
			commodity_ra:        13,
			input_yb:            66,
			shouldBeAliceResult: -60,
			shouldBeBobResult:   66,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
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
