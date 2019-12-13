package tasks

import (
	"math"
	"sort"
	"strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	r := strings.Compare(string(s), strings.ToUpper(string(s)))
	if r == 0 {
		return true
	}
	return false
}

func MaxRot(n int64) int64 {
	var combs []int64
	digits := convertToDigits(n)
	for i := 0; len(digits)-i > 0; i++ {
		combs = append(combs, convertFromDigits(digits))
		rotPart := digits[:len(digits)-i]
		last := rotPart[len(rotPart)-1]
		for j := len(rotPart) - 1; j > 0; j-- {
			rotPart[j] = rotPart[j-1]
		}
		rotPart[0] = last
	}
	sort.Slice(combs, func(i, j int) bool { return combs[i] > combs[j] })
	return combs[0]
}

func convertToDigits(n int64) []int64 {
	var r []int64
	for n > 0 {
		r = append(r, n%10)
		n = n / 10
	}
	return r
}

func convertFromDigits(a []int64) int64 {
	r := int64(0)
	for i, v := range a {
		r += v * int64(math.Pow10(i))
	}
	return r
}
