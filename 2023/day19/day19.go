package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	//fmt.Println("///")
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

	m := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			parseInput(line, m)
		}
	}

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		tmp := "in"
		for tmp != "R" && tmp != "A" {
			tmp = evaluateConditions(line, m[tmp])
		}
		if tmp == "A" {
			res += addAll(line)
		} else {

		}
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

	M := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			parseInput(line, M)
		}
	}

	lrange := make([][4]Range, 0)
	solve(M, "in", Range{1, 4000}, Range{1, 4000}, Range{1, 4000}, Range{1, 4000}, &lrange)

	res := 0
	x := 0
	m := 0
	a := 0
	s := 0
	for _, u := range lrange {
		x += (u[0].B - u[0].A + 1)
		m += (u[1].B - u[1].A + 1)
		a += (u[2].B - u[2].A + 1)
		s += (u[3].B - u[3].A + 1)
		res += (u[0].B - u[0].A + 1) * (u[1].B - u[1].A + 1) * (u[2].B - u[2].A + 1) * (u[3].B - u[3].A + 1)
	}
	//res = x * m * a * s

	//res := cardinality(lrange)

	return res
}

func parseInput(input string, m map[string][]string) {

	// Utilisation d'une expression régulière pour extraire les paires clé-valeur
	re := regexp.MustCompile(`([a-zA-Z]+){([^{}]+)}`)
	matches := re.FindAllStringSubmatch(input, -1)

	// Parcours des correspondances et construction de la map
	for _, match := range matches {
		key := match[1]
		values := strings.Split(match[2], ",")
		for i, value := range values {
			values[i] = strings.TrimSpace(value)
		}
		m[key] = values
	}

}

func evaluateConditions(data string, conditions []string) string {
	// Utilisation d'une expression régulière pour extraire les valeurs associées aux clés
	re := regexp.MustCompile(`([a-zA-Z])=(\d+)`)
	matches := re.FindAllStringSubmatch(data, -1)

	// Construction d'une map pour stocker les valeurs associées aux clés
	dataMap := make(map[string]int)
	for _, match := range matches {
		key := match[1]
		value, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Erreur de conversion de la valeur en entier.")
			return ""
		}
		dataMap[key] = value
	}

	// Évaluation des conditions
	for _, condition := range conditions {
		parts := strings.Split(condition, ":")
		if len(parts) != 2 {
			//fmt.Println("Condition mal formée:", condition)
			continue
		}

		// Extraction de la clé, de l'opérateur et de la valeur de la condition
		key := string(parts[0][0])
		operator := string(parts[0][1])          // Dernier caractère est l'opérateur
		value, err := strconv.Atoi(parts[0][2:]) // Tout sauf le dernier caractère est la valeur
		if err != nil {
			fmt.Println("Erreur de conversion de la valeur de la condition en entier.")
			continue
		}

		// Vérification des conditions
		switch operator {
		case "<":
			if dataMap[key] < value {
				return parts[1]
			}
		case ">":
			if dataMap[key] > value {
				return parts[1]
			}
			// Ajoutez d'autres opérateurs selon les besoins
		}
	}

	// Aucune condition n'est satisfaite, retourne la dernière valeur par défaut
	return conditions[len(conditions)-1]
}

