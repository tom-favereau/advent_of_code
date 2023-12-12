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

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " ")
		arr := strings.Split(sep[0], "")
		numstring := strings.Split(sep[1], ",")
		num := make([]int, 0, 10)
		for _, u := range numstring {
			n, _ := strconv.Atoi(u)
			num = append(num, n)
		}

		res += f(arr, num, 0, true)

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
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " ")
		arrtmp := strings.Split(sep[0], "")
		numstring := strings.Split(sep[1], ",")
		numtmp := make([]int, 0, 10)

		for _, u := range numstring {
			n, _ := strconv.Atoi(u)
			numtmp = append(numtmp, n)
		}

		num := make([]int, 0, 10)
		arr := make([]string, 0, 10)
		for i := 0; i < 5; i++ {
			arr = append(arr, arrtmp...)
			if i != 4 {
				arr = append(arr, "?")
			}
			num = append(num, numtmp...)
		}

		res += f(arr, num, 0, true)
		fmt.Println(i)
		i++

	}

	return res
}

func traitement(arr []string, num []int, tmp int) int {
	if len(num) == 0 {
		for _, u := range arr {
			if u != "." {
				return 0
			}
		}
		return 1
	} else {
		res := 0
		for i, u := range arr {
			if u == "#" {
				tmp++
				if tmp > num[0] {
					return 0
				} else if len(arr)-1 == i {
					return 1
				}
			} else if u == "?" {
				copied1 := make([]string, len(arr))
				copied2 := make([]string, len(arr))
				copy(copied1, arr)
				copy(copied2, arr)
				if tmp != 0 && tmp < num[0] {
					copied1[i] = "#"
					res += traitement(copied1[i:], num, tmp)
				} else if tmp == num[0] {
					if i < len(arr)-1 {
						res += traitement(copied1[i+1:], num[1:], 0)
					} else {
						res++
					}
					return res
				} else {
					copied1[i] = "#"
					res += traitement(copied1[i:], num, tmp)
					copied2[i] = "."
					res += traitement(copied2[i:], num, 0)
				}
				return res
			} else if u == "." {
				if tmp != 0 && tmp < num[0] {
					return 0
				} else if i < len(arr)-1 {
					return traitement(arr[i+1:], num, 0)
				} else {
					return 1
				}
			}
		}
		return res
	}
}

func f(arr []string, num []int, tmp int, first bool) int {
	if len(arr) == 0 {
		if len(num) == 0 {
			return 1
		} else {
			return 0
		}
	} else if len(num) == 0 {
		for _, u := range arr {
			if u == "#" {
				return 0
			}
		}
		return 1
	} else {
		if arr[0] == "#" {
			if tmp+1 > num[0] {
				return 0
			} else {
				if len(arr) == 1 {
					if len(num) == 1 && num[0] == tmp+1 {
						return 1
					} else {
						return 0
					}
				} else {
					return f(arr[1:], num, tmp+1, false)
				}
			}
		} else if arr[0] == "." {
			if first {
				return f(arr[1:], num, 0, first)
			} else {
				if tmp < num[0] {
					return 0
				} else {
					return f(arr[1:], num[1:], 0, true)
				}
			}
		} else if arr[0] == "?" {
			c1 := make([]string, len(arr))
			c2 := make([]string, len(arr))
			copy(c1, arr)
			c1[0] = "#"
			copy(c2, arr)
			c2[0] = "."
			r1 := f(c1, num, tmp, false)
			r2 := f(c2, num, tmp, first)
			return r1 + r2
		}
	}
	return tmp
}
