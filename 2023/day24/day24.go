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
	//fmt.Println("///")
	//fmt.Println(part2("input.txt"))
	//fmt.Println(part2("input2.txt"))

}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()
	var m float64
	var M float64
	if string2 == "input.txt" {
		m = 200000000000000
		M = 400000000000000
	} else {
		m = 7
		M = 27
	}

	scanner := bufio.NewScanner(file)
	pos := make([]Vector, 0)
	vel := make([]Vector, 0)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " @ ")
		p := strings.Split(sep[0], ", ")
		v := strings.Split(sep[1], ", ")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		z, _ := strconv.Atoi(p[2])
		pos = append(pos, Vector{float64(x), float64(y), float64(z)})

		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		vz, _ := strconv.Atoi(v[2])
		vel = append(vel, Vector{float64(vx), float64(vy), float64(vz)})
	}

	res := 0
	for i, _ := range pos {
		for j, _ := range pos {
			if i != j {
				A := [][]float64{
					{-vel[i].x, vel[j].x},
					{-vel[i].y, vel[j].y},
				}
				Y := []float64{pos[i].x - pos[j].x, pos[i].y - pos[j].y}
				t, e := gaussElimination(A, Y)
				x1 := pos[i].x + t[0]*vel[i].x
				y1 := pos[i].y + t[0]*vel[i].y
				x2 := pos[j].x + t[1]*vel[j].x
				y2 := pos[j].y + t[1]*vel[j].y
				if e == nil && t[0] >= 0 && t[1] >= 0 && x1 <= M && x1 >= m && x2 >= m && x2 <= M && y1 <= M && y1 >= m && y2 >= m && y2 <= M {
					res++
				}
			}
		}
	}
	return res / 2
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()
	var m float64
	var M float64
	if string2 == "input.txt" {
		m = 200000000000000
		M = 400000000000000
	} else {
		m = 7
		M = 27
	}

	scanner := bufio.NewScanner(file)
	pos := make([]Vector, 0)
	vel := make([]Vector, 0)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " @ ")
		p := strings.Split(sep[0], ", ")
		v := strings.Split(sep[1], ", ")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		z, _ := strconv.Atoi(p[2])
		pos = append(pos, Vector{float64(x), float64(y), float64(z)})

		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		vz, _ := strconv.Atoi(v[2])
		vel = append(vel, Vector{float64(vx), float64(vy), float64(vz)})
	}

	return 0
}

type Vector struct {
	x, y, z float64
}

// gaussElimination résout le système d'équations linéaires AX = Y par pivot de Gauss
func gaussElimination(A [][]float64, Y []float64) ([]float64, error) {
	n := len(Y)

	// Vérifier si la matrice A est carrée et de la bonne taille
	if len(A) != n || len(A[0]) != n {
		return nil, fmt.Errorf("La matrice A doit être carrée de taille %dx%d", n, n)
	}

	// Concaténer la matrice A et le vecteur Y pour former une matrice augmentée
	for i := 0; i < n; i++ {
		A[i] = append(A[i], Y[i])
	}

	// Appliquer la méthode du pivot de Gauss
	for i := 0; i < n; i++ {
		// Étape 1: Choix du pivot
		pivotRow := i
		for j := i + 1; j < n; j++ {
			if abs(A[j][i]) > abs(A[pivotRow][i]) {
				pivotRow = j
			}
		}

		// Échanger les lignes
		A[i], A[pivotRow] = A[pivotRow], A[i]

		// Étape 2: Élimination
		for j := i + 1; j < n; j++ {
			factor := A[j][i] / A[i][i]
			for k := i; k <= n; k++ {
				A[j][k] -= factor * A[i][k]
			}
		}
	}

	// Résoudre le système d'équations résultant
	X := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		X[i] = A[i][n] / A[i][i]
		for j := i - 1; j >= 0; j-- {
			A[j][n] -= A[j][i] * X[i]
		}
	}

	return X, nil
}

// Fonction utilitaire pour obtenir la valeur absolue d'un nombre
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
