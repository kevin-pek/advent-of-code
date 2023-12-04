package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, err := os.Open("4.in")
	if err != nil {
		return
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
    sum, cards, cardNum := 0, 0, 1
    cardCountMap := make(map[int]int)
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, ":")
        numbers := strings.Split(split[1], "|")
        winningNumbers := strings.Fields(numbers[0])
        givenNumbers := strings.Fields(numbers[1])
        winningNumbersMap := make(map[string]bool)
        for i := range winningNumbers {
            winningNumbersMap[winningNumbers[i]] = true
        }
        matches := 0
        for i := range givenNumbers {
            if winningNumbersMap[givenNumbers[i]] {
                matches++
            }
        }

        cardCountMap[cardNum]++
        if matches > 0 {
            sum += 1 << (matches - 1)
            for i := 1; i <= matches; i++ {
                cardCountMap[cardNum + i] += cardCountMap[cardNum]
            }
        }
        cardNum++
    }
    for _, count := range cardCountMap {
        cards += count
    }
    fmt.Printf("Part 1: %d\n", sum)
    fmt.Printf("Part 2: %d\n", cards)
}
