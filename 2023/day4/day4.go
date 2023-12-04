package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	//part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("erreur")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		separate := strings.Split(line[7:], "|")
		wining := strings.Split(separate[0], " ")
		mine := strings.Split(separate[1], " ")
		tmp := 0
		for _, u := range mine {
			for _, v := range wining {
				if u == v && u != "" {
					if tmp == 0 {
						tmp = 1
					} else {
						tmp *= 2
					}
				}
			}
		}
		res += tmp

	}
	fmt.Println(res)

}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("erreur")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	copyCard := make([]int, 300, 300)
	for i, _ := range copyCard {
		copyCard[i] = 1
	}
	cardNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		separate := strings.Split(line[7:], "|")
		wining := strings.Split(separate[0], " ")
		mine := strings.Split(separate[1], " ")
		tmp := 0
		count := 0
		for _, u := range mine {
			for _, v := range wining {
				if u == v && u != "" {
					if tmp == 0 {
						tmp = 1
					} else {
						tmp *= 2
					}
					count++
				}
			}
		}
		res += copyCard[cardNumber]
		for i := cardNumber + 1; i <= count+cardNumber; i++ {
			copyCard[i] += copyCard[cardNumber]
		}
		cardNumber++
	}
	fmt.Println(res)

}
