package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
    
var cubeLims = map[string] int {
    "red": 12,
    "green": 13,
    "blue": 14,
}

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
        game, _ := strconv.Atoi(strings.Split(subsets[0], " ")[1])
        picks := strings.Split(subsets[1], "; ")
        if checkValid(picks) {
            sum += game
        }
    }
    fmt.Println(sum)
}

func checkValid(picks []string) bool {
    for i := 0; i < len(picks); i++ {
        for _, cube := range strings.Split(picks[i], ", ") {
            words := strings.Split(cube, " ")
            i, _ := strconv.Atoi(words[0])
            if i > cubeLims[words[1]] {
                return false
            }
        }
    }
    return true
}
