package main

import (
	utils "advent_of_code/utils/grid"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(part1("input2.txt"))
	fmt.Println(part2opt("input.txt"))
}

func part1(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0

	grid := utils.NewGrid[string]()
	//sym := make([]int, 0)
	i := 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			end := false
			for k := 0; k < j-1; k++ {
				isSym := true
				for l := 0; l < j; l++ {
					for s := 0; s < i; s++ {
						v1, b1 := grid[utils.Pos{k - l, s}]
						v2, b2 := grid[utils.Pos{k + l + 1, s}]
						if b1 && b2 {
							if v1 != v2 {
								isSym = false
								break
							}
						}
					}
					if !isSym {
						break
					}
				}
				if isSym {
					res += k + 1
					end = true
					break
				}
			}
			if !end {
				for k := 0; k < i-1; k++ {
					isSym := true
					for l := 0; l < i; l++ {
						for s := 0; s < j; s++ {
							v1, b1 := grid[utils.Pos{s, k - l}]
							v2, b2 := grid[utils.Pos{s, k + l + 1}]
							if b1 && b2 {
								if v1 != v2 {
									isSym = false
									break
								}
							}
						}
						if !isSym {
							break
						}
					}
					if isSym {
						res += (k + 1) * 100
						break
					}
				}
			}
			grid = utils.NewGrid[string]()
			i = 0
			j = 0
		} else {
			arr := strings.Split(line, "")
			j = 0
			for _, u := range arr {
				grid.Add(u, utils.Pos{j, i})
				j++
			}
			i++
		}
	}
	end := false
	for k := 0; k < j-1; k++ {
		isSym := true
		for l := 0; l < j; l++ {
			for s := 0; s < i; s++ {
				v1, b1 := grid[utils.Pos{k - l, s}]
				v2, b2 := grid[utils.Pos{k + l + 1, s}]
				if b1 && b2 {
					if v1 != v2 {
						isSym = false
						break
					}
				}
			}
			if !isSym {
				break
			}
		}
		if isSym {
			res += k + 1
			end = true
			break
		}
	}
	if !end {
		for k := 0; k < i-1; k++ {
			isSym := true
			for l := 0; l < i; l++ {
				for s := 0; s < j; s++ {
					v1, b1 := grid[utils.Pos{s, k - l}]
					v2, b2 := grid[utils.Pos{s, k + l + 1}]
					if b1 && b2 {
						if v1 != v2 {
							isSym = false
							break
						}
					}
				}
				if !isSym {
					break
				}
			}
			if isSym {
				res += (k + 1) * 100
				break
			}
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

	res := 0

	grid := utils.NewGrid[string]()
	//sym := make([]int, 0)
	i := 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			tmp := 0
			for k := 0; k < j-1; k++ {
				isSym := true
				for l := 0; l < j; l++ {
					for s := 0; s < i; s++ {
						v1, b1 := grid[utils.Pos{k - l, s}]
						v2, b2 := grid[utils.Pos{k + l + 1, s}]
						if b1 && b2 {
							if v1 != v2 {
								isSym = false
								break
							}
						}
					}
					if !isSym {
						break
					}
				}
				if isSym {
					tmp = k + 1
					break
				}
			}

			for k := 0; k < i-1; k++ {
				isSym := true
				for l := 0; l < i; l++ {
					for s := 0; s < j; s++ {
						v1, b1 := grid[utils.Pos{s, k - l}]
						v2, b2 := grid[utils.Pos{s, k + l + 1}]
						if b1 && b2 {
							if v1 != v2 {
								isSym = false
							}
						}
					}
				}
				if isSym {
					tmp = (k + 1) * 100
					break
				}
			}
			for a := 0; a < i; a++ {
				for b := 0; b < j; b++ {
					if grid[utils.Pos{b, a}] == "#" {
						grid[utils.Pos{b, a}] = "."
					} else {
						grid[utils.Pos{b, a}] = "#"
					}
					for k := 0; k < j-1; k++ {
						isSym := true
						for l := 0; l < j; l++ {
							for s := 0; s < i; s++ {
								v1, b1 := grid[utils.Pos{k - l, s}]
								v2, b2 := grid[utils.Pos{k + l + 1, s}]
								if b1 && b2 {
									if v1 != v2 {
										isSym = false
									}
								}
							}
						}
						if isSym && tmp != k+1 {
							res += k + 1
							break
						}
					}

					for k := 0; k < i-1; k++ {
						isSym := true
						for l := 0; l < i; l++ {
							for s := 0; s < j; s++ {
								v1, b1 := grid[utils.Pos{s, k - l}]
								v2, b2 := grid[utils.Pos{s, k + l + 1}]
								if b1 && b2 {
									if v1 != v2 {
										isSym = false
									}
								}
							}
						}
						if isSym && tmp != (k+1)*100 {
							res += (k + 1) * 100
							break
						}
					}
					if grid[utils.Pos{b, a}] == "#" {
						grid[utils.Pos{b, a}] = "."
					} else {
						grid[utils.Pos{b, a}] = "#"
					}
				}

			}
			grid = utils.NewGrid[string]()
			i = 0
			j = 0
		} else {
			arr := strings.Split(line, "")
			j = 0
			for _, u := range arr {
				grid.Add(u, utils.Pos{j, i})
				j++
			}
			i++
		}
	}

	tmp := 0
	for k := 0; k < j-1; k++ {
		isSym := true
		for l := 0; l < j; l++ {
			for s := 0; s < i; s++ {
				v1, b1 := grid[utils.Pos{k - l, s}]
				v2, b2 := grid[utils.Pos{k + l + 1, s}]
				if b1 && b2 {
					if v1 != v2 {
						isSym = false
					}
				}
			}
		}
		if isSym {
			tmp = k + 1
		}
	}

	for k := 0; k < i-1; k++ {
		isSym := true
		for l := 0; l < i; l++ {
			for s := 0; s < j; s++ {
				v1, b1 := grid[utils.Pos{s, k - l}]
				v2, b2 := grid[utils.Pos{s, k + l + 1}]
				if b1 && b2 {
					if v1 != v2 {
						isSym = false
					}
				}
			}
		}
		if isSym {
			tmp = (k + 1) * 100
		}
	}
	for a := 0; a < i; a++ {
		for b := 0; b < j; b++ {
			if grid[utils.Pos{b, a}] == "#" {
				grid[utils.Pos{b, a}] = "."
			} else {
				grid[utils.Pos{b, a}] = "#"
			}
			for k := 0; k < j-1; k++ {
				isSym := true
				for l := 0; l < j; l++ {
					for s := 0; s < i; s++ {
						v1, b1 := grid[utils.Pos{k - l, s}]
						v2, b2 := grid[utils.Pos{k + l + 1, s}]
						if b1 && b2 {
							if v1 != v2 {
								isSym = false
							}
						}
					}
				}
				if isSym && tmp != k+1 {
					res += k + 1
					break
				}
			}

			for k := 0; k < i-1; k++ {
				isSym := true
				for l := 0; l < i; l++ {
					for s := 0; s < j; s++ {
						v1, b1 := grid[utils.Pos{s, k - l}]
						v2, b2 := grid[utils.Pos{s, k + l + 1}]
						if b1 && b2 {
							if v1 != v2 {
								isSym = false
							}
						}
					}
				}
				if isSym && tmp != (k+1)*100 {
					res += (k + 1) * 100
					break
				}
			}
			if grid[utils.Pos{b, a}] == "#" {
				grid[utils.Pos{b, a}] = "."
			} else {
				grid[utils.Pos{b, a}] = "#"
			}
		}

	}

	return res
}

