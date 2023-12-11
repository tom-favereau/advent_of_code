package main

import (
	utils "advent_of_code/utils/grid"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("input.txt"))
	//fmt.Println(part2("input2.txt"))
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := utils.NewGrid[string]()
	p := utils.Pos{0, 0}

	gal := make([]utils.Pos, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, "")

		p.X = 0
		for _, u := range arr {
			if u == "#" {
				gal = append(gal, p)
			}
			grid.Add(u, p)
			p.X++
		}

		p.Y++

	}

	res := 0
	for i, u1 := range gal {
		for j, u2 := range gal {
			fmt.Println(i, j)
			if j > i {
				tmp := 0
				for k := min(u1.X, u2.X); k < max(u1.X, u2.X); k++ {
					if cEmpty(grid, k, p.X) {
						tmp += 1000000
					} else {
						tmp++
					}
				}
				for k := min(u1.Y, u2.Y); k < max(u1.Y, u2.Y); k++ {
					if lEmpty(grid, k, p.X) {
						tmp += 1000000
					} else {
						tmp++
					}
				}
				res += tmp
			}
		}
	}

	return res
}

func cEmpty(g utils.Grid[string], p int, max int) bool {
	pos := utils.Pos{p, 0}
	for i := 0; i < max; i++ {
		if g[pos] != "." {
			return false
		}
		pos.Y++
	}
	return true
}

func lEmpty(g utils.Grid[string], p int, max int) bool {
	pos := utils.Pos{0, p}
	for i := 0; i < max; i++ {
		if g[pos] != "." {
			return false
		}
		pos.X++
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
