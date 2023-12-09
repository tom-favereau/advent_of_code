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
	fmt.Println(part2("input.txt"))
	//fmt.Println(part2("input2.txt"))
}

type Choix struct {
	g string
	d string
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	instruction := strings.Split(line, "")

	scanner.Scan()

	m := make(map[string]Choix)

	for scanner.Scan() {
		line = scanner.Text()
		key := line[:3]
		one := line[7:10]
		two := line[12:15]
		m[key] = Choix{g: one, d: two}
	}

	count := 0
	i := 0
	where := "AAA"
	for {
		ins := instruction[i]
		if ins == "L" {
			where = m[where].g
		} else {
			where = m[where].d
		}

		count++
		if where == "ZZZ" {
			break
		}

		if i == len(instruction)-1 {
			i = 0
		} else {
			i++
		}
	}

	return count

}

func part2(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	instruction := strings.Split(line, "")

	scanner.Scan()

	m := make(map[string]Choix)
	where := make([]string, 0, 10)
	for scanner.Scan() {
		line = scanner.Text()
		key := line[:3]
		one := line[7:10]
		two := line[12:15]
		m[key] = Choix{g: one, d: two}
		if key[2] == 'A' {
			where = append(where, key)
		}
	}

	count := make([]int, len(where))
	is_visit := make(map[int]bool)
	i := 0
	res := 0
	for {
		ins := instruction[i]
		for j, _ := range where {
			if ins == "L" {
				where[j] = m[where[j]].g
			} else {
				where[j] = m[where[j]].d
			}
			if !is_visit[j] {
				count[j]++
			}
			if where[j][2] == 'Z' {
				is_visit[j] = true
			}
		}
		if len(where) == len(is_visit) {
			res = lcmOfSlice(count)
			return res
		}

		if i == len(instruction)-1 {
			i = 0
		} else {
			i++
		}
	}

	return res

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
		fmt.Println("La liste doit contenir au moins deux éléments.")
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}