func part2opt(string2 string) int {
	file, err := os.Open(string2)
	if err != nil {
		fmt.Println("erreur")
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0

	grid := utils.NewGrid[string]()
	//sym := make([]int, 0)
	i := 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			end := false
			for k := 0; k < j-1; k++ {
				nb_diff := 0
				for l := 0; l < j; l++ {
					for s := 0; s < i; s++ {
						v1, b1 := grid[utils.Pos{k - l, s}]
						v2, b2 := grid[utils.Pos{k + l + 1, s}]
						if b1 && b2 {
							if v1 != v2 {
								nb_diff++
								if nb_diff > 1 {
									break
								}
							}
						}
					}
					if nb_diff > 1 {
						break
					}
				}
				if nb_diff == 1 {
					res += k + 1
					end = true
					break
				}
			}

			if !end {
				for k := 0; k < i-1; k++ {
					nb_diff := 0
					for l := 0; l < i; l++ {
						for s := 0; s < j; s++ {
							v1, b1 := grid[utils.Pos{s, k - l}]
							v2, b2 := grid[utils.Pos{s, k + l + 1}]
							if b1 && b2 {
								if v1 != v2 {
									nb_diff++
									if nb_diff > 1 {
										break
									}
								}
							}
						}
						if nb_diff > 1 {
							break
						}
					}
					if nb_diff == 1 {
						res += (k + 1) * 100
						break
					}
				}
			}
			grid = utils.NewGrid[string]()
			i = 0
			j = 0
		} else {
			arr := strings.Split(line, "")
			j = 0
			for _, u := range arr {
				grid.Add(u, utils.Pos{j, i})
				j++
			}
			i++
		}
	}

	end := false
	for k := 0; k < j-1; k++ {
		nb_diff := 0
		for l := 0; l < j; l++ {
			for s := 0; s < i; s++ {
				v1, b1 := grid[utils.Pos{k - l, s}]
				v2, b2 := grid[utils.Pos{k + l + 1, s}]
				if b1 && b2 {
					if v1 != v2 {
						nb_diff++
						if nb_diff > 1 {
							break
						}
					}
				}
			}
			if nb_diff > 1 {
				break
			}
		}
		if nb_diff == 1 {
			res += k + 1
			end = true
			break
		}
	}
	if !end {
		for k := 0; k < i-1; k++ {
			nb_diff := 0
			for l := 0; l < i; l++ {
				for s := 0; s < j; s++ {
					v1, b1 := grid[utils.Pos{s, k - l}]
					v2, b2 := grid[utils.Pos{s, k + l + 1}]
					if b1 && b2 {
						if v1 != v2 {
							nb_diff++
							if nb_diff > 1 {
								break
							}
						}
					}
				}
				if nb_diff > 1 {
					break
				}
			}
			if nb_diff == 1 {
				res += (k + 1) * 100
				break
			}
		}
	}

	return res
}
