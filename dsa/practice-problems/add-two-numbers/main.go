package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
	}

	result := reverse2(input)
	fmt.Println(result)
}

func reverse2(x int) int {
	outcome := x % 10
	if x != 0 {
		fmt.Printf("Result: %d\n", outcome)
		return reverse2(x / 10)
	}
	return x
}
func reverse(x int) int {
	sign := false
	if x < 0 {
		sign = true
	}

	input := strconv.Itoa(x)
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[:]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)

	if sign == true {
		output = output[:len(output)-1]
		result, err := strconv.Atoi(output)
		if err != nil {
			fmt.Println(err)
		}
		result = 0 - result
		return result
	} else {
		result, err := strconv.Atoi(output)
		if err != nil {
			fmt.Println(err)
		}
		return result
	}
}

func willOverflow(number, newDigit int) bool {
	caseOne := number > math.MaxInt32/10
	caseTwo := number < math.MinInt32/10
	caseThree := number == math.MaxInt32 && newDigit > 7
	caseFour := number == math.MinInt32 && newDigit < -8

	if caseOne || caseTwo || caseThree || caseFour {
		return true
	}

	return false
}
