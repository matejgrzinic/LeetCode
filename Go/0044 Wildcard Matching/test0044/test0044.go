package test0044

import (
	"fmt"
	"time"
)

type testFunc func(string, string) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals("aa", "a", false)
	equals("aa", "*", true)
	equals("cb", "?a", false)
	equals("adceb", "a*b", true)
	equals("acdcb", "a*c?b", false)
	equals("adceb", "*a*b", true)
	equals("", "*", true)
	equals("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*", true)
	equals("abbaabbbbababaababababbabbbaaaabbbbaaabbbabaabbbbbabbbbabbabbaaabaaaabbbbbbaaabbabbbbababbbaaabbabbabb", "***b**a*a*b***b*a*b*bbb**baa*bba**b**bb***b*a*aab*a**", true)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

func equals(n string, k string, expected bool) {
	numTests++
	got := testFunction(n, k)
	if got != expected {
		fmt.Println(n, k, "got:", got, "expected:", expected)
		failed++
	}
}
