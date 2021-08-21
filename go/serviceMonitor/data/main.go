package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value     int
	frequency int
	leftPtr   *node
	rightPtr  *node
}

func generate() {
	f, err := os.Create("./data/dataShort.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	i := 0
	for i < 100 {
		_, err := f.WriteString(fmt.Sprintf("%d\n", rand.Intn(10000)))
		if err != nil {
			fmt.Println(err)
		}
		i++
	}
}
func incrementNode(n *node) {
	n.frequency++
}
func createNode(v int) *node {
	newNode := node{
		value:     v,
		frequency: 1,
		leftPtr:   nil,
		rightPtr:  nil,
	}
	return &newNode
}
func sort(root *node, v int) {
	if v < root.value {
		//go left
		if root.leftPtr == nil {
			//create node here, point root.leftPtr toward it
			newNode := createNode(v)
			root.leftPtr = newNode
		} else {
			//recurse
			sort(root.leftPtr, v)
		}

	} else if v > root.value {
		//go right
		if root.rightPtr == nil {
			// create node, point root.rightPtr toward it
			newNode := createNode(v)
			root.rightPtr = newNode
		} else {
			//recurse
			sort(root.rightPtr, v)
		}

	} else if v == root.value {
		fmt.Printf("Frequency increased for: %v\n", v)
		incrementNode(root)
	}
}

func tree(nums *[]int) {
	root := &node{
		value:     (*nums)[0],
		frequency: 1,
		leftPtr:   nil,
		rightPtr:  nil,
	}
	for i := 1; i < len(*nums); i++ {
		sort(root, (*nums)[i])
	}

}
func main() {
	//generate()
	data, err := ioutil.ReadFile("./data/dataShort.txt")
	if err != nil {
		log.Fatal(err)
	}
	numsStr := strings.Split(string(data), "\n")
	var nums []int
	for i := 0; i < len(numsStr); i++ {
		if numsStr[i] != "" {
			temp, err := strconv.Atoi(numsStr[i])
			if err != nil {
				log.Println(err)
				continue
			} else {
				//fmt.Println(temp)
				nums = append(nums, temp)
			}
		}
	}
	tree(&nums)
}
