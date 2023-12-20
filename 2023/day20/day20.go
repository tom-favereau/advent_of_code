package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	//fmt.Println("///")
	fmt.Println(part2("input.txt"))
	// du graphe, valeur : 3739 3821 3943 4001
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

	m := make(map[string]Module)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " -> ")
		if arr[0] == "broadcaster" {
			m["broadcaster"] = Module{"", strings.Split(arr[1], ", "), false, nil}
		} else {
			tmp := strings.Split(arr[0], "")
			memo := make(map[string]bool)
			m[arr[0][1:]] = Module{tmp[0], strings.Split(arr[1], ", "), false, memo}
		}
	}

	for i, u := range m {
		for _, v := range u.next {
			if m[v].t == "&" {
				m[v].mem[i] = false
			}
		}
	}

	hight := 0
	low := 0
	for i := 0; i < 1000; i++ {
		tmp := solve(m)
		hight += tmp.hight
		low += tmp.low
	}
	return low * hight
}

func part2(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[string]Module)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " -> ")
		if arr[0] == "broadcaster" {
			m["broadcaster"] = Module{"", strings.Split(arr[1], ", "), false, nil}
		} else {
			tmp := strings.Split(arr[0], "")
			memo := make(map[string]bool)
			m[arr[0][1:]] = Module{tmp[0], strings.Split(arr[1], ", "), false, memo}
		}
	}

	for i, u := range m {
		for _, v := range u.next {
			if m[v].t == "&" {
				m[v].mem[i] = false
			}
		}
	}

	res := solveP2(m)
	return res
}

// modifier l'ordre dans lequel on fait les opération ....

type Info struct {
	name  string
	pulse string
	from  string
}

type Pair struct {
	hight int
	low   int
}

func solve(m map[string]Module) Pair {
	hight := 0
	low := 1
	queue := make([]Info, 0, 1)
	queue = append(queue, Info{"broadcaster", "low", ""})
	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]
		key := val.name
		pulse := val.pulse
		from := val.from
		mod := m[key]
		if mod.t == "" {
			for _, u := range mod.next {
				queue = append(queue, Info{u, "low", key})
			}
			low += len(mod.next)
		} else if mod.t == "%" {
			if pulse == "hight" {

			} else {
				if mod.on {
					m[key] = Module{mod.t, mod.next, false, mod.mem}
					for _, u := range mod.next {
						queue = append(queue, Info{u, "low", key})
					}
					low += len(mod.next)
				} else {
					m[key] = Module{mod.t, mod.next, true, mod.mem}
					for _, u := range mod.next {
						queue = append(queue, Info{u, "hight", key})
					}
					hight += len(mod.next)
				}

			}
		} else if mod.t == "&" {
			if pulse == "low" {
				mod.mem[from] = false
			} else {
				mod.mem[from] = true
			}
			tmp := true
			for _, u := range mod.mem {
				tmp = tmp && u
			}
			if tmp {
				for _, u := range mod.next {
					queue = append(queue, Info{u, "low", key})
				}
				low += len(mod.next)
			} else {
				for _, u := range mod.next {
					queue = append(queue, Info{u, "hight", key})
				}
				hight += len(mod.next)
			}
		}
	}
	return Pair{hight: hight, low: low}
}

func solveP2(m map[string]Module) int {
	i1 := 0
	i2 := 0
	i3 := 0
	i4 := 0
	for index := 1; index < 10000000; index++ {
		if i1 != 0 && i2 != 0 && i3 != 0 && i4 != 0 {
			break
		}
		hight := 0
		low := 1
		queue := make([]Info, 0, 1)
		queue = append(queue, Info{"broadcaster", "low", ""})
		for len(queue) > 0 {
			val := queue[0]
			queue = queue[1:]
			key := val.name
			pulse := val.pulse
			from := val.from
			mod := m[key]
			if mod.t == "" {
				for _, u := range mod.next {
					queue = append(queue, Info{u, "low", key})
				}
				low += len(mod.next)
			} else if mod.t == "%" {
				if pulse == "hight" {

				} else {
					if mod.on {
						m[key] = Module{mod.t, mod.next, false, mod.mem}
						for _, u := range mod.next {
							queue = append(queue, Info{u, "low", key})
						}
						low += len(mod.next)
					} else {
						m[key] = Module{mod.t, mod.next, true, mod.mem}
						for _, u := range mod.next {
							queue = append(queue, Info{u, "hight", key})
						}
						hight += len(mod.next)
					}

				}
			} else if mod.t == "&" {
				if pulse == "low" {
					mod.mem[from] = false
				} else {
					mod.mem[from] = true
				}
				tmp := true
				for _, u := range mod.mem {
					tmp = tmp && u
				}
				if tmp {
					if key == "mj" && i1 == 0 {
						//fmt.Println("mj", index)
						i1 = index
					} else if key == "rd" && i2 == 0 {
						//fmt.Println("rd", index)
						i2 = index
					} else if key == "qs" && i3 == 0 {
						//fmt.Println("qs", index)
						i3 = index
					} else if key == "cs" && i4 == 0 {
						//fmt.Println("cs", index)
						i4 = index
					}
					for _, u := range mod.next {
						queue = append(queue, Info{u, "low", key})
					}
					low += len(mod.next)
				} else {
					for _, u := range mod.next {
						queue = append(queue, Info{u, "hight", key})
					}
					hight += len(mod.next)
				}
			}
		}
	}
	return i1 * i2 * i3 * i4
}

type Module struct {
	t    string
	next []string
	on   bool            //si on est on ou pas
	mem  map[string]bool // && de tout les élément
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func lcmOfSlice(numbers []int) int {
	if len(numbers) < 2 {
		return numbers[0]
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}
