package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	count := 0
	validCount := 0
	oldCount := 0
	oldValidCount := 0
	data, err := ioutil.ReadFile("../input")
        if err != nil {
                panic(err)
        }

	strData := string(data)
	strDataArr := strings.Split(strData,"\n\n")

	for _, i := range strDataArr {
		count += contains([]string{i, "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"})
		parts := regexp.MustCompile("[\\s,\\n]+").Split(i, -1)
		validCount += validate(parts)
		if count == oldCount && oldValidCount != validCount {
			validCount -= 1
		}
		oldCount = count
		oldValidCount = validCount
	}
	fmt.Println(count)
	fmt.Println(validCount)
}

func contains(looking []string) (int) {
	flag := 1
	correct := 1
	for _, i := range looking {
		if flag == 1 {
			flag = 0
			continue
		}
		if !strings.Contains(looking[0], i) {
			correct = 0
		}
	}
	return correct
}

func validate(parts []string) (int) {
	valid := 1
	for _, i := range parts {
	if i == "" {
        	valid = 0
        	continue
        }
	switch string(i[0:3]) {
	case "byr":
		year, _ := strconv.Atoi(string(i[4:8]))
		if (year > 2002) || (year < 1920) {
			valid = 0
		}
		break;
	case "iyr":
		year, _ := strconv.Atoi(string(i[4:8]))
		if year > 2020 || year < 2010 {
			valid = 0
		}
		break;
	case "eyr":
		year, _ := strconv.Atoi(string(i[4:8]))
		if year < 2020 || year > 2030 {
			valid = 0
		}
		break;
	case "hgt":
		if len(i) <= 7 {
			// Definitely not valid
			return 0
		}
		unit := string(i[7])
		fmt.Println(unit)
		if unit == "c" {
			height, _ := strconv.Atoi(string(i[4:7]))
			if height < 150 || height > 193 {
				valid = 0
			}
		} else if unit == "n" {
			height, _ := strconv.Atoi(string(i[4:6]))
			if height < 59 || height > 76 {
				valid = 0
			}
		} else {
			fmt.Println("Invalid unit")
			valid = 0
		}
		break;
	case "hcl":
		regex := "#([a-f]|[0-9]){6}"
		r, _ := regexp.MatchString(regex, i)
		if !r {
			valid = 0
		}
		break;
	case "ecl":
		regex := "amb|blu|brn|gry|grn|hzl|oth"
		r, _ := regexp.MatchString(regex, i)
		if !r {
			valid = 0
		}
		break;
	case "pid":
		regex := "\\d{9}"
		r, _ := regexp.MatchString(regex, i)
		if !r {
			valid = 0
		}
		break;
	}
	}
	return valid
}
