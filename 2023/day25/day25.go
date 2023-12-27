package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

	scanner := bufio.NewScanner(file)

	gr := make(map[string][]string)
	vert := 0
	ind := 0
	indexes := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, ": ")
		neights := strings.Split(sep[1], " ")
		gr[sep[0]] = neights
		vert += 1
		_, b := indexes[sep[0]]
		if !b {
			indexes[sep[0]] = ind
			ind++
		}
		for _, u := range neights {
			_, b2 := indexes[u]
			if !b2 {
				indexes[u] = ind
				ind++
			}
		}

	}

	graph := newGraph(ind)

	for k, lu := range gr {
		for _, u := range lu {
			addEdge(graph, indexes[k], indexes[u])
		}
	}

	res := solve(graph)
	return res
}

func solve(graph *Graph) int {

	freq := make(map[[2]int]int)
	for startNode := 5; startNode < graph.Vertices; startNode++ {
		prev := make([]int, graph.Vertices)
		visited := make([]bool, graph.Vertices)
		fifo := make([]int, 0, graph.Vertices)
		fifo = append(fifo, startNode)
		visited[startNode] = true

		for len(fifo) > 0 {
			node := fifo[0]
			fifo = fifo[1:]

			for i := 0; i < graph.Vertices; i++ {
				if graph.Edges[node][i] == 1 && !visited[i] {
					visited[i] = true
					prev[i] = node
					fifo = append(fifo, i)
				}
			}
		}

		for i := 0; i < graph.Vertices; i++ {
			node := i
			for node != startNode {
				tmp := prev[node]
				freq[[2]int{Min(tmp, node), Max(node, tmp)}]++
				node = tmp
			}
		}
		//fmt.Println(freq[[2]int{0, 1}])

	}
	//fmt.Println(freq)
	m1 := 0
	m2 := 0
	m3 := 0
	var e1 [2]int
	var e2 [2]int
	var e3 [2]int
	for i := 0; i < len(freq); i++ {
		for j := 0; j < len(freq); j++ {
			v := freq[[2]int{i, j}]
			if v > m1 {
				m3 = m2
				e3 = e2
				m2 = m1
				e2 = e1
				m1 = v
				e1 = [2]int{i, j}
			} else if v > m2 {
				m3 = m2
				e3 = e2
				m2 = v
				e2 = [2]int{i, j}
			} else if v > m3 {
				m3 = v
				e3 = [2]int{i, j}
			}
		}
	}
	/*
		fmt.Println(m1, e1)
		fmt.Println(m2, e2)
		fmt.Println(m3, e3)
	*/
	removeEdge(graph, e1[0], e1[1])
	removeEdge(graph, e2[0], e2[1])
	removeEdge(graph, e3[0], e3[1])

	c := connectedComponentsSizes(graph)
	//fmt.Println(c)
	return c[0] * c[1]
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Graph struct {
	Vertices int
	Edges    [][]int
	val      []int
}

func newGraph(vertices int) *Graph {
	graph := &Graph{
		Vertices: vertices,
		Edges:    make([][]int, vertices),
		val:      make([]int, vertices),
	}

	for i := range graph.Edges {
		graph.Edges[i] = make([]int, vertices)
	}
	for i := range graph.val {
		graph.val[i] = 1
	}

	return graph
}

func addEdge(graph *Graph, src, dest int) {
	graph.Edges[src][dest] = 1
	graph.Edges[dest][src] = 1
}

func removeEdge(graph *Graph, src, dest int) {
	graph.Edges[src][dest] = 0
	graph.Edges[dest][src] = 0
}

func dfs(graph *Graph, node int, visited []bool, size *int) {
	visited[node] = true
	*size++

	for i := 0; i < graph.Vertices; i++ {
		if graph.Edges[node][i] == 1 && !visited[i] {
			//freq[[2]int{node, i}]++
			dfs(graph, i, visited, size)
		}
	}
}

func connectedComponentsSizes(graph *Graph) []int {
	visited := make([]bool, graph.Vertices)
	sizes := []int{}

	for i := 0; i < graph.Vertices; i++ {
		if !visited[i] {
			size := 0
			dfs(graph, i, visited, &size)
			sizes = append(sizes, size)
		}
	}

	return sizes
}

func minCut(graph *Graph) [2]int {
	rand.Seed(time.Now().UnixNano())

	// Copier le graphe pour éviter de le modifier directement
	tempGraph := &Graph{
		Vertices: graph.Vertices,
		Edges:    make([][]int, graph.Vertices),
		val:      make([]int, graph.Vertices),
	}
	copy(tempGraph.Edges, graph.Edges)
	for i := 0; i < tempGraph.Vertices; i++ {
		tempGraph.val[i] = 1
	}

	// Tableau pour suivre les arêtes contractées
	contractedEdges := make([][]int, 0)
	ind := 0
	for tempGraph.Vertices > 2 {
		// Choisir une arête aléatoire
		src := rand.Intn(tempGraph.Vertices)
		dest := rand.Intn(tempGraph.Vertices)

		notnull := make([][2]int, 0)
		for i := 0; i < tempGraph.Vertices; i++ {
			for j := 0; j < tempGraph.Vertices; j++ {
				if tempGraph.Edges[i][j] != 0 {
					notnull = append(notnull, [2]int{i, j})
				}
			}
		}
		//randomElem := notnull[rand.Intn(len(notnull))]
		//src := randomElem[0]
		//dest := randomElem[1]

		ind++
		if ind > 10000000 {
			//fmt.Println(tempGraph.Vertices)
			//fmt.Println(tempGraph.val)
			break
		}

		// Ignorer les boucles
		if src == dest {
			continue
		}

		// Ignorer les arêtes déjà contractées
		//if tempGraph.Edges[src][dest] == 0 {
		//	continue
		//}

		// Contracter l'arête
		for i := 0; i < tempGraph.Vertices; i++ {
			if tempGraph.Edges[i][dest] == 1 {
				tempGraph.Edges[i][src] += tempGraph.Edges[i][dest]
				tempGraph.Edges[i][dest] = 0
			}
		}
		tempGraph.val[src] += tempGraph.val[dest]
		tempGraph.val = append(tempGraph.val[:dest], tempGraph.val[dest+1:]...)

		// Supprimer la colonne et la ligne du nœud contracté
		tempGraph.Edges = append(tempGraph.Edges[:dest], tempGraph.Edges[dest+1:]...)
		for i := range tempGraph.Edges {
			tempGraph.Edges[i] = append(tempGraph.Edges[i][:dest], tempGraph.Edges[i][dest+1:]...)
		}

		// Ajouter l'arête contractée à la liste
		contractedEdges = append(contractedEdges, []int{src, dest})
		tempGraph.Vertices--
	}
	//fmt.Println(tempGraph.val, tempGraph.val[0]*tempGraph.val[1])

	return [2]int{tempGraph.val[0], tempGraph.val[1]}
}
