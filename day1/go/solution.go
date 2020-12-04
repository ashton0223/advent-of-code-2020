package main 

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("../input")
	if err != nil {
		 panic(err)
	}

	array := separate(string(dat))

	numArray := convert(array)

	twoNums := lookForTwoSum(numArray)
	threeNums := lookForThreeSum(numArray)

	fmt.Println(fmt.Sprintf("%d + %d = %d", twoNums[0], twoNums[1], twoNums[0] + twoNums[1]))
	fmt.Println(fmt.Sprintf("%d * %d = %d", twoNums[0], twoNums[1], twoNums[0] * twoNums[1]))

	fmt.Println(fmt.Sprintf("%d + %d + %d = %d", threeNums[0], threeNums[1], threeNums[2], threeNums[0] + threeNums[1] + threeNums[2]))
	fmt.Println(fmt.Sprintf("%d * %d * %d = %d", threeNums[0], threeNums[1], threeNums[2], threeNums[0] * threeNums[1] * threeNums[2]))
}

func separate(data string) ([]string) {
	return strings.Split(data, "\n")
}

func convert(arr []string) ([]int) {
	numArr := make([]int, len(arr) - 1)
	for i := 0; i < len(arr); i++ {
		// Checks for empty characters
		if arr[i] == "" {
			continue
		}

		j, err := strconv.Atoi(arr[i])
		if err != nil {
			panic(err)
		}

		numArr[i] = j
	}
	return numArr
}

func lookForTwoSum(nums []int) ([]int) {
	sumNums := make([]int, 2)

ExitTwos:
	for i := 0; i < len(nums); i++{
		for j := i; j < len(nums); j++ {
			if nums[i] + nums[j] == 2020 {
				sumNums[0] = nums[i]
				sumNums[1] = nums[j]
				continue ExitTwos
			}
		}
	}
	return sumNums
}

func lookForThreeSum(nums []int) ([]int) {
	sumNums := make([]int, 3)

ExitThrees:
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				if (nums[i] + nums[j] + nums[k]) == 2020 {
					sumNums[0] = nums[i]
					sumNums[1] = nums[j]
					sumNums[2] = nums[k]
					continue ExitThrees
				}
			}
		}
	}
	return sumNums
}
