package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")

		tmp := make([]string, 0, len(arr))
		for _, u := range arr {
			tmp = append(tmp, u)
		}

		g = append(g, tmp)
	}

	m := make(map[P]bool)
	info := make(map[Info]bool)

	f(g, P{0, 0}, "right", m, info)

	res := 0
	for _, u := range m {
		if u {
			res++
		}
	}

	return res
}

type P struct {
	i int
	j int
}

type Info struct {
	p   P
	dir string
}

func f(g [][]string, p P, dir string, ener map[P]bool, info map[Info]bool) {
	if p.i >= len(g) || p.j >= len(g[0]) || p.i < 0 || p.j < 0 {
		return
	} else if info[Info{p, dir}] {
		return
	} else {
		ener[p] = true
		info[Info{p, dir}] = true
		//fmt.Println(p.i, p.j)
		u := g[p.i][p.j]
		if u == "." {
			if dir == "up" {
				f(g, P{p.i - 1, p.j}, dir, ener, info)
			} else if dir == "down" {
				f(g, P{p.i + 1, p.j}, dir, ener, info)
			} else if dir == "left" {
				f(g, P{p.i, p.j - 1}, dir, ener, info)
			} else if dir == "right" {
				f(g, P{p.i, p.j + 1}, dir, ener, info)
			}
		} else if u == "|" {
			if dir == "up" {
				f(g, P{p.i - 1, p.j}, dir, ener, info)
			} else if dir == "down" {
				f(g, P{p.i + 1, p.j}, dir, ener, info)
			} else {
				f(g, P{p.i + 1, p.j}, "down", ener, info)
				f(g, P{p.i - 1, p.j}, "up", ener, info)
			}
		} else if u == "-" {
			if dir == "left" {
				f(g, P{p.i, p.j - 1}, dir, ener, info)
			} else if dir == "right" {
				f(g, P{p.i, p.j + 1}, dir, ener, info)
			} else {
				f(g, P{p.i, p.j - 1}, "left", ener, info)
				f(g, P{p.i, p.j + 1}, "right", ener, info)
			}
		} else if u == "/" {
			if dir == "up" {
				f(g, P{p.i, p.j + 1}, "right", ener, info)
			} else if dir == "down" {
				f(g, P{p.i, p.j - 1}, "left", ener, info)
			} else if dir == "right" {
				f(g, P{p.i - 1, p.j}, "up", ener, info)
			} else if dir == "left" {
				f(g, P{p.i + 1, p.j}, "down", ener, info)
			}
		} else if u == "\\" {
			if dir == "down" {
				f(g, P{p.i, p.j + 1}, "right", ener, info)
			} else if dir == "up" {
				f(g, P{p.i, p.j - 1}, "left", ener, info)
			} else if dir == "left" {
				f(g, P{p.i - 1, p.j}, "up", ener, info)
			} else if dir == "right" {
				f(g, P{p.i + 1, p.j}, "down", ener, info)
			}
		}
	}
	return
}

func part2(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")

		tmp := make([]string, 0, len(arr))
		for _, u := range arr {
			tmp = append(tmp, u)
		}

		g = append(g, tmp)
	}
	res := 0
	for i := 0; i < len(g); i++ {
		m := make(map[P]bool)
		info := make(map[Info]bool)

		f(g, P{i, 0}, "right", m, info)

		tmp := 0
		for _, u := range m {
			if u {
				tmp++
			}
		}
		if tmp > res {
			res = tmp
		}
	}

	for i := 0; i < len(g); i++ {
		m := make(map[P]bool)
		info := make(map[Info]bool)

		f(g, P{i, len(g[0]) - 1}, "left", m, info)

		tmp := 0
		for _, u := range m {
			if u {
				tmp++
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	for j := 0; j < len(g[0]); j++ {
		m := make(map[P]bool)
		info := make(map[Info]bool)

		f(g, P{0, j}, "down", m, info)

		tmp := 0
		for _, u := range m {
			if u {
				tmp++
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	for j := 0; j < len(g[0]); j++ {
		m := make(map[P]bool)
		info := make(map[Info]bool)

		f(g, P{len(g) - 1, j}, "up", m, info)

		tmp := 0
		for _, u := range m {
			if u {
				tmp++
			}
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}
