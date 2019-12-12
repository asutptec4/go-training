package tasks

import (
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
	return 0
}
