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
	passports := getPassports()
	countValid := 0
	for _, p := range passports {
		if isValidPassport(p) {
			countValid++
		}
	}
	fmt.Println(countValid)
}

func getPassports() []map[string]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	ans := []map[string]string{}
	curr := ""
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			ans = append(ans, newPassport(curr[1:]))
			curr = ""
		} else {
			curr += " " + line
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return append(ans, newPassport(curr[1:]))
}

func newPassport(data string) map[string]string {
	passport := map[string]string{}
	parts := strings.Split(data, " ")
	for _, part := range parts {
		p := strings.Split(part, ":")
		passport[p[0]] = p[1]
	}
	return passport
}

func isValidPassport(passport map[string]string) bool {
	if !isValidYear(passport["byr"], 1920, 2002) {
		// fmt.Println("invalid birth year", passport)
		return false
	}
	if !isValidYear(passport["iyr"], 2010, 2020) {
		// fmt.Println("isValidYear iyr", passport)
		return false
	}
	if !isValidYear(passport["eyr"], 2020, 2030) {
		// fmt.Println("isValidYear eyr", passport)
		return false
	}
	if !isValidHeight(passport["hgt"]) {
		// fmt.Println("isValidHeight", passport)
		return false
	}
	if !isValidHair(passport["hcl"]) {
		// fmt.Println("isValidHair", passport)
		return false
	}
	if !isValidEye(passport["ecl"]) {
		// fmt.Println("isValidEye", passport)
		return false
	}
	if !isValidPID(passport["pid"]) {
		// fmt.Println("isValidPID", passport)
		return false
	}
	return true
}

func isValidYear(s string, min, max int) bool {
	year, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return min <= year && year <= max
}

func isValidHeight(s string) bool {
	if len(s) == 4 {
		if s[2:] != "in" {
			return false
		}
		inches, err := strconv.Atoi(s[:2])
		if err != nil {
			return false
		}
		return 59 <= inches && inches <= 76
	}

	if len(s) == 5 {
		if s[3:] != "cm" {
			return false
		}
		cms, err := strconv.Atoi(s[:3])
		if err != nil {
			return false
		}
		return 150 <= cms && cms <= 193
	}

	return false
}

func isValidHair(s string) bool {
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	s = s[1:]
	for i := range s {
		if ('0' <= s[i] && s[i] <= '9') || ('a' <= s[i] && s[i] <= 'f') {
			continue
		}
		return false
	}
	return true
}

func isValidEye(s string) bool {
	valid := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return valid[s]
}

func isValidPID(s string) bool {
	if len(s) != 9 {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}
