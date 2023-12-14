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
	fmt.Println(part2("input2.txt"))
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

	res := make([]int, 0)
	isCyle := make(map[int]int)
	k := 0
	for {
		cycle(grid, j, i, 0)
		cycle(grid, j, i, 1)
		cycle(grid, j, i, 2)
		t := cycle(grid, j, i, 3)
		tmp := t.res

		fmt.Println(t)
		fmt.Println("///////")

		a, b := isCyle[tmp]
		if b {
			t := 1000000000 % (k - a)
			return res[t+a]
		} else {
			isCyle[tmp] = k
			res = append(res, tmp)
		}
		k++

	}

	/*
		res := 0
		for k := 0; k < 1000000000; k++ {
			res = cycle(grid, j, i, 0).res
			res = cycle(grid, j, i, 1).res
			res = cycle(grid, j, i, 2).res
			res = cycle(grid, j, i, 3).res
		}
		return res

	*/
	//return res[len(res)-1]
}

type Return struct {
	g   utils.Grid[string]
	res int
}

func cycle(g utils.Grid[string], mx int, my int, c int) Return {
	if c == 0 {
		r := make(map[int]int)
		res := 0
		nb := 0
		for j := 0; j < my; j++ {
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
	} else if c == 1 {
		r := make(map[int]int)
		res := 0
		nb := 0
		for i := 0; i < mx; i++ {
			for j := 0; j < my; j++ {
				u := g[utils.Pos{i, j}]
				if u == "#" {
					r[i] = j + 1
				} else if u == "O" {
					g[utils.Pos{i, j}] = "."
					g[utils.Pos{i, r[i]}] = "O"
					res += j
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
	} else if c == 2 {
		r := make(map[int]int)
		res := 0
		nb := 0
		for j := my - 1; j >= 0; j-- {
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
		re := Return{g, res}
		return re
	} else if c == 3 {
		r := make(map[int]int)
		res := 0
		nb := 0
		for i := mx - 1; i >= 0; i-- {
			for j := 0; j < my; j++ {
				u := g[utils.Pos{i, j}]
				if u == "#" {
					r[i] = j + 1
				} else if u == "O" {
					g[utils.Pos{i, j}] = "."
					g[utils.Pos{i, r[i]}] = "O"
					res += j
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
	}
	re := Return{g, 0}
	return re
}
