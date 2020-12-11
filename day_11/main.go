package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	originalSeats := getSeats()

	// part 1
	taken := -1
	seats, nextTaken := getNextRound(4, originalSeats, getAdjacentTaken)
	for taken != nextTaken {
		taken = nextTaken
		seats, nextTaken = getNextRound(4, seats, getAdjacentTaken)
	}
	fmt.Println("taken seats after folks stopped moving", taken)

	// part 2
	taken = -1
	seats, nextTaken = getNextRound(5, originalSeats, getAdjacentTakenLOS)
	for taken != nextTaken {
		taken = nextTaken
		seats, nextTaken = getNextRound(5, seats, getAdjacentTakenLOS)
	}
	fmt.Println("part2", taken)
}

func getSeats() [][]byte {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seats := [][]byte{}
	for scanner.Scan() {
		seats = append(seats, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return seats
}

func getAdjacentTaken(r, c int, seats [][]byte) int {
	return isSeatTaken(r-1, c, seats) +
		isSeatTaken(r-1, c-1, seats) +
		isSeatTaken(r-1, c+1, seats) +
		isSeatTaken(r, c-1, seats) +
		isSeatTaken(r, c+1, seats) +
		isSeatTaken(r+1, c-1, seats) +
		isSeatTaken(r+1, c, seats) +
		isSeatTaken(r+1, c+1, seats)
}
func isSeatTaken(r, c int, seats [][]byte) int {
	if r >= 0 && c >= 0 && r < len(seats) && c < len(seats[r]) && seats[r][c] == '#' {
		return 1
	}
	return 0
}

func getAdjacentTakenLOS(r, c int, seats [][]byte) int {
	ans := 0
	for i := 0; i < 8; i++ {
		ans += getTakenSeatsLOS(i, r, c, seats)
	}
	return ans
}

func getTakenSeatsLOS(dir, r, c int, seats [][]byte) int {
	switch dir {
	case 0:
		r--
		for r >= 0 && seats[r][c] == '.' {
			r--
		}
		if r < 0 || seats[r][c] != '#' {
			return 0
		}
		return 1

	case 1:
		r++
		for r < len(seats) && seats[r][c] == '.' {
			r++
		}
		if r == len(seats) || seats[r][c] != '#' {
			return 0
		}
		return 1

	case 2:
		c--
		for c >= 0 && seats[r][c] == '.' {
			c--
		}
		if c < 0 || seats[r][c] != '#' {
			return 0
		}
		return 1

	case 3:
		c++
		for c < len(seats[0]) && seats[r][c] == '.' {
			c++
		}
		if c == len(seats[0]) || seats[r][c] != '#' {
			return 0
		}
		return 1

	case 4:
		r--
		c--
		for c >= 0 && r >= 0 && seats[r][c] == '.' {
			r--
			c--
		}
		if r < 0 || c < 0 || seats[r][c] != '#' {
			return 0
		}
		return 1

	case 5:
		r++
		c--
		for c >= 0 && r < len(seats) && seats[r][c] == '.' {
			r++
			c--
		}
		if c < 0 || r == len(seats) || seats[r][c] != '#' {
			return 0
		}
		return 1

	case 6:
		c++
		r++
		for r < len(seats) && c < len(seats[0]) && seats[r][c] == '.' {
			c++
			r++
		}
		if r == len(seats) || c == len(seats[0]) || seats[r][c] != '#' {
			return 0
		}
		return 1

	default:
		r--
		c++
		for r >= 0 && c < len(seats[0]) && seats[r][c] == '.' {
			c++
			r--
		}
		if r < 0 || c == len(seats[0]) || seats[r][c] != '#' {
			return 0
		}
		return 1
	}
}

func getNextRound(emptySeatsThreshold int, seats [][]byte, getAdjacentTaken func(r, c int, seats [][]byte) int) ([][]byte, int) {
	numTaken := 0
	next := make([][]byte, len(seats))
	for r := range seats {
		next[r] = make([]byte, len(seats[r]))
		for c := range seats[r] {
			at := getAdjacentTaken(r, c, seats)
			if seats[r][c] == 'L' && at == 0 {
				next[r][c] = '#'
			} else if seats[r][c] == '#' && at >= emptySeatsThreshold {
				next[r][c] = 'L'
			} else {
				next[r][c] = seats[r][c]
			}

			if next[r][c] == '#' {
				numTaken++
			}
		}
	}
	return next, numTaken
}
