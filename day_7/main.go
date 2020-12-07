package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// part 1
	graph := createGraph()
	ans := 0
	for _, bag := range graph {
		if bag.color == "shiny gold" {
			continue
		}
		if bag.carries("shiny gold") {
			ans++
		}
	}
	fmt.Println("bags that can carry shiny gold bag:", ans)

	// part 2, subtract one for shiny gold bag itself
	fmt.Println("num bags required for shiny gold: ", graph["shiny gold"].count()-1)
}

func createGraph() map[string]*node {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	graph := map[string]*node{}
	for scanner.Scan() {
		line := scanner.Text()
		addRule(line, graph)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return graph
}

type node struct {
	color       string
	contains    []*node
	numContains []int
}

func (n *node) carries(color string) bool {
	if n.color == color {
		return true
	}
	ans := false
	for _, c := range n.contains {
		ans = ans || c.carries(color)
	}
	return ans
}

func (n *node) count() int {
	ans := 1
	for i := range n.contains {
		ans += n.numContains[i] * n.contains[i].count()
	}
	return ans
}

func addRule(rule string, graph map[string]*node) {
	parts := strings.Split(rule, " bags contain ")
	parentColor := parts[0]
	if graph[parentColor] == nil {
		graph[parentColor] = &node{
			color:       parentColor,
			contains:    []*node{},
			numContains: []int{},
		}
	}
	if parts[1] == "no other bags." {
		return
	}

	contents := strings.Split(parts[1][:len(parts[1])-1], ", ")
	for _, c := range contents {
		num, childColor := parseContent(c)
		if graph[childColor] == nil {
			graph[childColor] = &node{
				color:       childColor,
				contains:    []*node{},
				numContains: []int{},
			}
		}
		graph[parentColor].contains = append(graph[parentColor].contains, graph[childColor])
		graph[parentColor].numContains = append(graph[parentColor].numContains, num)
	}
}

func parseContent(s string) (int, string) {
	if s[len(s)-1] == 's' {
		s = s[:len(s)-5]
	} else {
		s = s[:len(s)-4]
	}
	i := 0
	num := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		num = 10*num + int(s[i]) - 48
		i++
	}
	return num, s[i+1:]
}
