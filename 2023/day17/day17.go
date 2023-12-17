package main

import (
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

	grid := Grid{0, 0, make([][]int, 0)}
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, "")

		grid.cols = 0
		grid.values = append(grid.values, make([]int, 0))
		for _, u := range arr {
			num, _ := strconv.Atoi(u)
			grid.values[grid.rows] = append(grid.values[grid.rows], num)
			grid.cols++
		}
		grid.rows++
	}
	res := dijkstra(grid)
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

	grid := Grid{0, 0, make([][]int, 0)}
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, "")

		grid.cols = 0
		grid.values = append(grid.values, make([]int, 0))
		for _, u := range arr {
			num, _ := strconv.Atoi(u)
			grid.values[grid.rows] = append(grid.values[grid.rows], num)
			grid.cols++
		}
		grid.rows++
	}
	res := dijkstra2(grid)
	return res
}

type Grid struct {
	rows, cols int
	values     [][]int
}

type Info struct {
	row      int
	col      int
	longueur int
	dir      int
}

func dijkstra(grid Grid) int {
	rows, cols := grid.rows, grid.cols
	distances := make(map[Info]int)
	chemin := make(map[Info][]Test)

	//start := Node{0, 0, grid.values[0][0]}
	distances[Info{0, 0, 0, -1}] = 0

	queue := []Info{Info{0, 0, 0, -1}}
	directions := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for i, dir := range directions {
			if current.dir != -1 && ((current.longueur == 3 && current.dir == i) || current.dir == (i+2)%4) {

			} else {
				newRow, newCol := current.row+dir[0], current.col+dir[1]
				newLongueur := 0
				if current.dir == i {
					newLongueur = current.longueur + 1
				} else {
					newLongueur = 1
				}

				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
					newDistance := distances[current] + grid.values[newRow][newCol]

					pr, b := distances[Info{newRow, newCol, newLongueur, i}]
					if !b || newDistance < pr { //&& dir[0] != 0 && dir[1] != 0 {
						// Check if not more than 3 steps in the same direction
						distances[Info{newRow, newCol, newLongueur, i}] = newDistance
						chemin[Info{newRow, newCol, newLongueur, i}] = append(chemin[Info{current.row, current.col, current.longueur, current.dir}], Test{dir, newRow, newCol})
						queue = append(queue, Info{newRow, newCol, newLongueur, i})
					}
				}
			}
		}
	}
	//fmt.Println(chemin[rows-1][cols-1])
	res := 10000000000000
	//ch := make([]Test, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			tmp, b := distances[Info{rows - 1, cols - 1, j, i}]
			if b && tmp < res {
				res = tmp
				//ch = chemin[Info{rows - 1, cols - 1, j, i}]
			}
		}
	}
	//fmt.Println(ch)
	return res
}

type Test struct {
	dir [2]int
	i   int
	j   int
}

func dijkstra2(grid Grid) int {
	rows, cols := grid.rows, grid.cols
	distances := make(map[Info]int)
	//chemin := make(map[Info][]Test)

	//start := Node{0, 0, grid.values[0][0]}
	distances[Info{0, 0, 0, -1}] = 0

	queue := []Info{Info{0, 0, 0, -1}}
	directions := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for i, dir := range directions {
			if current.col == cols-1 && current.row == rows-1 {

			} else if current.dir != -1 && ((current.longueur < 4 && current.dir != i) || (current.longueur == 10 && current.dir == i) || ((i+2)%4 == current.dir)) {

			} else {
				newRow, newCol := current.row+dir[0], current.col+dir[1]
				newLongueur := 0
				if current.dir == i {
					newLongueur = current.longueur + 1
				} else {
					newLongueur = 1
				}

				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
					newDistance := distances[current] + grid.values[newRow][newCol]

					pr, b := distances[Info{newRow, newCol, newLongueur, i}]
					if !b || newDistance < pr { //&& dir[0] != 0 && dir[1] != 0 {
						// Check if not more than 3 steps in the same direction
						distances[Info{newRow, newCol, newLongueur, i}] = newDistance
						//chemin[Info{newRow, newCol, newLongueur, i}] = append(chemin[Info{newRow, newCol, newLongueur, i}], chemin[Info{current.row, current.col, current.longueur, current.dir}]...)
						//chemin[Info{newRow, newCol, newLongueur, i}] = append(chemin[Info{newRow, newCol, newLongueur, i}], Test{dir, newRow, newCol})
						queue = append(queue, Info{newRow, newCol, newLongueur, i})
					}
				}
			}
		}
	}
	//fmt.Println(chemin[rows-1][cols-1])
	res := 10000000000000
	//ch := make([]Test, 0)
	for i := 0; i < 4; i++ {
		for j := 4; j < 11; j++ {
			tmp, b := distances[Info{rows - 1, cols - 1, j, i}]
			//t, b2 := chemin[Info{rows - 1, cols - 1, j, i}]
			if b && tmp < res {
				//fmt.Println(tmp)
				//ch = chemin[Info{rows - 1, cols - 1, j, i}]
				res = tmp
			}
		}
	}
	//fmt.Println(ch)
	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
