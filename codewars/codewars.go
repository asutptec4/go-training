package main

import (
	"fmt"

	"github.com/asutptec4/go-training/codewars/tasks"
)

func main() {
	fmt.Println("Run Codewars tasks..")
	fmt.Println("String array - ", tasks.SplitString("abcdeasgs"))
	fmt.Println("Is valid IP - ", tasks.IsValidIP("123.456.789.0"))
	fmt.Println("Is valid IP - ", tasks.IsValidIP2("124.056.32.0"))
	fmt.Println("Count of ones - ", tasks.CalcBits(255))
	fmt.Println("Row sum - ", tasks.SumTriangleOdd(3))
	fmt.Println("Duplicate count - ", tasks.DuplicateCount("aabbcde"))
	fmt.Println("Complementary DNA - ", tasks.DNAStrand("ATTGC"))
	fmt.Println("Uppercase string - ", tasks.MyString("HELLO I AM DONALD").IsUpperCase())
	fmt.Println("MaxRot - ", tasks.MaxRot(56789))
}
