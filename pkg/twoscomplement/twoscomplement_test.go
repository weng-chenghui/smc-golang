package twoscomplement_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	. "github.com/weng-chenghui/smc-golang/pkg/twoscomplement"
)

func TestFromInt32(t *testing.T) {
	testcases := []struct {
		name     string
		input    int32
		shouldBe Vec
	}{
		{
			name:     "13 should be 01101",
			input:    13,
			shouldBe: Vec{0, 1, 1, 0, 1},
		},
		{
			name:     "-13 should be 11101",
			input:    -13,
			shouldBe: Vec{1, 1, 1, 0, 1},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			got := FromInt32(tc.input)
			if !cmp.Equal(got, tc.shouldBe) {
				tt.Fatalf("%s", cmp.Diff(got, tc.shouldBe))
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	testcases := []struct {
		name     string
		input    Vec
		shouldBe int32
	}{
		{
			name:     "01101 should be 13",
			input:    Vec{0, 1, 1, 0, 1},
			shouldBe: 13,
		},
		{
			name:     "11101 should be -13",
			input:    Vec{1, 1, 1, 0, 1},
			shouldBe: -13,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			got := ToInt32(tc.input)
			if !cmp.Equal(got, tc.shouldBe) {
				tt.Fatalf("%s", cmp.Diff(got, tc.shouldBe))
			}
		})
	}
}

func TestDotProduct(t *testing.T) {
	testcases := []struct {
		name     string
		lvalue   int32
		rvalue   int32
		shouldBe int32
	}{
		{
			name:     "13 (Vec{0,1,1,0,1}) . -13 Vec{1,1,1,0,1} = should be 3",
			lvalue:   13,
			rvalue:   -13,
			shouldBe: 3,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			rvec := FromInt32(tc.rvalue)
			lvec := FromInt32(tc.lvalue)
			gotFloat := lvec.Dot(rvec)
			got := int32(gotFloat)
			if !cmp.Equal(got, tc.shouldBe) {
				tt.Fatalf("%s", cmp.Diff(got, tc.shouldBe))
			}
		})
	}
}
