package main

import (
	"regexp"
	"strings"

	"./test0071"
)

func main() {
	test0071.Test(simplifyPath)
}

func simplifyPath(path string) string {
	r := regexp.MustCompile("/+")
	locations := r.Split(path, -1)

	shortest := []string{}
	for _, loc := range locations {
		if loc != "." {
			if loc == ".." {
				if len(shortest) > 0 {
					shortest = shortest[:len(shortest)-1]
				}
			} else {
				shortest = append(shortest, loc)
			}
		}
	}
	shortestClean := []string{}
	for _, val := range shortest {
		if val != "" {
			shortestClean = append(shortestClean, val)
		}
	}
	return "/" + strings.Join(shortestClean, "/")
}
