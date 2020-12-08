package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type cpu struct {
	pc  int
	acc int
}

func (c *cpu) execLine(line string) {
	// op[0] is the instruction, op[1] is what is given to the instruction
	op := strings.Split(line, " ")

	// Number passed to the instruction
	num, _ := strconv.Atoi(op[1])

	switch op[0] {
	case "acc":
		c.acc += num
		c.pc++
		break
	case "jmp":
		c.pc += num
		break
	case "nop":
		c.pc++
		break
	default:
		panic("Invalid instruction")
	}
}

func main() {
	file, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	program := strings.Split(string(file), "\n")

	noFix := getRunProgram(program, false)
	acc, _ := noFix()
	fmt.Print("Accumulator value is: ")
	fmt.Println(acc)

	fix := getRunProgram(program, true)
	acc, _ = fix()
	fmt.Print("Accumulator value without loop is: ")
	fmt.Println(acc)
}

func getRunProgram(program []string, withChecks bool) func() (int, int) {
	// Initialize cpu with pc and acc set to 0, w/ no op
	c := cpu{0, 0}

	if withChecks {
		return func() (int, int) {
			return execWithChecks(c, program)
		}
	} else {
		return func() (int, int) {
			return execWithoutChecks(c, program)
		}
	}
}

func execWithoutChecks(c cpu, program []string) (int, int) {
	// Keep track of what lines have been run
	ran := make(map[int]bool)

	for c.pc < len(program) {
		// Check for lines already run, and break if so
		if ran[c.pc] {
			break
		}

		ran[c.pc] = true
		c.execLine(program[c.pc])
	}

	return c.acc, c.pc
}

func execWithChecks(c cpu, program []string) (int, int) {
	// Used to swap values
	swap := map[string]string{"nop": "jmp", "jmp": "nop"}

	found := false
	for c.pc < len(program) {
		program2 := make([]string, len(program))
		copy(program2, program)
		// Spawn new CPU without checks if instruction is nop or jmp
		if instruction, ok := swap[program2[c.pc][0:3]]; ok && !found {
			program2[c.pc] = strings.Replace(program2[c.pc], swap[instruction], instruction, 1)
			sc := cpu{c.pc, 0}
			_, ran := execWithoutChecks(sc, program2)
			if ran == len(program) {
				found = true
				program[c.pc] = strings.Replace(program[c.pc], swap[instruction], instruction, 1)
			}
		}
		c.execLine(program[c.pc])
	}
	return c.acc, c.pc
}
