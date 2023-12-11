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
	fmt.Println(part2("input.txt"))
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

	isEmptyX := make(map[int]bool)
	for i := 0; i < p.X; i++ {
		empty := true
		for j := 0; j < p.Y; j++ {
			if grid[utils.Pos{i, j}] == "#" {
				empty = false
				break
			}
		}
		if empty {
			isEmptyX[i] = true
		}
	}

	isEmptyY := make(map[int]bool)
	for j := 0; j < p.Y; j++ {
		empty := true
		for i := 0; i < p.X; i++ {
			if grid[utils.Pos{i, j}] == "#" {
				empty = false
				break
			}
		}
		if empty {
			isEmptyY[j] = true
		}
	}

	res := 0
	for i, u1 := range gal {
		for j, u2 := range gal {
			if j > i {
				tmp := 0
				for k := min(u1.X, u2.X); k < max(u1.X, u2.X); k++ {
					if isEmptyY[k] {
						tmp += 2
					} else {
						tmp++
					}
				}
				for k := min(u1.Y, u2.Y); k < max(u1.Y, u2.Y); k++ {
					if isEmptyX[k] {
						tmp += 2
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

func part2(string2 string) int {
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

	isEmptyX := make(map[int]bool)
	for i := 0; i < p.X; i++ {
		empty := true
		for j := 0; j < p.Y; j++ {
			if grid[utils.Pos{i, j}] == "#" {
				empty = false
				break
			}
		}
		if empty {
			isEmptyX[i] = true
		}
	}

	isEmptyY := make(map[int]bool)
	for j := 0; j < p.Y; j++ {
		empty := true
		for i := 0; i < p.X; i++ {
			if grid[utils.Pos{i, j}] == "#" {
				empty = false
				break
			}
		}
		if empty {
			isEmptyY[j] = true
		}
	}

	res := 0
	for i, u1 := range gal {
		for j, u2 := range gal {
			if j > i {
				tmp := 0
				for k := min(u1.X, u2.X); k < max(u1.X, u2.X); k++ {
					if isEmptyY[k] {
						tmp += 1000000
					} else {
						tmp++
					}
				}
				for k := min(u1.Y, u2.Y); k < max(u1.Y, u2.Y); k++ {
					if isEmptyX[k] {
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
