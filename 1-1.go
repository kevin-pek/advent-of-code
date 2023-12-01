package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("1-1.test")
	if err != nil {
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
		first, last := -1, -1
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if first == -1 {
					first = int(char - '0') * 10
					sum += first
				}
				last = int(char - '0')
			}
		}
		sum += last
	}
	fmt.Printf("Sum of calibration values: %d\n", sum) // 142
}

