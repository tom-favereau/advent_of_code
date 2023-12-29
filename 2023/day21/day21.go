package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// la figure forme un diamand, on calcule le nombre de case impair dans la grille
// on fait 26501365//131 = 202300
// l'aire du diamand est donné par 2*(203200**2)+2*203200
// je tente les polynome de lagrange
// c'est pas stupide puisque mon calcul avec les le diamand est quadraticn
// l'aire est en x**2, la frontière est en x, et les nord sud est ouest sont constant
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

	var stepMax int
	if string2 == "input.txt" {
		stepMax = 64
	} else {
		stepMax = 6
	}

	scanner := bufio.NewScanner(file)
	g := make([][]string, 0)
	i := 0
	ps := P{0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")
		tmp := make([]string, 0)
		for j, u := range arr {
			if u == "S" {
				ps = P{i, j}
			}
			tmp = append(tmp, u)
		}
		g = append(g, tmp)
		i++
	}
	/*
		vis := make(map[P]bool)
		solve(g, ps, 0, vis)
		res := 0
		//tmp := make(map[P]bool)
		for k, _ := range vis {
			if (k.i+k.j)%2 == 0 {
				res++
			}
		}

	*/
	res := solve(g, ps, stepMax)
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
	i := 0
	ps := P{0, 0}
	pair := 0
	impair := 0
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")
		tmp := make([]string, 0)
		for j, u := range arr {
			if u == "S" {
				ps = P{i, j}
			} else if u == "." {
				if (i+j)%2 == 0 {
					pair++
				} else {
					impair++
				}
			}
			tmp = append(tmp, u)
		}
		g = append(g, tmp)
		i++
	}

	//fmt.Println(pair, impair)
	res := solveP24real(g, ps)
	return res
}

type Visit struct {
	p     P
	index int
}

type P struct {
	i, j int
}

func solveP24real(g [][]string, ps P) int {
	sizeDiamond := 26501365 / len(g) //203200
	end := 26501365 % len(g)         // 65

	firstPoint := solve(g, ps, end)
	secondPoint := solve(g, ps, 2*len(g)+end)
	thirdPoint := solve(g, ps, 4*len(g)+end)

	x := []float64{0, 1, 2}
	y := []float64{float64(firstPoint), float64(secondPoint), float64(thirdPoint)}
	//fmt.Println(lagrangeInterpolation(x, y, float64(x0)))
	res := lagrangeInterpolation(x, y, float64(sizeDiamond+end))

	return int(res)
}

func lagrangeInterpolation(x []float64, y []float64, xVal float64) float64 {
	result := 0.0
	n := len(x)

	for i := 0; i < n; i++ {
		term := y[i]
		for j := 0; j < n; j++ {
			if j != i {
				term *= (xVal - x[j]) / (x[i] - x[j])
			}
		}
		result += term
	}

	return result
}

func f(g [][]string, p P, impair int) int {
	nb := 0
	i := 1
	pr := -1
	for nb != pr {
		if i%2 == 1 {
			pr = nb
			//fmt.Println(nb)
		}
		vis := make(map[Visit]bool)
		solve2(g, p, 0, vis, i)
		nb = 0
		tmp := make(map[P]bool)
		for k, _ := range vis {
			if k.index%2 == 1 {
				if !tmp[k.p] {
					//fmt.Println(k.p.i, " ", k.p.j)
					nb++
					tmp[k.p] = true
				}
			}
		}
		i++
	}
	fmt.Println("nb case impair dans le cycle", nb)
	return i - 1
}

/*
func solve(g [][]string, p P, index int, vis map[P]bool) {
	i := p.i
	j := p.j
	if i >= 0 && i < len(g) && j >= 0 && j < len(g[0]) && index <= 6 {

		cur := g[i][j]
		if (cur == "S" || cur == ".") && !vis[p] {
			vis[p] = true
			solve(g, P{i + 1, j}, index+1, vis)
			solve(g, P{i, j + 1}, index+1, vis)
			solve(g, P{i - 1, j}, index+1, vis)
			solve(g, P{i, j - 1}, index+1, vis)
		}
	}
}

*/

func solve(g [][]string, start P, stepMax int) int {
	queue := make([]Visit, 0, stepMax*stepMax)
	queue = append(queue, Visit{start, 0})
	visited := make(map[P]bool)
	for len(queue) > 0 {
		cur := queue[0]
		i := cur.p.i % len(g)
		if i < 0 {
			i += len(g)
		}
		j := cur.p.j % len(g)
		if j < 0 {
			j += len(g)
		}
		symb := g[i][j]
		queue = queue[1:]

		if (symb == "S" || symb == ".") && !visited[cur.p] {
			visited[cur.p] = true
			if cur.index < stepMax {
				queue = append(queue, Visit{P{cur.p.i + 1, cur.p.j}, cur.index + 1})
				queue = append(queue, Visit{P{cur.p.i, cur.p.j + 1}, cur.index + 1})
				queue = append(queue, Visit{P{cur.p.i - 1, cur.p.j}, cur.index + 1})
				queue = append(queue, Visit{P{cur.p.i, cur.p.j - 1}, cur.index + 1})
			}
		}
	}
	res := 0
	for p, val := range visited {
		if val && (p.i+p.j)%2 == stepMax%2 {
			res++
		}
	}
	return res
}

func solve2(g [][]string, p P, index int, vis map[Visit]bool, indM int) {
	i := p.i % len(g)
	j := p.j % len(g)
	if i < 0 {
		i += len(g)
	}
	if j < 0 {
		j += len(g)
	}
	if index <= indM {
		//fmt.Println(i%len(g), j%len(g))
		cur := g[i][j]
		if (cur == "S" || cur == ".") && !vis[Visit{p, index}] {
			vis[Visit{p, index}] = true
			solve2(g, P{p.i + 1, p.j}, index+1, vis, indM)
			solve2(g, P{p.i, p.j + 1}, index+1, vis, indM)
			solve2(g, P{p.i - 1, p.j}, index+1, vis, indM)
			solve2(g, P{p.i, p.j - 1}, index+1, vis, indM)
		}
	}
}

func solve3(g [][]string, p P, index int, vis map[Visit]bool) {
	i := p.i
	j := p.j
	if i%131 >= 0 && i%131 < len(g) && j%131 >= 0 && j%131 < len(g[0]) && index <= 3*131 {

		cur := g[i%131][j%131]
		if (cur == "S" || cur == ".") && !vis[Visit{p, index}] {
			vis[Visit{p, index}] = true
			solve3(g, P{i + 1, j}, index+1, vis)
			solve3(g, P{i, j + 1}, index+1, vis)
			solve3(g, P{i - 1, j}, index+1, vis)
			solve3(g, P{i, j - 1}, index+1, vis)
		}
	}
}
