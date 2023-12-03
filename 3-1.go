package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coord struct {
    x int
    y int
}

type Number struct {
    row   int
    start int
    end   int
}

func main() {
	file, err := os.Open("3.in")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    symbols := map[Coord]bool{}
    numbers := map[Number]int{}
    i := 0
    for scanner.Scan() {
        line := scanner.Text()
        start := -1
        num := ""
        for j := range line {
            if '0' <= line[j] && line[j] <= '9' {
                // if start of new number, store its starting index
                if start == -1 {
                    start = j
                }
                num += string(line[j])
            } else {
                addNumber(i, &start, j - 1, &num, numbers)
                if isSymbol(line[j]) {
                    addAdjacentPoints(j, i, symbols)
                }
            }
        }
        // if last digit is number, save it to the map
        addNumber(i, &start, len(line) - 1, &num, numbers)
        i++
    }
    sum := 0
    // sum numbers that are adjacent to symbols
    for number, value := range numbers {
        for x := number.start; x <= number.end ; x++ {
            if symbols[Coord{x: x, y: number.row}] {
                sum += value
                break
            }
        }
    }
    fmt.Println(sum)
}

func isSymbol(s byte) bool {
    return s != '.' && !('0' <= s && s <= '9')
}

func addAdjacentPoints(x int, y int, symbolMap map[Coord]bool) {
    dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
    for _, i := range dy {
        for _, j := range dx {
            symbolMap[Coord{y: y + i, x: x + j}] = true
        }
    }
}

func addNumber(row int, start *int, end int, num *string, numberMap map[Number]int) {
    if *start != -1 {
        n, _ := strconv.Atoi(*num)
        numberMap[Number{row, *start, end}] = n
        *start = -1
        *num = ""
    }
}
