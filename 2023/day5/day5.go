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
	fmt.Println(part2("input.txt"))

}

func part2(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := make([]int, 0, 10)
	scanner.Scan()
	line := scanner.Text()

	for _, u := range strings.Split(line[6:], " ") {
		if u != "" {
			num, _ := strconv.Atoi(u)
			seeds = append(seeds, num)
		}
	}

	//m := make(map[int]int)
	res := make([]Intervale, 0, 10)
	for i := 0; i < len(seeds); i += 2 {
		res = append(res, Intervale{a: seeds[i], b: seeds[i] + seeds[i+1] - 1})
	}

	scanner.Scan()
	scanner.Scan()
	tmp := make([]Intervale, 0, 10)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			//m = make(map[int]int)
			for _, u := range res {
				if u.a != -1 {
					tmp = append(tmp, u)
				}
			}
			res = tmp
			tmp = make([]Intervale, 0, 10)
			scanner.Scan()
		} else {
			arr := strings.Split(line, " ")
			dest, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
			source, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
			step, _ := strconv.Atoi(strings.TrimSpace(arr[2]))
			for i, u := range res {
				if u.a >= source && u.b < source+step {
					tmp = append(tmp, Intervale{a: dest + (u.a - source), b: dest + (u.b - source)})
					res[i] = Intervale{-1, -1}
				} else if u.b < source {

				} else if u.a >= source+step {

				} else if u.a < source && u.b < source+step {
					tmp = append(tmp, Intervale{a: dest, b: dest + (u.b - source)})
					res[i].b = source - 1
				} else if u.a >= source && u.b >= source+step {
					tmp = append(tmp, Intervale{a: dest + (u.a - source), b: dest + step - 1})
					res[i].a = source + step
				}
			}
		}
	}
	for _, u := range res {
		if u.a != -1 {
			tmp = append(tmp, u)
		}
	}
	res = tmp
	t := res[0].a
	for _, u := range res {
		if u.a < t && u.a > 0 {
			t = u.a
		}
	}
	return t
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := make([]int, 0, 10)
	scanner.Scan()
	line := scanner.Text()

	for _, u := range strings.Split(line[6:], " ") {
		if u != "" {
			num, _ := strconv.Atoi(u)
			seeds = append(seeds, num)
		}
	}

	//m := make(map[int]int)
	res := make([]Intervale, 0, 10)
	for i := 0; i < len(seeds); i++ {
		res = append(res, Intervale{a: seeds[i], b: seeds[i]})
	}

	scanner.Scan()
	scanner.Scan()
	tmp := make([]Intervale, 0, 10)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			//m = make(map[int]int)
			for _, u := range res {
				if u.a != -1 {
					tmp = append(tmp, u)
				}
			}
			res = tmp
			tmp = make([]Intervale, 0, 10)
			scanner.Scan()
		} else {
			arr := strings.Split(line, " ")
			dest, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
			source, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
			step, _ := strconv.Atoi(strings.TrimSpace(arr[2]))
			for i, u := range res {
				if u.a >= source && u.b < source+step {
					tmp = append(tmp, Intervale{a: dest + (u.a - source), b: dest + (u.b - source)})
					res[i] = Intervale{-1, -1}
				} else if u.b < source {

				} else if u.a >= source+step {

				} else if u.a < source && u.b < source+step {
					tmp = append(tmp, Intervale{a: dest, b: dest + (u.b - source)})
					res[i].b = source - 1
				} else if u.a >= source && u.b >= source+step {
					tmp = append(tmp, Intervale{a: dest + (u.a - source), b: dest + step - 1})
					res[i].a = source + step
				}
			}
		}
	}
	for _, u := range res {
		if u.a != -1 {
			tmp = append(tmp, u)
		}
	}
	res = tmp
	t := res[0].a
	for _, u := range res {
		if u.a < t && u.a > 0 {
			t = u.a
		}
	}
	return t
}

type Intervale struct {
	a int
	b int
}
