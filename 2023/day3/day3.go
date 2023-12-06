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
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	//digit := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	res := 0
	for i, u := range matrix {
		currentNumber := 0
		nextToSymbol := false
		for j, v := range u {
			number, e := strconv.Atoi(v)
			if e == nil {
				currentNumber = currentNumber*10 + number
				if i == 0 && j == 0 {
					if isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == 0 && j == len(matrix[0])-1 {
					if isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 && j == 0 {
					if isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 && j == len(matrix[0])-1 {
					if isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j-1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == 0 {
					if isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == 0 {
					if isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i+1][j]) || isSymbo(matrix[i-1][j]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 {
					if isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j-1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == len(matrix[0])-1 {
					if isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j-1]) || isSymbo(matrix[i-1][j-1]) || isSymbo(matrix[i+1][j]) || isSymbo(matrix[i-1][j]) {
						nextToSymbol = true
					}
				} else {
					if isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i][j-1]) || isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j-1]) {
						nextToSymbol = true
					}
				}
			} else {
				if nextToSymbol {
					res += currentNumber
				} else if currentNumber != 0 {
				}
				currentNumber = 0
				nextToSymbol = false
			}

		}
		if nextToSymbol {
			res += currentNumber
		}
		currentNumber = 0
		nextToSymbol = false
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

	matrix := make([][]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	res := 0
	for i, u := range matrix {
		for j, v := range u {
			if v == "*" {
				a1 := int(math.Max(float64(i-1), 0))
				a2 := int(math.Min(float64(i+1), float64(len(matrix)-1)))
				b1 := int(math.Max(float64(j-3), 0))
				b2 := int(math.Min(float64(j+3), float64(len(matrix[0])-1)))
				subMatrix := make([][]string, 0, 3)
				for n := 0; n <= a2-a1; n++ {
					tmp := make([]string, 0, 5)
					for m := 0; m <= b2-b1; m++ {
						if isSymbo(matrix[a1+n][b1+m]) {
							tmp = append(tmp, ".")
						} else {
							tmp = append(tmp, matrix[a1+n][b1+m])
						}
					}
					subMatrix = append(subMatrix, tmp)
				}
				subMatrix[1][3] = "*"
				res += aux_part2(subMatrix)
			}
		}
	}

	return res

}

func aux_part2(matrix [][]string) int {
	res := 1
	conteur := 0
	for i, u := range matrix {
		currentNumber := 0
		nextToSymbol := false
		for j, v := range u {
			number, e := strconv.Atoi(v)
			if e == nil {
				currentNumber = currentNumber*10 + number
				if i == 0 && j == 0 {
					if isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == 0 && j == len(matrix[0])-1 {
					if isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 && j == 0 {
					if isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 && j == len(matrix[0])-1 {
					if isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j-1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == 0 {
					if isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == 0 {
					if isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i+1][j]) || isSymbo(matrix[i-1][j]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 {
					if isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j-1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == len(matrix[0])-1 {
					if isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j-1]) || isSymbo(matrix[i-1][j-1]) || isSymbo(matrix[i+1][j]) || isSymbo(matrix[i-1][j]) {
						nextToSymbol = true
					}
				} else {
					if isSymbo(matrix[i+1][j]) || isSymbo(matrix[i+1][j+1]) || isSymbo(matrix[i+1][j-1]) || isSymbo(matrix[i][j+1]) || isSymbo(matrix[i][j-1]) || isSymbo(matrix[i-1][j+1]) || isSymbo(matrix[i-1][j]) || isSymbo(matrix[i-1][j-1]) {
						nextToSymbol = true
					}
				}
			} else {
				if nextToSymbol {
					res *= currentNumber
					conteur++
				} else if currentNumber != 0 {
				}
				currentNumber = 0
				nextToSymbol = false
			}

		}
		if nextToSymbol {
			res *= currentNumber
			conteur++
		}
		currentNumber = 0
		nextToSymbol = false
	}
	if conteur == 2 {
		return res
	} else {
		return 0
	}
}

func isSymbo(s string) bool {
	if s == "." || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" || s == "0" {
		return false
	} else {
		return true
	}
}
