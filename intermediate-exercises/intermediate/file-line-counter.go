package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileLineCounter() {
	dat, err := os.ReadFile("./intermediate/industry.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileContent := string(dat)
	lines := len(strings.Split(strings.TrimSpace(fileContent), "\n"))
	fmt.Printf("File lines count: %d\n", lines)
}

func FileLineCounterOptimized() {
	file, err := os.Open("./intermediate/industry.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	firstNonEmpty := ""
	longestLine := ""
	emptyLines := 0
	for scanner.Scan() {
		lines++
		current := scanner.Text()
		current = strings.TrimSpace(current)
		if current != "" && firstNonEmpty == "" {
			firstNonEmpty = current
		}
		if len(current) > len(longestLine) {
			longestLine = current
		}
		if current == "" {
			emptyLines++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}

	fmt.Printf("File lines count: %d\n", lines)
	fmt.Printf("First non empty: %s\n", firstNonEmpty)
	fmt.Printf("Longest line: %s\n", longestLine)
	fmt.Printf("Empty lines: %d\n", emptyLines)
}
