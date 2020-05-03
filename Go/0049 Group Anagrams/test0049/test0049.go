package test0049

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type testFunc func([]string) [][]string

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals([]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}})
	equals(readInput("./tests/test1input.txt"), readOutput("./tests/test1output.txt"))

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func readInput(fileName string) []string {
	fileContent, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(fileContent)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	fileContent.Close()
	return fileTextLines
}

func readOutput(fileName string) [][]string {
	fileContent, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(fileContent)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines [][]string

	for fileScanner.Scan() {
		text := fileScanner.Text()
		fileTextLines = append(fileTextLines, strings.Split(text, " "))
	}
	fileContent.Close()
	return fileTextLines
}

func equals(words []string, expected [][]string) {
	numTests++
	got := testFunction(words)
	if !sliceEquals(got, expected) {
		fmt.Println(words, "got:", got, "expected:", expected)
		failed++
	}
}

func sliceEquals(slice1 [][]string, slice2 [][]string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}

		for j := 0; j < len(slice1[i]); j++ {
			ok := false
			for k := 0; k < len(slice1[i]); k++ {
				if slice1[i][j] == slice2[i][k] {
					ok = true
					break
				}
			}
			if !ok {
				return false
			}
		}
	}
	return true
}
