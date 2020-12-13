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
	start, buses := getInput()
	fmt.Println(start, buses)

	// part 1
	earliest, busID := getEarliestBus(start, buses)
	fmt.Println(earliest, busID, (earliest-start)*busID)

	// part2
	fmt.Println("winGoldCoin", winGoldCoin(buses))
}

func getInput() (int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	busStrings := strings.Split(scanner.Text(), ",")
	buses := [][]int{}
	for i, s := range busStrings {
		if s == "x" {
			continue
		}
		b, _ := strconv.Atoi(s)
		buses = append(buses, []int{b, i})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return start, buses
}

func getEarliestBus(time int, buses [][]int) (int, int) {
	for true {
		for _, b := range buses {
			if time%b[0] == 0 {
				return time, b[0]
			}
		}
		time++
	}
	return -1, -1
}

func winGoldCoin(buses [][]int) int64 {
	step := int64(buses[0][0])
	var time int64
	for i := 1; i < len(buses); i++ {
		for (time+int64(buses[i][1]))%int64(buses[i][0]) != 0 {
			time += step
		}
		step *= int64(buses[i][0])
	}
	return time
}
