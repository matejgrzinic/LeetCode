package main

func destCity(paths [][]string) string {
	cityIn := make(map[string]bool)
	cityOut := make(map[string]bool)
	for _, path := range paths {
		cityOut[path[0]] = true
		cityIn[path[1]] = true

	}
	for k := range cityIn {
		if !cityOut[k] {
			return k
		}
	}
	return ""
}
