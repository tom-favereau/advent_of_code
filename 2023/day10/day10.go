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
}

func Aire(points []utils.Pos) int {
	n := len(points)
	if n < 3 {
		return 0.0
	}

	aire := 0

	for i := 0; i < n; i++ {
		x1, y1 := points[i].X, points[i].Y
		x2, y2 := points[(i+1)%n].X, points[(i+1)%n].Y
		aire += (x1*y2 - x2*y1)
	}

	aire = int(0.5 * float64(aire))
	if aire < 0 {
		aire = -aire
	}

	return aire
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
	for scanner.Scan() {
		line := scanner.Text()

		p.X = 0
		arr := strings.Split(line, "")
		for _, u := range arr {
			grid.Add(u, p)
			p.X++
		}
		p.Y++
	}

	loop := make([]utils.Pos, 0, 10)
	//loop = append(loop, utils.Pos{1, 1})
	//loop = Loop(grid, utils.Pos{2, 1}, "right", loop)
	loop = append(loop, utils.Pos{18, 74})
	loop = Loop(grid, utils.Pos{19, 74}, "right", loop)

	return len(loop) / 2
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
	for scanner.Scan() {
		line := scanner.Text()

		p.X = 0
		arr := strings.Split(line, "")
		for _, u := range arr {
			grid.Add(u, p)
			p.X++
		}
		p.Y++
	}

	loop := make([]utils.Pos, 0, 10)
	//loop = append(loop, utils.Pos{1, 1})
	//loop = Loop(grid, utils.Pos{2, 1}, "right", loop)
	loop = append(loop, utils.Pos{18, 74})
	loop = Loop(grid, utils.Pos{19, 74}, "right", loop)

	res := 0
	for i := 0; i < p.Y; i++ {
		for j := 0; j < p.X; j++ {
			if !contain(loop, utils.Pos{j, i}) {
				inx := false
				for k := j; k < p.X; k++ {
					if contain(loop, utils.Pos{k, i}) && grid[utils.Pos{k, i}] == "|" {
						inx = !inx
					}
				}
				iny := false
				for k := i; k < p.Y; k++ {
					if contain(loop, utils.Pos{j, k}) && grid[utils.Pos{j, k}] == "-" {
						iny = !iny
					}
				}
				if !inx || !iny {

				} else {
					res++
				}
			}
		}
	}

	aire := Aire(loop)
	res = aire - len(loop)/2 + 1

	return res
}

func Loop(g utils.Grid[string], p utils.Pos, dir string, ans []utils.Pos) []utils.Pos {
	switch g[p] {
	case "S":
		return ans
	case "|":
		if dir == "down" {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X, p.Y + 1}, "down", ans)
		} else {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X, p.Y - 1}, "up", ans)
		}
	case "-":
		if dir == "right" {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X + 1, p.Y}, "right", ans)
		} else {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X - 1, p.Y}, "left", ans)
		}
	case "L":
		if dir == "left" {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X, p.Y - 1}, "up", ans)
		} else {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X + 1, p.Y}, "right", ans)
		}
	case "J":
		if dir == "right" {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X, p.Y - 1}, "up", ans)
		} else {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X - 1, p.Y}, "left", ans)
		}
	case "F":
		if dir == "left" {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X, p.Y + 1}, "down", ans)
		} else {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X + 1, p.Y}, "right", ans)
		}
	case "7":
		if dir == "right" {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X, p.Y + 1}, "down", ans)
		} else {
			ans = append(ans, p)
			return Loop(g, utils.Pos{p.X - 1, p.Y}, "left", ans)
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func contain(p []utils.Pos, pos utils.Pos) bool {
	for _, u := range p {
		if u.X == pos.X && u.Y == pos.Y {
			return true
		}
	}
	return false
}
