package main

import (
	"Test-BE/logical"
	"fmt"
)

func main() {

	result1 := logical.ReverseSentence("italem irad irigayaj")
	fmt.Println(result1)

	logical.FizzBuzz(100)

	result3 := logical.Fibonacci(9)
	fmt.Println(result3)

	soal4 := []int{20, 17, 15, 14, 10}
	result4 := logical.FindMinPrice(soal4)
	fmt.Println(result4)

	soal5 := []string{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"}
	result5 := logical.CekListStr(soal5)
	fmt.Println(result5)

}
