package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	instructions := getInstructions()

	// part 1
	fmt.Println("accumulator value at time of loop:", getValueAtLoop(instructions))

	// part 2
	fmt.Println("accumulator after fix", accumulatorValAfterFix(instructions))
}

func getInstructions() []*instruction {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	out := []*instruction{}
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line[4:])
		out = append(out, &instruction{
			action: line[:3],
			val:    v,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out

}

type instruction struct {
	action string
	val    int
}

func getValueAtLoop(instructions []*instruction) int {
	visited := map[int]bool{}
	ans := 0
	i := 0
	for i < len(instructions) && !visited[i] {
		visited[i] = true
		switch instructions[i].action {
		case "nop":
			i++
		case "acc":
			ans += instructions[i].val
			i++
		case "jmp":
			i += instructions[i].val
		}
	}
	return ans
}

func accAtEnd(i, acc int, instructions []*instruction) int {
	visited := map[int]bool{}
	for i < len(instructions) && !visited[i] {
		visited[i] = true
		switch instructions[i].action {
		case "nop":
			i++
		case "acc":
			acc += instructions[i].val
			i++
		case "jmp":
			i += instructions[i].val
		}
	}
	if i == len(instructions) {
		return acc
	}
	return math.MinInt64
}

func accumulatorValAfterFix(instructions []*instruction) int {
	acc := 0
	i := 0
	for i < len(instructions) {
		switch instructions[i].action {
		case "nop":
			instructions[i].action = "jmp"
			val := accAtEnd(i, acc, instructions)
			if val != math.MinInt64 {
				return val
			}
			instructions[i].action = "nop"
			i++
		case "jmp":
			instructions[i].action = "nop"
			val := accAtEnd(i, acc, instructions)
			if val != math.MinInt64 {
				return val
			}
			instructions[i].action = "jmp"
			i += instructions[i].val
		case "acc":
			acc += instructions[i].val
			i++
		}
	}
	return math.MinInt64
}
