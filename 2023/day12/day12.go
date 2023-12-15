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

		t := f(arr, num, 0, true)
		res += t

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

		opt := make(map[Info]int)

		res += f2(arr, num, 0, 0, 0, opt)
		//res += f3(arr, num, 0, true, opt)

	}
	return res
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

type Info struct {
	i1  int
	i2  int
	tmp int
}

func f2(arr []string, num []int, i1, i2, tmp int, opt map[Info]int) int {
	k := Info{i1, i2, tmp}

	val, ok := opt[k]
	if ok {
		return val
	}

	if i1 == len(arr) {
		if i2 == len(num) && tmp == 0 {
			return 1
		} else if i2 == len(num)-1 && num[i2] == tmp {
			return 1
		} else {
			return 0
		}
	}

	ans := 0

	if arr[i1] == "?" {
		if tmp == 0 {
			ans += f2(arr, num, i1+1, i2, 0, opt)
		} else if i2 < len(num) && num[i2] == tmp {
			ans += f2(arr, num, i1+1, i2+1, 0, opt)
		}
		ans += f2(arr, num, i1+1, i2, tmp+1, opt)

	} else if arr[i1] == "." {
		if tmp == 0 {
			ans += f2(arr, num, i1+1, i2, 0, opt)
		} else if i2 < len(num) && num[i2] == tmp {
			ans += f2(arr, num, i1+1, i2+1, 0, opt)
		}
	} else if arr[i1] == "#" {
		ans += f2(arr, num, i1+1, i2, tmp+1, opt)
	}

	opt[k] = ans
	return ans
}

func f3(arr []string, num []int, tmp int, first bool, opt map[Info]int) int {
	r, b := opt[Info{len(arr), len(num), tmp}]
	if b {
		return r
	} else if len(arr) == 0 {
		if len(num) == 0 {
			opt[Info{len(arr), len(num), tmp}] = 1
			return 1
		} else {
			opt[Info{len(arr), len(num), tmp}] = 0
			return 0
		}
	} else if len(num) == 0 {
		for _, u := range arr {
			if u == "#" {
				opt[Info{len(arr), len(num), tmp}] = 0
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
						opt[Info{len(arr), len(num), tmp}] = 1
						return 1
					} else {
						opt[Info{len(arr), len(num), tmp}] = 0
						return 0
					}
				} else {
					res := f3(arr[1:], num, tmp+1, false, opt)
					opt[Info{len(arr), len(num), tmp}] = res
					return res
				}
			}
		} else if arr[0] == "." {
			if first {
				res := f3(arr[1:], num, 0, first, opt)
				opt[Info{len(arr), len(num), tmp}] = res
				return res
			} else {
				if tmp < num[0] {
					opt[Info{len(arr), len(num), tmp}] = 0
					return 0
				} else {
					res := f3(arr[1:], num[1:], 0, true, opt)
					opt[Info{len(arr), len(num), tmp}] = res
					return res
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
			opt[Info{len(arr), len(num), tmp}] = r1 + r2
			return r1 + r2
		}
	}
	return tmp
}
