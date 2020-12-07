package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}
	rules := strings.Split(string(file), "\n")
	fmt.Println(len(findBig(rules, "shiny gold")))
	fmt.Println(findSmall(rules, "shiny gold", 1))
}

func findBig(rules []string, bagtype string) []string {
	bags := make([]string, 0)
	for _, i := range rules {
		find := strings.Split(i, " contain")
		if strings.Contains(find[1], bagtype) {
			findMore := strings.Split(find[0], " bag")
			bags = append(bags, findMore[0])
			morebags := findBig(rules, findMore[0])
			for _, j := range morebags {
				bags = append(bags, j)
			}
		}
	}
	return removeDuplicates(bags)
}

func findSmall(rules []string, bagtype string, amount int) int {
	bagsAmount := 0
	for _, i := range rules {
		find := strings.SplitN(i, " bags contain", 2)
		if strings.Contains(find[0], bagtype) {
			rgx := regexp.MustCompile("(\\d)\\s(\\w+\\s\\w+)(?:(\\sbags?))")
			if len(find) < 2 {
				continue
			}
			res := rgx.FindAllStringSubmatch(find[1], -1)
			for _, j := range res {
				inAmount, _ := strconv.Atoi(j[1])
				bagsAmount += inAmount
				//inAmount *= amount
				inBagtype := j[2]
				bagsAmount += findSmall(rules, inBagtype, inAmount)
			}
		}
	}
	return bagsAmount * amount
}

func removeDuplicates(inputs []string) []string {
	strMap := make(map[string]bool)
	outputs := make([]string, 0)
	for _, input := range inputs {
		if !strMap[input] {
			strMap[input] = true
			outputs = append(outputs, input)
		}
	}
	return outputs
}
