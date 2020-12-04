package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	ans := 0
	for _, line := range input {
		if isValid(parseLine(line)) {
			ans++
		}
	}
	fmt.Println(ans)
}

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func parseLine(input string) (int, int, byte, string) {
	p1 := strings.Split(input, ": ")
	password := p1[1]
	p2 := strings.Split(p1[0], " ")
	char := p2[1][0]
	nums := strings.Split(p2[0], "-")
	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])
	return min, max, char, password
}

func isValid(min, max int, char byte, password string) bool {
	count := 0
	if password[min-1] == char {
		count++
	}
	if password[max-1] == char {
		count++
	}
	return count == 1
}
