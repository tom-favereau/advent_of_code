package main

import (
	utils "advent_of_code/utils/grid"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println("///")
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

	pos := make([]utils.Pos, 0)
	pos = append(pos, utils.Pos{0, 0})
	nb := 1
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		d, _ := strconv.Atoi(arr[1])
		nb += d
		pr := pos[len(pos)-1]
		if arr[0] == "R" {
			pos = append(pos, utils.Pos{pr.X + d, pr.Y})
		} else if arr[0] == "L" {
			pos = append(pos, utils.Pos{pr.X - d, pr.Y})
		} else if arr[0] == "U" {
			pos = append(pos, utils.Pos{pr.X, pr.Y + d})
		} else if arr[0] == "D" {
			pos = append(pos, utils.Pos{pr.X, pr.Y - d})
		}
	}

	aire := Aire(pos) - nb/2

	return aire + nb
}

func part2(string2 string) int64 {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := make([]P, 0)
	pos = append(pos, P{0, 0})
	var nb int64 = 1
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		hex := strings.Split(arr[2], "")
		//d := 0
		d, _ := strconv.ParseInt(arr[2][2:len(hex)-2], 16, 64)
		//fmt.Println(arr[2][1 : len(hex)-1])
		/*
			for i := len(hex) - 3; i >= 0; i-- {
				n, _ := strconv.Atoi(hex[i])
				//d += int(math.Pow(10, float64(i)) * float64(n))
			}

		*/
		nb += d
		pr := pos[len(pos)-1]
		last := hex[len(hex)-2]
		if last == "0" {
			pos = append(pos, P{pr.X + d, pr.Y})
		} else if last == "2" {
			pos = append(pos, P{pr.X - d, pr.Y})
		} else if last == "3" {
			pos = append(pos, P{pr.X, pr.Y + d})
		} else if last == "1" {
			pos = append(pos, P{pr.X, pr.Y - d})
		}
	}

	aire := Aire2(pos) - nb/2

	return aire + nb
}

type P struct {
	X int64
	Y int64
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

func Aire2(points []P) int64 {
	n := len(points)
	if n < 3 {
		return 0.0
	}

	var aire int64 = 0

	for i := 0; i < n; i++ {
		x1, y1 := points[i].X, points[i].Y
		x2, y2 := points[(i+1)%n].X, points[(i+1)%n].Y
		aire += (x1*y2 - x2*y1)
	}

	aire = int64(0.5 * float64(aire))
	if aire < 0 {
		aire = -aire
	}

	return aire
}
