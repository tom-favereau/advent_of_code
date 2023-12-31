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

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	index := 1
	possible := true
	for scanner.Scan() {
		line := scanner.Text()

		arrayLine := strings.Split(line[7:], ";")

		for _, u := range arrayLine {
			possible = true
			red := 0
			blue := 0
			green := 0
			part := strings.Split(u, ",")
			for _, v := range part {
				tmp := strings.Split(v, " ")
				if tmp[2] == "green" {
					number, _ := strconv.Atoi(tmp[1])
					green += number
				} else if tmp[2] == "blue" {
					number, _ := strconv.Atoi(tmp[1])
					blue += number
				} else if tmp[2] == "red" {
					number, _ := strconv.Atoi(tmp[1])
					red += number
				}
			}
			if red > 12 || green > 13 || blue > 14 {
				possible = false
				break
			}
		}
		if possible {
			res += index
		}
		index++

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

		arrayLine := strings.Split(line[7:], ";")

		redmax := 0
		bluemax := 0
		greenmax := 0
		for _, u := range arrayLine {
			red := 0
			blue := 0
			green := 0
			part := strings.Split(u, ",")
			for _, v := range part {
				tmp := strings.Split(v, " ")
				if tmp[2] == "green" {
					number, _ := strconv.Atoi(tmp[1])
					green += number
				} else if tmp[2] == "blue" {
					number, _ := strconv.Atoi(tmp[1])
					blue += number
				} else if tmp[2] == "red" {
					number, _ := strconv.Atoi(tmp[1])
					red += number
				}
			}
			if green > greenmax {
				greenmax = green
			}
			if blue > bluemax {
				bluemax = blue
			}
			if red > redmax {
				redmax = red
			}
		}
		res += greenmax * bluemax * redmax

	}

	return res
}
