package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read in file called "input"
	foo := []int{1, 2}
	fmt.Printf("length of foo: %d\n", len(foo))

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("booooo")
	}
	defer file.Close()

	dist_x, dist_y := []int{}, []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dists := strings.Fields(line)
		x, _ := strconv.Atoi(dists[0])
		y, _ := strconv.Atoi(dists[1])

		if len(dist_x) == 0 { dist_x = append(dist_x, x) }
		if len(dist_y) == 0 { dist_y = append(dist_y, y) }

		for idx, val := range dist_x {
			if x <= val {
				foo := []
			}
		}
	}
}
