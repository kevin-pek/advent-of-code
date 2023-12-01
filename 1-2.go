package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode"
)

var numMap = map[string] int {
	"one":   1,
    "two":   2,
    "three": 3,
    "four":  4,
    "five":  5,
    "six":   6,
    "seven": 7,
    "eight": 8,
    "nine":  9,
}

func main() {
    sum := sumValues("1.in") // 53894
	fmt.Printf("Sum of calibration values: %d\n", sum)
}

func sumValues(filepath string) int {
	file, _ := os.Open(filepath)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first, last := -1, -1
        for i := 0; i < len(line); i++ {
            if unicode.IsDigit(rune(line[i])) {
                digit := int(line[i] - '0')
                if first == -1 {
                    first = digit
                }
                last = digit
            } else {
                segment := line[i:]
                for word, num := range numMap {
                    if strings.HasPrefix(segment, word) {
                        if first == -1 {
                            first = num
                        }
                        last = num
                        i += len(word) - 2 // jump index ahead but leave the last character of word in
                    }
                }
            }
        }
		sum += first * 10 + last
	}
    return sum
}

func TestPartTwo(t *testing.T) {
    sum := sumValues("1-2.test")
    if sum != 281 {
        t.Errorf("Expected %d, got %d", 281, sum)
    }
}
