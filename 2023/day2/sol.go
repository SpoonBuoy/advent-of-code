package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	Red   = 12
	Green = 13
	Blue  = 14
)

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func solve(line string, cnt int) int {
	newLine := strings.Split(line, ":")
	drawn := strings.Split(newLine[1], ";")
	var (
		blue  = 0
		red   = 0
		green = 0
	)
	for _, once := range drawn {
		chunks := strings.Split(once, " ")

		for i := 1; i <= len(chunks)-2; i++ {
			color := chunks[i+1]
			num, _ := strconv.Atoi(chunks[i])
			//fmt.Println(num)
			if strings.Contains(color, "blue") {
				blue = max(blue, num)
			} else if strings.Contains(color, "red") {
				red = max(red, num)
			} else if strings.Contains(color, "green") {
				green = max(green, num)
			}
			i++
		}
	}
	if red <= Red && blue <= Blue && green <= Green {
		return cnt
	}
	return 0
}

func main() {
	file, _ := os.Open("day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0
	cnt := 1
	for scanner.Scan() {
		line := scanner.Text()
		ans += solve(line, cnt)
		cnt++
	}
	fmt.Println(ans)

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
