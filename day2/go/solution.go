package main

import (
        "fmt"
	"io/ioutil"
        "strings"
	"strconv"
)

func main() {
	amount := 0
	amount2 := 0
        dat, err := ioutil.ReadFile("../input")
        if err != nil {
                 panic(err)
        }

	arr := strings.Split(string(dat),"\n")

	for i := 0; i < len(arr) -2; i++ {
		arr2 := strings.Split(arr[i], " ")
		arr3 := strings.Split(arr2[0], "-")
		min, _ := strconv.Atoi(arr3[0])
		max, _ := strconv.Atoi(arr3[1])
		letter := arr2[1][0]
		count := strings.Count(arr2[2], string(letter))
		if !(count > max || count < min) {
			amount++
		}
		if max > len(arr2[2]) {
			max = 1
			min = 1
		}
		maxindex := arr2[2][max-1]
		minindex := arr2[2][min-1]
		if (string(maxindex) == string(letter)) && !(string(minindex) == string(letter)) {
			amount2++
			fmt.Println("max only",string(maxindex),max,min)
			fmt.Println(string(arr2[2]))
		}
		if !(string(maxindex) == string(letter)) && (string(minindex) == string(letter)) {
			fmt.Println("min only",string(minindex),max,min)
			fmt.Println(string(arr2[2]))
			amount2++
		}
	}
	fmt.Println(amount)
	fmt.Println(amount2)
}
