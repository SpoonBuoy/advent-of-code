package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func solve(line string) int64 {
	var res []rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			res = append(res, c)
		}
	}
	num := (res[0]-'0')*10 + res[len(res)-1] - '0'

	return int64(num)
}

func main() {
	file, _ := os.Open("day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ans int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		ans += solve(line)
	}
	fmt.Println(ans)

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
