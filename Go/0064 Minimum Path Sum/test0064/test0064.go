package test0064

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type testFunc func([][]int) int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7)
	equals(readInput("./tests/test1input.txt"), 83)
	equals(readInput("./tests/test2input.txt"), 91)
	equals(readInput("./tests/test3input.txt"), 83)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(matrix [][]int, expected int) {
	numTests++
	got := testFunction(matrix)
	if got != expected {
		fmt.Println(matrix, "got:", got, "expected:", expected)
		failed++
	}
}

func readInput(fileName string) [][]int {
	fileContent, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(fileContent)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines [][]int

	for fileScanner.Scan() {
		row := fileScanner.Text()
		nums := strings.Split(row, ",")
		numRow := []int{}
		for _, val := range nums {
			n, _ := strconv.Atoi(val)
			numRow = append(numRow, n)
		}
		fileTextLines = append(fileTextLines, numRow)
	}
	fileContent.Close()
	return fileTextLines
}
