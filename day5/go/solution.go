package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}
	strFile := string(file)
	strFileLines := strings.Split(strFile, "\n")

	partArr := make([]int, 0)
	for _, line := range strFileLines {
		total := 128
		totalCol := 8
		addnum := 0
		colnum := 0
		for indexL, i := range line {
			if string(i) == "B" {
				addnum += (total / int(math.Pow(float64(2), float64(indexL+1))))
			}
			if string(i) == "R" {
				colnum += (totalCol / int(math.Pow(float64(2), float64(indexL-6))))
			}
		}
		partArr = append(partArr, (addnum*8)+colnum)
	}
	fmt.Println(findLargest(partArr))
	fmt.Println(findMissing(partArr))
}

func findLargest(nums []int) int {
	largest := 0
	for _, i := range nums {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func findMissing(nums []int) int {
	sort.Ints(nums)
	for i, id := range nums {
		next := id + 1
		if next != nums[i+1] {
			return id + 1
		}
	}
	return 0
}
