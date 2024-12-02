package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read in file called "input"fmt.Printf("and our final diff is: %d\n", diff)

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

		dist_x = append(dist_x, x)
		dist_y = append(dist_y, y)
	}

	sort.Ints(dist_x)
	sort.Ints(dist_y)

	diff := 0
	for i := 0; i < 1000; i++ {
		tmp := dist_x[i] - dist_y[i]
		if tmp < 0 {
			tmp = -tmp
		}
		diff += tmp
	}
	fmt.Printf("and our final diff is: %d\n", diff)
}
