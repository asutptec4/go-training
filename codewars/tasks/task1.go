package tasks

import (
	"math/bits"
	"net"
	"strconv"
	"strings"
)

// SplitString Complete the solution so that it splits the string into pairs of two characters.
// If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
func SplitString(str string) []string {
	var r = []string{""}
	var n int = 0
	for i := 0; i < len(str); i++ {
		if len(r[n]) >= 2 {
			r = append(r, "")
			n++
		}
		r[n] += string(str[i])
	}
	if len(r[n]) == 1 {
		r[n] += "_"
	}
	return r
}

// IsValidIP Write an algorithm that will identify valid IPv4 addresses in dot-decimal format.
// IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.
func IsValidIP(ip string) bool {
	octs := strings.Split(ip, ".")
	if len(octs) != 4 {
		return false
	}
	for _, oct := range octs {
		if len(oct) > 1 && oct[0] == '0' {
			return false
		}
		v, err := strconv.Atoi(oct)
		if err != nil {
			return false
		}
		if v < 0 || v > 255 {
			return false
		}
	}
	return true
}

// IsValidIP2 use net package
func IsValidIP2(ip string) bool {
	if r := net.ParseIP(ip); r == nil {
		return false
	}
	return true
}

// CalcBits Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number.
// You can guarantee that input is non-negative.
func CalcBits(input uint) int {
	var r uint = 0
	for input > 0 {
		r += input % 2
		input /= 2
	}
	return int(r)
}

// CalcBits2 use bits package
func CalcBits2(input uint) int {
	return bits.OnesCount(input)
}

// SumTriangleOdd Given the triangle of consecutive odd numbers.
// Calculate the row sums of this triangle from the row index (starting at index 1)
func SumTriangleOdd(n int) int {
	sum := 0
	for index := 1; index <= n; index++ {
		sum += index
	}
	return sumNOddNumbers(sum) - sumNOddNumbers(sum-n)
}

func sumNOddNumbers(n int) int {
	sum := 0
	for index := 0; index < n; index++ {
		sum += 2*index + 1
	}
	return sum
}

// DuplicateCount Write a function that will return the count of distinct case-insensitive alphabetic characters
// and numeric digits that occur more than once in the input string.
// The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.
func DuplicateCount(s1 string) int {
	var res = 0
	s1 = strings.ToLower(s1)
	m := make(map[rune]int)
	for _, ch := range s1 {
		if v, ok := m[ch]; ok {
			m[ch] = v + 1
		} else {
			m[ch] = 1
		}
	}
	for _, v := range m {
		if v > 1 {
			res++
		}
	}
	return res
}

// DNAStrand In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G".
//  You have function with one side of the DNA; you need to get the other complementary side.
func DNAStrand(dna string) string {
	var out []rune = make([]rune, len(dna))
	for i, ch := range dna {
		switch ch {
		case 'A':
			out[i] = 'T'
		case 'T':
			out[i] = 'A'
		case 'G':
			out[i] = 'C'
		case 'C':
			out[i] = 'G'
		}
	}
	return string(out)
}
