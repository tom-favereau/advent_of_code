package main

import (
	"bufio"
	"fmt"
	"os"
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

	g := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")

		tmp := make([]string, 0)
		for _, u := range arr {
			tmp = append(tmp, u)
		}
		g = append(g, tmp)
	}
	start := [2]int{0, 1}
	end := [2]int{len(g) - 1, len(g) - 2}
	res := findLongestPath(g, start, end)
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

	g := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")

		tmp := make([]string, 0)
		for _, u := range arr {
			tmp = append(tmp, u)
		}
		g = append(g, tmp)
	}
	start := [2]int{0, 1}
	end := [2]int{len(g) - 1, len(g) - 2}
	res := findLongestPath3(g, start, end)
	return res
}

func findLongestPath(maze [][]string, start [2]int, end [2]int) int {
	var dfs func([2]int, int, map[[2]int]bool)
	maxPathLength := 0
	dfs = func(current [2]int, pathLength int, visited map[[2]int]bool) {
		// Arrêt de la récursion si on atteint la sortie
		if current == end {
			if pathLength > maxPathLength {
				maxPathLength = pathLength
			}
			return
		}

		// Marquer la case actuelle comme visitée
		visited[current] = true

		// Explorer les cases voisines non visitées
		for _, neighbor := range neighbors(current, maze) {
			if !visited[neighbor] {
				dfs(neighbor, pathLength+1, visited)
			}
		}

		// Retirer la marque de la case actuelle pour permettre d'autres explorations
		delete(visited, current)
	}

	maxPathLength = 0
	visited := make(map[[2]int]bool)
	dfs(start, 0, visited)

	return maxPathLength
}

type Edge struct {
	dist int
	to   [2]int
}

func findLongestPath3(maze [][]string, start [2]int, end [2]int) int {
	var dfs func([2]int, [2]int, [2]int, int, map[[2]int]int)
	var auxP2 func([2]int, int, map[[2]int]bool)
	dist := make(map[[2]int][]Edge)
	dfs = func(current [2]int, from [2]int, pr [2]int, pathLength int, visited map[[2]int]int) {
		// Arrêt de la récursion si on atteint la sortie
		if current == end {
			dist[current] = append(dist[current], Edge{pathLength + 1, from})
			dist[from] = append(dist[from], Edge{pathLength + 1, current})
			return
		}

		neights := neighbors2(current, maze)
		// Marquer la case actuelle comme visitée
		visited[current] += 1
		if len(neights) > 2 {
			dist[current] = append(dist[current], Edge{pathLength + 1, from})
			dist[from] = append(dist[from], Edge{pathLength + 1, current})
			for _, neighbor := range neighbors2(current, maze) {
				if visited[neighbor] < len(neighbors2(neighbor, maze))-1 && neighbor != pr {
					dfs(neighbor, current, current, 0, visited)
				}
			}
		} else {
			// Explorer les cases voisines non visitées
			for _, neighbor := range neights {
				if (visited[neighbor] < len(neighbors2(neighbor, maze))-1 && neighbor != pr) || neighbor == end {
					dfs(neighbor, from, current, pathLength+1, visited)
				}
			}
		}
		//delete(visited, current)

	}

	visited := make(map[[2]int]int)
	dfs(start, start, start, 0, visited)
	//fmt.Println(dist)
	maxPathLength := 0
	auxP2 = func(current [2]int, pathLength int, visited map[[2]int]bool) {
		// Arrêt de la récursion si on atteint la sortie
		if current == end {
			if pathLength > maxPathLength {
				maxPathLength = pathLength
			}
			return
		}

		// Marquer la case actuelle comme visitée
		visited[current] = true

		// Explorer les cases voisines non visitées
		for _, neighbor := range dist[current] {
			if !visited[neighbor.to] {
				auxP2(neighbor.to, pathLength+neighbor.dist, visited)
			}
		}

		// Retirer la marque de la case actuelle pour permettre d'autres explorations
		delete(visited, current)
	}
	visited2 := make(map[[2]int]bool)
	auxP2(start, 0, visited2)

	return maxPathLength - 1
}

// Fonction utilitaire pour obtenir les voisins valides d'une case
func neighbors(pos [2]int, maze [][]string) [][2]int {
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var result [][2]int
	if maze[pos[0]][pos[1]] == ">" {
		newPos := [2]int{pos[0], pos[1] + 1}
		if isValid(newPos, maze) {
			result = append(result, newPos)
		}
	} else if maze[pos[0]][pos[1]] == "<" {
		newPos := [2]int{pos[0], pos[1] - 1}
		if isValid(newPos, maze) {
			result = append(result, newPos)
		}
	} else if maze[pos[0]][pos[1]] == "v" {
		newPos := [2]int{pos[0] + 1, pos[1]}
		if isValid(newPos, maze) {
			result = append(result, newPos)
		}
	} else if maze[pos[0]][pos[1]] == "^" {
		newPos := [2]int{pos[0] - 1, pos[1]}
		if isValid(newPos, maze) {
			result = append(result, newPos)
		}
	} else {
		for _, dir := range directions {
			newPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if isValid(newPos, maze) {
				result = append(result, newPos)
			}
		}
	}

	return result
}

func neighbors2(pos [2]int, maze [][]string) [][2]int {
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var result [][2]int

	for _, dir := range directions {
		newPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if isValid(newPos, maze) {
			result = append(result, newPos)
		}
	}

	return result
}

// Fonction utilitaire pour vérifier si une position est valide dans le labyrinthe
func isValid(pos [2]int, maze [][]string) bool {
	rows, cols := len(maze), len(maze[0])
	return pos[0] >= 0 && pos[0] < rows && pos[1] >= 0 && pos[1] < cols &&
		maze[pos[0]][pos[1]] != "#"
}
