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

	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")

		arrnum := make([]int, 0, 10)
		for _, u := range arr {
			num, _ := strconv.Atoi(u)
			arrnum = append(arrnum, num)
		}
		res += result(arrnum)
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

	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")

		arrnum := make([]int, 0, 10)
		for _, u := range arr {
			num, _ := strconv.Atoi(u)
			arrnum = append(arrnum, num)
		}
		res += result2(arrnum)
	}

	return res
}

func result(arr []int) int {
	zero := true
	tmp := make([]int, 0, 10)
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff != 0 {
			zero = false
		}
		tmp = append(tmp, diff)
	}

	if zero {
		return arr[len(arr)-1]
	} else {
		return arr[len(arr)-1] + result(tmp)
	}
}

func result2(arr []int) int {
	zero := true
	tmp := make([]int, 0, 10)
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff != 0 {
			zero = false
		}
		tmp = append(tmp, diff)
	}

	if zero {
		//fmt.Println("ok")
		return arr[0]
	} else {
		//fmt.Println("arr", arr[0])
		//fmt.Println("test", arr[0], result2(tmp))
		return arr[0] - result2(tmp)
	}
}
