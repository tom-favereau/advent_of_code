package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
j'ai voulue faire le dernier jour en caml, je code doit être dans AOC_ocaml/bin/day25.ml
mon ford fulkerson ne marche pas bien

chat gpt m'a conseillé networkx et ça marche bien.
j'ai fait en go l'algorithme de wikipedia qui est heuristique, ça marche de temps en temps si on laisse tourné longtemps
donc pas concluant

je suis pas satisfait du tout de mes deux dernier jour, j'ai l'impression d'avoir juste utilisé
des module python et que le plus gros problème c'etait de lire la doc (et dialoguer avec chatgpt)
*/

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
	/*
		res := 0
		//fmt.Println(dfs(graph, ))
		//visited := make(map[[2]int]bool)
		for i := 0; i < ind; i++ {
			fmt.Println("ok1")
			for j := i + 1; j < ind; j++ {
				for k := 0; k < ind; k++ {
					fmt.Println("ok2")
					for l := k + 1; l < ind; l++ {
						if (k != i) || (l != j) {
							for m := 0; m < ind; m++ {
								fmt.Println("ok3")
								for n := m + 1; n < ind; n++ {
									if ((m != i) || (n != j)) && ((m != k) || (n != l)) {
										e1 := graph.Edges[i][j]
										e2 := graph.Edges[k][l]
										e3 := graph.Edges[m][n]
										removeEdge(graph, i, j)
										removeEdge(graph, k, l)
										removeEdge(graph, m, n)
										sizes := connectedComponentsSizes(graph)
										//fmt.Println(sizes)
										if i == indexes["hfx"] && j == indexes["pzl"] && k == indexes["bvb"] && l == indexes["cmg"] && m == indexes["nvd"] && n == indexes["jqt"] {
											fmt.Println(sizes)
											removeEdge(graph, indexes["hfx"], indexes["pzl"])
											removeEdge(graph, indexes["bvb"], indexes["cmg"])
											removeEdge(graph, indexes["nvd"], indexes["jqt"])
											size := connectedComponentsSizes(graph)
											fmt.Println(size)
										}
										if len(sizes) == 2 {
											res = sizes[0] * sizes[1]
											return res
										}
										graph.Edges[i][j] = e1
										graph.Edges[j][i] = e1
										graph.Edges[k][l] = e2
										graph.Edges[l][k] = e2
										graph.Edges[m][n] = e3
										graph.Edges[n][m] = e3
									}
								}
							}
						}
					}
				}
			}
		}
	*/

	//removeEdge(graph, indexes["hfx"], indexes["pzl"])
	//removeEdge(graph, indexes["bvb"], indexes["cmg"])
	//removeEdge(graph, indexes["nvd"], indexes["jqt"])
	//sizes := connectedComponentsSizes(graph)
	//fmt.Println(sizes)
	minCutComponents := minCut(graph)
	fmt.Println(minCutComponents)
	return minCutComponents[0] * minCutComponents[1]
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
	}

	for i := range graph.Edges {
		graph.Edges[i] = make([]int, vertices)
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
			fmt.Println(tempGraph.Vertices)
			fmt.Println(tempGraph.val)
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
	fmt.Println(tempGraph.val, tempGraph.val[0]*tempGraph.val[1])

	return [2]int{tempGraph.val[0], tempGraph.val[1]}
}
