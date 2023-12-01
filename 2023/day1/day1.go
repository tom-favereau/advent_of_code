package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()
}

func part1() {
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("erreur")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		arrayLine := strings.Split(line, "")
		first := -1
		end := -1
		for _, u := range arrayLine {
			number, er := strconv.Atoi(u)
			if er == nil {
				if first == -1 {
					first = number
				}
				end = number
			}
		}
		res += 10*first + end
	}
	fmt.Println(res)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("erreur")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		arrayLine := strings.Split(line, "")
		first := -1
		end := -1
		for i, u := range arrayLine {
			number, er := strconv.Atoi(u)
			if er == nil {
				if first == -1 {
					first = number
				}
				end = number
			} else if len(arrayLine)-i >= 3 && arrayLine[i] == "o" && arrayLine[i+1] == "n" && arrayLine[i+2] == "e" {
				if first == -1 {
					first = 1
				}
				end = 1
			} else if len(arrayLine)-i >= 3 && arrayLine[i] == "t" && arrayLine[i+1] == "w" && arrayLine[i+2] == "o" {
				if first == -1 {
					first = 2
				}
				end = 2
			} else if len(arrayLine)-i >= 5 && arrayLine[i] == "t" && arrayLine[i+1] == "h" && arrayLine[i+2] == "r" && arrayLine[i+3] == "e" && arrayLine[i+4] == "e" {
				if first == -1 {
					first = 3
				}
				end = 3
			} else if len(arrayLine)-i >= 4 && arrayLine[i] == "f" && arrayLine[i+1] == "o" && arrayLine[i+2] == "u" && arrayLine[i+3] == "r" {
				if first == -1 {
					first = 4
				}
				end = 4
			} else if len(arrayLine)-i >= 4 && arrayLine[i] == "f" && arrayLine[i+1] == "i" && arrayLine[i+2] == "v" && arrayLine[i+3] == "e" {
				if first == -1 {
					first = 5
				}
				end = 5
			} else if len(arrayLine)-i >= 3 && arrayLine[i] == "s" && arrayLine[i+1] == "i" && arrayLine[i+2] == "x" {
				if first == -1 {
					first = 6
				}
				end = 6
			} else if len(arrayLine)-i >= 5 && arrayLine[i] == "s" && arrayLine[i+1] == "e" && arrayLine[i+2] == "v" && arrayLine[i+3] == "e" && arrayLine[i+4] == "n" {
				if first == -1 {
					first = 7
				}
				end = 7
			} else if len(arrayLine)-i >= 5 && arrayLine[i] == "e" && arrayLine[i+1] == "i" && arrayLine[i+2] == "g" && arrayLine[i+3] == "h" && arrayLine[i+4] == "t" {
				if first == -1 {
					first = 8
				}
				end = 8
			} else if len(arrayLine)-i >= 4 && arrayLine[i] == "n" && arrayLine[i+1] == "i" && arrayLine[i+2] == "n" && arrayLine[i+3] == "e" {
				if first == -1 {
					first = 9
				}
				end = 9
			}
		}
		res += 10*first + end
	}
	fmt.Println(res)
}
