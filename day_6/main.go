package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part2())
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	curr := ""
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += part1Helper(curr)
			curr = ""
		} else {
			curr += line
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total + part1Helper(curr)
}

func part1Helper(s string) int {
	chars := map[byte]int{}
	for i := range s {
		chars[s[i]]++
	}
	return len(chars)
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	curr := []string{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += part2Helper(curr)
			curr = []string{}
		} else {
			curr = append(curr, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total + part2Helper(curr)
}

func part2Helper(input []string) int {
	chars := map[byte]int{}
	for _, s := range input {
		for i := range s {
			chars[s[i]]++
		}
	}

	ans := 0
	for _, f := range chars {
		if f == len(input) {
			ans++
		}
	}

	return ans
}
