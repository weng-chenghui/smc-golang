package twoscomplement

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/quartercastle/vector"
)

type Vec = vector.Vector

func FromInt32(i int32) Vec {
	arr := make([]float64, 0)
	str := fmt.Sprintf("%b", i)
	for pos, char := range str {
		strChar := fmt.Sprintf("%c", char)
		if pos == 0 {
			if strChar == "-" {
				arr = append(arr, 1)
				// in this case, the pos 0 is "-" only.
				continue
			} else {
				// in this case, the pos 0 is bit of the int.
				// so append MSB here.
				arr = append(arr, 0)
			}
		}
		converted, _ := strconv.ParseFloat(strChar, 64)
		arr = append(arr, float64(converted))
	}
	v := vector.Vector(arr)
	return v
}

// Ignore the neg. TODO: confirming about how to deal with MSB neg when dot producting.
// The scalar-product of two integers' bitvectors says nothing about the neg bit.
func Dot(lvec, rvec Vec) int32 {
	r := lvec[1:].Dot(rvec[1:])
	return int32(r)
}

func ToInt32(v Vec) int32 {
	var b strings.Builder
	for pos, bit := range v {
		if pos == 0 && bit == 1 {
			b.WriteString("-")
			continue
		}
		b.WriteString(strconv.Itoa(int(bit)))
	}
	result, _ := strconv.ParseInt(b.String(), 2, 32)
	return int32(result)
}

func IsOverflowed(v Vec) bool {
	return !(v[1] < 2)
}
