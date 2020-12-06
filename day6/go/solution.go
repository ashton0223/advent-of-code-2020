package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	groups := strings.Split(string(file), "\n\n")
	groupCounts := countAllGroups(groups, "anyone")
	total := addAll(groupCounts)
	fmt.Println(total)
	groupCounts = countAllGroups(groups, "everyone")
	total = addAll(groupCounts)
	fmt.Println(total)
}

// Req is used to know whether you are trying to find anyone or everyone who answered yes
func countAllGroups(groups []string, req string) []int {
	count := make([]int, 0)

	function := getFunc(req)

	for _, group := range groups {
		count = append(count, function(strings.Split(group, "\n")))
	}
	return count
}

func countGroup(group []string) int {
	combined := strings.Join(group, "")
	return len(getUniqueChars(combined))
}

func countEveryoneGroup(group []string) int {
	count := 0
	people := len(group)
	combined := strings.Join(group, "")
	unique := getUniqueChars(combined)
	for _, char := range unique {
		charStr := string(char)
		rgx := regexp.MustCompile(charStr)
		occur := rgx.FindAllString(combined, -1)
		if len(occur) >= people {
			count += 1
		}
	}
	return count
}

func addAll(nums []int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}

func getUniqueChars(input string) string {
	output := ""
	for _, char := range input {
		charStr := string(char)
		if !strings.Contains(output, charStr) {
			output += charStr
		}
	}
	return output
}

func getFunc(req string) func([]string) int {
	switch req {
	case "anyone":
		return func(groups []string) int {
			return countGroup(groups)
		}
	case "everyone":
		return func(groups []string) int {
			return countEveryoneGroup(groups)
		}
	}
	return nil
}
