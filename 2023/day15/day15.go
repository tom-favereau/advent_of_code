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

	arr := strings.Split(line, ",")

	res := 0
	for _, u := range arr {
		l := strings.Split(u, "")
		tmp := 0
		for _, t := range l {
			asc := int(t[0])
			tmp += asc
			tmp *= 17
			tmp = tmp % 256

		}
		res += tmp
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
	line := scanner.Text()

	arr := strings.Split(line, ",")

	res := make([][]string, 256)

	for _, u := range arr {
		indexOp := 0
		for i := range u {
			if u[i] == '=' || u[i] == '-' {
				indexOp = i
				break
			}
		}
		h := hash(u[:indexOp])
		if u[indexOp] == '-' {
			for k, t := range res[h] {
				if len(t) > indexOp && u[:indexOp] == t[:indexOp] {
					res[h] = append(res[h][:k], res[h][k+1:]...)
				}
			}
		} else if u[indexOp] == '=' {
			if len(res[h]) == 0 {
				res[h] = append(res[h], u)
			} else {
				find := false
				for k, t := range res[h] {
					if len(t) > indexOp && (t[indexOp] == '-' || t[indexOp] == '=') && u[:indexOp] == t[:indexOp] {
						res[h][k] = u
						find = true
					}
				}
				if !find {
					res[h] = append(res[h], u)
				}
			}
		}
	}

	ans := 0
	for i, u := range res {
		for j, t := range u {
			num, _ := strconv.Atoi(string(t[len(t)-1]))
			ans += (i + 1) * (j + 1) * num
		}
	}

	return ans
}

func hash(s string) int {
	res := 0
	for _, u := range s {
		asc := int(u)
		res += asc
		res *= 17
		res = res % 256
	}
	return res
}
