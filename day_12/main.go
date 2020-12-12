package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	r, c := getDest()
	fmt.Println(r, c)

	r, c = getDestWithWaypoint()
	fmt.Println(r, c)
}

func getDest() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	r, c, dir := 0, 0, 0
	for scanner.Scan() {
		move(&r, &c, &dir, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return r, c
}

func getDestWithWaypoint() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	r, c, wr, wc := 0, 0, 10, 1
	for scanner.Scan() {
		moveWP(&r, &c, &wr, &wc, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return r, c
}

// dir values
// 0 is east
// 1 is north
// 2 is west
// 3 is south
func move(r, c, dir *int, s string) {
	val, _ := strconv.Atoi(s[1:])
	switch s[0] {
	case 'L':
		*dir = (*dir + val/90) % 4
	case 'R':
		*dir = (*dir - val/90) % 4
		if *dir < 0 {
			*dir += 4
		}
	case 'F':
		switch *dir {
		case 0:
			*r += val
		case 1:
			*c += val
		case 2:
			*r -= val
		case 3:
			*c -= val
		}
	case 'E':
		*r += val
	case 'N':
		*c += val
	case 'W':
		*r -= val
	case 'S':
		*c -= val
	}
}

func moveWP(r, c, wr, wc *int, s string) {
	val, _ := strconv.Atoi(s[1:])
	switch s[0] {
	case 'R':
		switch val {
		case 90:
			*wc, *wr = -1*(*wr), *wc
		case 180:
			*wc, *wr = -1*(*wc), -1*(*wr)
		case 270:
			*wc, *wr = (*wr), -1*(*wc)
		}
	case 'L':
		switch val {
		case 90:
			*wc, *wr = (*wr), -1*(*wc)
		case 180:
			*wc, *wr = -1*(*wc), -1*(*wr)
		case 270:
			*wc, *wr = -1*(*wr), *wc
		}
	case 'F':
		*r += val * (*wr)
		*c += val * (*wc)
	case 'E':
		*wr += val
	case 'N':
		*wc += val
	case 'W':
		*wr -= val
	case 'S':
		*wc -= val
	}
}
