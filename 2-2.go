package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2.in")
	if err != nil {
		return
	}
	defer file.Close()

    sum := 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        subsets := strings.Split(line, ": ")
        sets := strings.Split(subsets[1], "; ")
        sum += getPower(sets)
    }
    fmt.Println(sum)
}

func getPower(picks []string) int {
    colorMap := map[string] int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }

    for i := 0; i < len(picks); i++ {
        for _, cube := range strings.Split(picks[i], ", ") {
            words := strings.Split(cube, " ")
            i, _ := strconv.Atoi(words[0])
            if i > colorMap[words[1]] {
                colorMap[words[1]] = i
            }
        }
    }
    power := 1
    for _, n := range colorMap {
        power *= n
    }
    return power
}
