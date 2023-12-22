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
	fmt.Println(part1("input.txt"))
	//fmt.Println("///")
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

	order := make(Order, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "~")
		S := strings.Split(arr[0], ",")
		x0, _ := strconv.Atoi(S[0])
		y0, _ := strconv.Atoi(S[1])
		z0, _ := strconv.Atoi(S[2])
		E := strings.Split(arr[1], ",")
		x1, _ := strconv.Atoi(E[0])
		y1, _ := strconv.Atoi(E[1])
		z1, _ := strconv.Atoi(E[2])
		order = append(order, Brick{x0, y0, z0, x1, y1, z1})
	}

	sort.Sort(order)

	m := make(map[P]Brick)
	mapSupport := make(map[Brick][]Brick)
	mapSupportReci := make(map[Brick][]Brick)
	for _, u := range order {
		brickSupport := Brick{0, 0, 0, 0, 0, -1}
		support := make([]Brick, 0)
		for x := u.x0; x <= u.x1; x++ {
			for y := u.y0; y <= u.y1; y++ {
				bs, b := m[P{x, y}]
				if b && bs.z1 > brickSupport.z1 {
					if bs.z1 != -1 {
						support = make([]Brick, 0)
						support = append(support, bs)
					}
					brickSupport = m[P{x, y}]
				} else if b && bs.z1 == brickSupport.z1 && bs != brickSupport {
					support = append(support, bs)
				}
			}
		}
		for x := u.x0; x <= u.x1; x++ {
			for y := u.y0; y <= u.y1; y++ {
				m[P{x, y}] = Brick{u.x0, u.y0, brickSupport.z1 + 1, u.x1, u.y1, brickSupport.z1 + 1 + u.z1 - u.z0}
			}
		}
		for _, v := range support {
			mapSupport[v] = append(mapSupport[v], u)
		}
		mapSupportReci[u] = support

	}

	//counted := make(map[Brick]bool)
	res := len(order) - len(mapSupport) //ce qui sont tout en haut
	for i, u := range mapSupport {
		sup := true
		for _, v := range u {
			tmp := false
			for j, g := range mapSupport {
				for _, f := range g {
					if v == f && i != j {
						tmp = true
					}
				}
			}
			sup = sup && tmp
		}
		if sup {
			res++
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

	order := make(Order, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "~")
		S := strings.Split(arr[0], ",")
		x0, _ := strconv.Atoi(S[0])
		y0, _ := strconv.Atoi(S[1])
		z0, _ := strconv.Atoi(S[2])
		E := strings.Split(arr[1], ",")
		x1, _ := strconv.Atoi(E[0])
		y1, _ := strconv.Atoi(E[1])
		z1, _ := strconv.Atoi(E[2])
		order = append(order, Brick{x0, y0, z0, x1, y1, z1})
	}

	sort.Sort(order)

	m := make(map[P]Brick)
	mapSupport := make(map[Brick][]Brick)
	mapSupportReci := make(map[Brick][]Brick)
	for _, u := range order {
		brickSupport := Brick{0, 0, 0, 0, 0, -1}
		support := make([]Brick, 0)
		for x := u.x0; x <= u.x1; x++ {
			for y := u.y0; y <= u.y1; y++ {
				bs, b := m[P{x, y}]
				if b && bs.z1 > brickSupport.z1 {
					if bs.z1 != -1 {
						support = make([]Brick, 0)
						support = append(support, bs)
					}
					brickSupport = m[P{x, y}]
				} else if b && bs.z1 == brickSupport.z1 && bs != brickSupport {
					support = append(support, bs)
				}
			}
		}
		for x := u.x0; x <= u.x1; x++ {
			for y := u.y0; y <= u.y1; y++ {
				m[P{x, y}] = Brick{u.x0, u.y0, brickSupport.z1 + 1, u.x1, u.y1, brickSupport.z1 + 1 + u.z1 - u.z0}
			}
		}
		for _, v := range support {
			mapSupport[v] = append(mapSupport[v], Brick{u.x0, u.y0, brickSupport.z1 + 1, u.x1, u.y1, brickSupport.z1 + 1 + u.z1 - u.z0})
		}
		mapSupportReci[Brick{u.x0, u.y0, brickSupport.z1 + 1, u.x1, u.y1, brickSupport.z1 + 1 + u.z1 - u.z0}] = support

	}
	res := 0

	order = make([]Brick, 0)
	for j, _ := range mapSupportReci {
		order = append(order, j)
	}
	sort.Sort(order)

	for i, _ := range mapSupport {
		fall := make([]Brick, 0)
		fall = append(fall, i)
		for _, j := range order {
			v := mapSupportReci[j]
			tmp := 0
			for _, f := range v {
				for _, g := range fall {
					if g == f && j != f {
						tmp++
					}
				}
			}
			if tmp == len(v) && len(v) != 0 {
				fall = append(fall, j)
			}
		}
		res += len(fall) - 1
	}

	return res

}

type P struct {
	x, y int
}

type Brick struct {
	x0, y0, z0 int
	x1, y1, z1 int
}

type Order []Brick

func (o Order) Len() int      { return len(o) }
func (o Order) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

func (o Order) Less(i, j int) bool {
	// Compare les valeurs de z1
	if o[i].z0 < o[j].z0 {
		return true
	} else if o[i].z0 == o[j].z0 {
		// En cas d'égalité de z1, compare les valeurs de z2
		return o[i].z1 < o[j].z1
	}
	// Sinon, o[i].z1 > o[j].z1, donc retourne false
	return false
}
