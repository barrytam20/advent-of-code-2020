package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	_, taken := getIDs()
	fmt.Println(findSeat(taken))
}

func getIDs() (int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	ids := []int{}
	max := 0
	for scanner.Scan() {
		id := calcID(scanner.Text())
		if id > max {
			max = id
		}
		ids = append(ids, id)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return max, ids
}

func calcID(s string) int {
	front, back := 0, 127
	for i := 0; i < 7; i++ {
		mid := (front + back) / 2
		if s[i] == 'F' {
			back = mid
		} else {
			front = mid + 1
		}
	}

	left, right := 0, 7
	for i := 7; i < len(s); i++ {
		mid := (left + right) / 2
		if s[i] == 'R' {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return front*8 + left
}

func findSeat(taken []int) int {
	seats := map[int]bool{}
	for row := 1; row < 127; row++ {
		for col := 0; col < 8; col++ {
			seats[row*8+col] = true
		}
	}
	for _, id := range taken {
		delete(seats, id)
	}
	for k := range seats {
		if !seats[k-1] && !seats[k+1] {
			return k
		}
	}
	return -1
}
