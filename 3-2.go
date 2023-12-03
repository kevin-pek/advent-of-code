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
    gearMap := map[Coord][]int{} // map coordinates to index of gears
    gears := []int{} // stores gears and their gear ratio
    numbers := map[Number]int{}
    grid := []string{}
    i := 0
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, line)
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
            }
        }
        // if last digit is number, save it to the map
        addNumber(i, &start, len(line) - 1, &num, numbers)
        i++
    }

    // find gears and update gearMap for adjacent points
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if isGear(r, c, grid) {
                fmt.Println(r, c)
                addNewGear(c, r, gearMap, &gears)
            }
        }
    }

    for numRange, number := range numbers {
        for x := numRange.start; x <= numRange.end ; x++ {
            coord := Coord{x: x, y: numRange.row}
            if gearIndices, ok := gearMap[coord]; ok {
                for _, i := range gearIndices {
                    gears[i] *= number
                }
                break
            }
        }
    }
    
    sum := 0
    for _, ratio := range gears {
        sum += ratio
    }
    fmt.Println(sum)
}

func addNumber(row int, start *int, end int, num *string, numberMap map[Number]int) {
    if *start != -1 {
        n, _ := strconv.Atoi(*num)
        numberMap[Number{row, *start, end}] = n
        *start = -1
        *num = ""
    }
}

func isGear(r int, c int, grid []string) bool {
    if grid[r][c] != '*' {
        return false
    }

    count := 0
    if r > 0 {
        if '0' <= grid[r - 1][c] && grid[r - 1][c] <= '9' {
            count++
        } else {
            count += getAdjacentColumnsCount(r - 1, c, grid)
        }
    }
    if r < len(grid) {
        if '0' <= grid[r + 1][c] && grid[r + 1][c] <= '9' {
            count++
        } else {
            count += getAdjacentColumnsCount(r + 1, c, grid)
        }
    }
    
    count += getAdjacentColumnsCount(r, c, grid)

    return count == 2
}

func getAdjacentColumnsCount(r int, c int, grid []string) int {
    count := 0
    if c > 0 && '0' <= grid[r][c - 1] && grid[r][c - 1] <= '9' {
        count++
    }
    if c < len(grid[0]) && '0' <= grid[r][c + 1] && grid[r][c + 1] <= '9' {
        count++
    }
    return count
}

func addNewGear(x int, y int, gearMap map[Coord][]int, gears *[]int) {
    // map each adjacent point on the grid to corresponding gears
    dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
    for i := 0; i < len(dx); i++ {
        coord := Coord{y: y + dy[i], x: x + dx[i]}
        adjGears, ok := gearMap[coord]
        if ok {
            gearMap[coord] = append(adjGears, len(*gears))
        } else {
            gearMap[coord] = []int{len(*gears)}
        }
    }
    // add new gear
    *gears = append(*gears, 1)
}
