package logical

import (
	"fmt"
	"strings"
)

// Soal 1
func ReverseSentence(sentence string) string {
	words := strings.Fields(sentence)
	var reversedWords []string

	for _, word := range words {
		// Balikkan setiap kata
		r := []rune(word)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		reversedWords = append(reversedWords, string(r))
	}

	return strings.Join(reversedWords, " ")
}

// Soal 2
func FizzBuzz(a int) {
	for i := 1; i <= a; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Soal 3
func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	result := make([]int, n)
	result[0] = 0
	result[1] = 1

	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result
}

// Soal 4
func FindMinPrice(arr []int) int {
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

// Soal 5
func CekListStr(arr []string) int {
	count := 0
	for _, val := range arr {
		if len(val) == 1 {
			if val[0] >= '0' && val[0] <= '9' {
				count++
			}
		}
	}
	return count
}
