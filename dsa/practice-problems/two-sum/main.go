package main

import (
	"fmt"
)

func main() {
	numArr := []int{3, 2, 4}
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	}
	result := calcHash(numArr, i)
	fmt.Println(result)
}

func calc(nums []int, target int) []int {
	for i, value := range nums {
		temp := target - value
		for j, value := range nums {
			if value == temp {
				return []int{i, j}
			}
		}
	}
	return nil
}

func calcHash(nums []int, target int) []int {
	//map[KeyType]ValueType
	hashmap := make(map[int]int)
	for i, value := range nums {
		hashmap[value] = i
	}
	for i, value := range nums {
		temp := target - value
		j, ok := hashmap[temp]
		if ok == true && j != i {
			return []int{i, j}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {

	m := map[int]int{}
	for i, x := range nums {
		if j, ok := m[target-x]; ok {
			return []int{i, j}
		}
		m[x] = i
	}
	return []int{}
}
