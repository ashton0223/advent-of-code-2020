package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	strData := string(data)
	strDataArr := strings.Split(strData, "\n")
	lineLength := len(strDataArr[0])
	hCount := 0
	vCount := 0 // Only used when the vertical distance is more than 1
	vCNeeded := false // Used to know whether to check for vertical count
	treeCount := 0
	hArr := []int{1,3,5,7,1}
	vArr := []int{1,1,1,1,2}
	results := make([]int, len(hArr))

	for i, distance := range hArr {
		if vArr[i] > 1 {
			vCNeeded = true
		} else {
			vCNeeded = false
		}
		for _,line := range strDataArr {
			// Check for going down more than once at a time
			// Need to skip the first line in order to go down
			if hCount == 0 {
				hCount += distance
				if vCNeeded {
					vCount += 1
				}
				continue
			}

			// Checks if you have moved down enough
			if vCNeeded {
				if vCount % vArr[i] != 0 {
					vCount += 1
					continue
				}
			}

			// Checks for empty lines
			if len(line) == 0 {
				continue
			}
			place := line[hCount % lineLength]
			if string(place) == "#" {
				treeCount++
			}

			hCount += distance
			vCount += 1
		}
		hCount = 0
		vCount = 0
		fmt.Println(treeCount)
		results[i] = treeCount
		treeCount = 0
	}
	product := 1
	for _, num := range results {
		product *= num
	}
	fmt.Println(product)
}
