package main

import "fmt"

type pair struct {
	value int
}

func main() {
	numArr := [8]int{1, 2, 3, 1, 2, 3, 1, 4}
	matchValues(&numArr)
}
func matchValues(numArr *[8]int) {
	pairs := make([]pair, 0)
	for i := 0; i < len(numArr); i++ {
		for j := i + 1; j < len(numArr); j++ {
			if numArr[i] == numArr[j] {
				pairs = append(pairs, pair{numArr[i]})
				temp := numArr[j]
				numArr[j] = numArr[i+1]
				numArr[i+1] = temp
				i = i + 1
				break
			}
		}
	}
	fmt.Println(pairs)
}
