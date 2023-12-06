package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	pars := strings.Split(line1, " ")
	T := make([]string, 0, 10)
	for i, u := range pars {
		if u != "" && i != 0 {
			T = append(T, u)
		}
	}
	pars = strings.Split(line2, " ")
	D := make([]string, 0, 10)
	for i, u := range pars {
		if u != "" && i != 0 {
			D = append(D, u)
		}
	}

	res := 1
	for i, _ := range D {
		// t*(T-t) - D = 0 => - t**2 + T*t - D = 0
		b, _ := strconv.Atoi(T[i])
		a := -1
		d, _ := strconv.Atoi(D[i])
		c := -(float64(d) + 0.0001)
		delta := math.Pow(float64(b), 2) - float64(4*a)*c
		x1 := int((-float64(b) - math.Sqrt(delta)) / (float64(2 * a)))
		x2 := int((-float64(b)+math.Sqrt(delta))/(float64(2*a))) + 1

		res *= int(math.Abs(float64(x2-x1))) + 1

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

	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	pars := strings.Split(line1, " ")
	parsT := make([]string, 0, 10)
	for i, u := range pars {
		if u != "" && i != 0 {
			parsT = append(parsT, u)
		}
	}
	pars = strings.Split(line2, " ")
	parsD := make([]string, 0, 10)
	for i, u := range pars {
		if u != "" && i != 0 {
			parsD = append(parsD, u)
		}
	}

	D := 0
	T := 0
	for i, _ := range parsD {
		if parsD[i] != "" {
			size := len(parsD[i])
			d, _ := strconv.Atoi(parsD[i])
			D = int(math.Pow(10, float64(size)))*D + d
		}
		if parsT[i] != "" {
			size := len(parsT[i])
			t, _ := strconv.Atoi(parsT[i])
			T = int(math.Pow(10, float64(size)))*T + t
		}
	}

	b := T
	a := -1
	d := D
	c := -(float64(d) + 0.0001)
	delta := math.Pow(float64(b), 2) - float64(4*a)*c
	x1 := int((-float64(b) - math.Sqrt(delta)) / (float64(2 * a)))
	x2 := int((-float64(b)+math.Sqrt(delta))/(float64(2*a))) + 1

	res := int(math.Abs(float64(x2-x1))) + 1

	return res
}
