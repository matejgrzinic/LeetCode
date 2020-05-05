package test0055

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type testFunc func([]int) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]int{2, 3, 1, 1, 4}, true)
	equals(readInput("./tests/test1input.txt"), false)
	equals([]int{3, 2, 1, 0, 4}, false)
	equals([]int{1, 3, 2}, true)
	equals([]int{2, 5, 0, 0}, true)
	equals([]int{1}, true)
	equals([]int{2, 0}, true)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(nums []int, expected bool) {
	numTests++
	got := testFunction(nums)
	if got != expected {
		fmt.Println(nums, "got:", got, "expected:", expected)
		failed++
	}
}

func readInput(fileName string) []int {
	fileContent, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(fileContent)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []int

	for fileScanner.Scan() {
		num, _ := strconv.Atoi(fileScanner.Text())
		fileTextLines = append(fileTextLines, num)
	}
	fileContent.Close()
	return fileTextLines
}
