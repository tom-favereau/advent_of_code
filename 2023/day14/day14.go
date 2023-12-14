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

	//grid := utils.NewGrid[string]()
	r := make(map[int]int)
	j := 0
	res := 0
	nb := 0
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, "")
		for i, u := range arr {
			if u == "#" {
				r[i] = j + 1
			} else if u == "O" {
				res += r[i]
				r[i]++
				nb++
			} else if u == "." {
				_, b := r[i]
				if !b {
					r[i] = j
				}
			}
		}
		j++
	}

	return (nb * j) - res
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
	i := 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")

		j = 0
		for _, u := range arr {
			grid[utils.Pos{j, i}] = u
			j++
		}
		i++
	}
	/*
		cycle(grid, j, i, 0)
		cycle(grid, j, i, 1)
		cycle(grid, j, i, 2)
		r := cycle(grid, j, i, 3).res
		fmt.Println("cycle 1")
		fmt.Println(toArray(grid, j, i))
		fmt.Println(r)
		cycle(grid, j, i, 0)
		cycle(grid, j, i, 1)
		cycle(grid, j, i, 2)
		r = cycle(grid, j, i, 3).res
		fmt.Println("cycle 2")
		fmt.Println(toArray(grid, j, i))
		fmt.Println(r)
		cycle(grid, j, i, 0)
		cycle(grid, j, i, 1)
		cycle(grid, j, i, 2)
		r = cycle(grid, j, i, 3).res
		fmt.Println("cycle 3")
		fmt.Println(toArray(grid, j, i))
		fmt.Println(r)

	*/

	res := 0
	memo := make(map[[100][100]string]int)
	list_res := make([]int, 1)
	for k := 1; k <= 1000000000; k++ {
		cycle(grid, j, i, 0)
		//fmt.Println("r0", calcul(grid, j, i))
		cycle(grid, j, i, 1)
		//fmt.Println("r1", calcul(grid, j, i))
		cycle(grid, j, i, 2)
		//fmt.Println("r2", calcul(grid, j, i))
		cycle(grid, j, i, 3)
		//fmt.Println("r3", calcul(grid, j, i))
		res = calcul(grid, j, i)
		isVisit, b := memo[toArray(grid, j, i)]

		if b {
			index_res := (1000000000 - isVisit) % (k - isVisit)
			return list_res[index_res+isVisit]
		} else {
			memo[toArray(grid, j, i)] = k
			list_res = append(list_res, res)
		}

	}
	return res

	//return res[len(res)-1]
}

type Return struct {
	g   utils.Grid[string]
	res int
}

func cycle(g utils.Grid[string], mx int, my int, c int) Return {
	if c == 0 { //haut
		r := make(map[int]int)
		res := 0
		nb := 0
		for j := 0; j < my; j++ { //haut
			for i := 0; i < mx; i++ {
				u := g[utils.Pos{i, j}]
				if u == "#" {
					r[i] = j + 1
				} else if u == "O" {
					g[utils.Pos{i, j}] = "."
					g[utils.Pos{i, r[i]}] = "O"
					res += r[i]
					r[i]++
					nb++
				} else if u == "." {
					_, b := r[i]
					if !b {
						r[i] = j
					}
				}
			}
		}
		re := Return{g, (nb * mx) - res}
		return re
	} else if c == 1 { //gauche
		//fmt.Println(g)
		//fmt.Println("/////")
		r := make(map[int]int)
		res := 0
		nb := 0
		for j := 0; j < mx; j++ {
			for i := 0; i < my; i++ {
				u := g[utils.Pos{j, i}]
				if u == "#" {
					r[i] = j + 1
				} else if u == "O" {
					g[utils.Pos{j, i}] = "."
					//fmt.Println(i, j, r[i])
					g[utils.Pos{r[i], i}] = "O"
					res += i
					r[i]++
					nb++
				} else if u == "." {
					_, b := r[i]
					if !b {
						r[i] = j
					}
				}
			}
		}
		re := Return{g, (nb * mx) - res}
		return re
	} else if c == 2 { //bas
		r := make(map[int]int)
		res := 0
		nb := 0
		for j := my - 1; j >= 0; j-- { //haut
			for i := 0; i < mx; i++ {
				u := g[utils.Pos{i, j}]
				if u == "#" {
					r[i] = j - 1
				} else if u == "O" {
					_, b := r[i]
					if !b {
						r[i] = 9
					}
					g[utils.Pos{i, j}] = "."
					g[utils.Pos{i, r[i]}] = "O"
					res += r[i]
					r[i]--
					nb++
				} else if u == "." {
					_, b := r[i]
					if !b {
						r[i] = j
					}
				}
			}
		}
		re := Return{g, res}
		return re
	} else if c == 3 { //droite
		r := make(map[int]int)
		res := 0
		nb := 0
		for j := mx - 1; j >= 0; j-- {
			for i := 0; i < my; i++ {
				u := g[utils.Pos{j, i}]
				if u == "#" {
					r[i] = j - 1
				} else if u == "O" {
					_, b := r[i]
					if !b {
						r[i] = 9
					}
					g[utils.Pos{j, i}] = "."
					//fmt.Println(i, j, r[i])
					g[utils.Pos{r[i], i}] = "O"
					res += i
					r[i]--
					nb++
				} else if u == "." {
					_, b := r[i]
					if !b {
						r[i] = j
					}
				}
			}
		}
		re := Return{g, (nb * mx) - res}
		return re
	}
	re := Return{g, 0}
	return re
}

func toArray(grid utils.Grid[string], mx, my int) [100][100]string {
	var arr [100][100]string
	for i := 0; i < mx; i++ {
		for j := 0; j < my; j++ {
			arr[j][i] = grid[utils.Pos{i, j}]
		}
	}
	return arr
}

func contain(t []int, val int) int {
	for i, u := range t {
		if u == val {
			return i
		}
	}
	return -1
}

func calcul(g utils.Grid[string], nx, ny int) int {
	res := 0
	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			if g[utils.Pos{i, j}] == "O" {
				res += nx - j
			}
		}
	}
	return res
}
