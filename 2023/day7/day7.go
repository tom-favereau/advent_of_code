package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
	//fmt.Println(part2("input2.txt"))
}

type Hand struct {
	h     map[string]int
	card1 []int
	point int
}

type Order []Hand

func (o Order) Len() int      { return len(o) }
func (o Order) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o Order) Less(i, j int) bool {
	// Comparer par taille de h décroissante
	if len(o[i].h) != len(o[j].h) {
		return len(o[j].h) > len(o[i].h)
	} else {
		maxi := -1
		maxj := -1
		for key := range o[i].h {
			if o[i].h[key] > maxi {
				maxi = o[i].h[key]
			}
		}
		for key := range o[j].h {
			if o[j].h[key] > maxj {
				maxj = o[j].h[key]
			}
		}

		if maxi != maxj {
			return maxi > maxj
		} else {
			for k, _ := range o[i].card1 {
				if o[i].card1[k] == o[j].card1[k] {
					continue
				} else {
					return o[i].card1[k] > o[j].card1[k]
				}

			}
		}

	}
	return true

	// En cas d'égalité (ou à 2 éléments si les deux ont 4 éléments ou 3 et 2),
	// comparer par card1 décroissante

}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	order := make(Order, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")
		cards := strings.Split(arr[0], "")
		num, _ := strconv.Atoi(arr[1])

		card1 := make([]int, 0, 10)

		for i, _ := range cards {
			first, e := strconv.Atoi(cards[i])
			if e != nil {
				if cards[i] == "T" {
					first = 10
				} else if cards[i] == "J" {
					first = 11
				} else if cards[i] == "Q" {
					first = 12
				} else if cards[i] == "K" {
					first = 13
				} else if cards[i] == "A" {
					first = 14
				}
			}
			card1 = append(card1, first)
		}

		tmp := Hand{make(map[string]int), card1, num}

		for _, u := range cards {
			tmp.h[u]++
		}
		order = append(order, tmp)
	}

	sort.Sort(order)
	res := 0
	for i, u := range order {
		res += (len(order) - i) * u.point
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

	order := make(Order, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")
		cards := strings.Split(arr[0], "")
		num, _ := strconv.Atoi(arr[1])

		card1 := make([]int, 0, 10)

		for i, _ := range cards {
			first, e := strconv.Atoi(cards[i])
			if e != nil {
				if cards[i] == "T" {
					first = 10
				} else if cards[i] == "J" {
					first = 0
				} else if cards[i] == "Q" {
					first = 12
				} else if cards[i] == "K" {
					first = 13
				} else if cards[i] == "A" {
					first = 14
				}
			}
			card1 = append(card1, first)
		}

		tmp := Hand{make(map[string]int), card1, num}
		nbJ := 0
		for _, u := range cards {
			if u == "J" {
				nbJ++

			} else {
				tmp.h[u]++
			}

		}
		maxval := -1
		maxkey := ""
		for key := range tmp.h {
			if tmp.h[key] > maxval {
				maxval = tmp.h[key]
				maxkey = key
			}
		}
		tmp.h[maxkey] += nbJ
		if len(tmp.h) == 0 {
			tmp.h["J"] = 5
		}

		order = append(order, tmp)
	}

	sort.Sort(order)
	res := 0
	for i, u := range order {
		res += (len(order) - i) * u.point
	}
	return res
}
