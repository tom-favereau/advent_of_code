package main

import (
	utils "advent_of_code/utils/grid"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Point représente un point dans l'espace en 2D.

// CalculerAirePolygone calcule l'aire d'un polygone donné par une liste de points.
func Aire(points []utils.Pos) int {
	n := len(points)
	if n < 3 {
		// Un polygone doit avoir au moins 3 points pour former une figure fermée.
		return 0.0
	}

	aire := 0

	// Utiliser l'algorithme de l'aire signée pour calculer l'aire du polygone.
	for i := 0; i < n; i++ {
		x1, y1 := points[i].X, points[i].Y
		x2, y2 := points[(i+1)%n].X, points[(i+1)%n].Y
		aire += (x1*y2 - x2*y1)
	}

	// Diviser par 2 et prendre la valeur absolue pour obtenir l'aire totale.
	aire = int(0.5 * float64(aire))
	if aire < 0 {
		aire = -aire
	}

	return aire
}

func main() {
	fmt.Println(part1("input2.txt"))
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
	loop = append(loop, utils.Pos{1, 1})
	loop = Loop(grid, utils.Pos{2, 1}, "right", loop)

	farest := 0
	for _, u := range loop {
		if farest < abs(u.X-1)+(u.Y-1) {
			farest = (u.X - 1) + (u.Y - 1)
		}
	}

	polygon := make([]utils.Pos, 0, 10)
	for _, u := range loop {
		if grid[u] != "|" && grid[u] != "-" {
			polygon = append(polygon, u)
		}
	}
	x := true
	y := false
	for i := 1; i < len(polygon); i++ {
		if x {
			polygon[i].X++
			x = false
		} else if y {
			polygon[i].Y++
			y = false
		}
	}

	fmt.Println(farest)
	fmt.Println(Aire(polygon))
	return 0
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