func addAll(line string) int {
	re := regexp.MustCompile(`([a-zA-Z])=(\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)

	res := 0
	for i, _ := range matches {
		num, _ := strconv.Atoi(matches[i][2])
		res += num
	}
	return res
}

type Range struct {
	A int
	B int
}

func solve(M map[string][]string, k string, x, m, a, s Range, lrange *[][4]Range) {
	conditions := M[k]
	dataMap := make(map[string]Range)
	dataMap["x"] = x
	dataMap["m"] = m
	dataMap["a"] = a
	dataMap["s"] = s
	// Évaluation des conditions
	res := 0
	for _, condition := range conditions {
		parts := strings.Split(condition, ":")
		if len(parts) != 2 {
			//fmt.Println("Condition mal formée:", condition)
			continue
		}

		// Extraction de la clé, de l'opérateur et de la valeur de la condition
		key := string(parts[0][0])
		operator := string(parts[0][1])          // Dernier caractère est l'opérateur
		value, err := strconv.Atoi(parts[0][2:]) // Tout sauf le dernier caractère est la valeur
		if err != nil {
			fmt.Println("Erreur de conversion de la valeur de la condition en entier.")
			continue
		}

		// Vérification des conditions
		switch operator {
		case "<":
			if dataMap[key].A < value {
				if parts[1] == "A" {
					if key == "a" {
						r := [4]Range{x, m, {a.A, Min(a.B, value-1)}, s}
						*lrange = append(*lrange, r)
						a = Range{value, a.B}
					} else if key == "s" {
						r := [4]Range{x, m, a, {s.A, Min(value-1, s.B)}}
						*lrange = append(*lrange, r)
						s = Range{value, s.B}
					} else if key == "m" {
						r := [4]Range{x, {m.A, Min(value-1, m.B)}, a, s}
						*lrange = append(*lrange, r)
						m = Range{value, m.B}
					} else if key == "x" {
						r := [4]Range{{x.A, Min(value-1, x.B)}, m, a, s}
						*lrange = append(*lrange, r)
						x = Range{value, x.B}
					}
				} else if parts[1] == "R" {
					if key == "a" {
						a = Range{value, a.B}
					} else if key == "s" {
						s = Range{value, s.B}
					} else if key == "m" {
						m = Range{value, m.B}
					} else if key == "x" {
						x = Range{value, x.B}
					}
				} else {
					if key == "m" {
						solve(M, parts[1], x, Range{m.A, Min(value-1, m.B)}, a, s, lrange)
						m = Range{value, m.B}
					} else if key == "s" {
						solve(M, parts[1], x, m, a, Range{s.A, Min(value-1, s.B)}, lrange)
						s = Range{value, s.B}
					} else if key == "a" {
						solve(M, parts[1], x, m, Range{a.A, Min(value-1, a.B)}, s, lrange)
						a = Range{value, a.B}
					} else if key == "x" {
						solve(M, parts[1], Range{x.A, Min(value-1, x.B)}, m, a, s, lrange)
						x = Range{value, x.B}
					}
				}
			}
		case ">":
			if dataMap[key].B > value {
				if parts[1] == "A" {
					if key == "a" {
						r := [4]Range{x, m, {Max(value+1, a.A), a.B}, s}
						*lrange = append(*lrange, r)
						a = Range{a.A, value}
					} else if key == "s" {
						r := [4]Range{x, m, a, {Max(value+1, s.A), s.B}}
						*lrange = append(*lrange, r)
						s = Range{s.A, value}
					} else if key == "m" {
						r := [4]Range{x, {Max(value+1, m.A), m.B}, a, s}
						*lrange = append(*lrange, r)
						m = Range{m.A, value}
					} else if key == "x" {
						r := [4]Range{{Max(value+1, x.A), x.B}, m, a, s}
						*lrange = append(*lrange, r)
						x = Range{x.A, value}
					}
				} else if parts[1] == "R" {
					if key == "a" {
						a = Range{a.A, value}
					} else if key == "s" {
						s = Range{s.A, value}
					} else if key == "m" {
						m = Range{m.A, value}
					} else if key == "x" {
						x = Range{x.A, value}
					}
				} else {
					if key == "m" {
						solve(M, parts[1], x, Range{Max(m.A, value+1), m.B}, a, s, lrange)
						m = Range{m.A, value}
					} else if key == "s" {
						solve(M, parts[1], x, m, a, Range{Max(s.A, value+1), s.B}, lrange)
						s = Range{s.A, value}
					} else if key == "a" {
						solve(M, parts[1], x, m, Range{Max(a.A, value+1), a.B}, s, lrange)
						a = Range{a.A, value}
					} else if key == "x" {
						solve(M, parts[1], Range{Max(x.A, value+1), x.B}, m, a, s, lrange)
						x = Range{x.A, value}
					}
				}
			}
			// Ajoutez d'autres opérateurs selon les besoins
		}
	}

	// Aucune condition n'est satisfaite, retourne la dernière valeur par défaut
	tmp := conditions[len(conditions)-1]
	if tmp == "A" {
		*lrange = append(*lrange, [4]Range{x, m, a, s})
	} else if tmp == "R" {
		res += 0
	} else {
		solve(M, M[k][len(M[k])-1], x, m, a, s, lrange)
	}
	return
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
