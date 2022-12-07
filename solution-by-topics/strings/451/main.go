package main

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	freq := make(map[rune]int)
	for _, v := range s {
		_, ok := freq[v]
		if !ok {
			freq[v] = 1
			continue
		}
		freq[v]++
	}
	res := make([]rune, 0, len(freq))

	for key := range freq {
		res = append(res, key)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return freq[res[i]] > freq[res[j]]
	})
	var sb strings.Builder
	for _, v := range res {
		for j := 0; j < freq[v]; j++ {
			sb.WriteString(string(v))
		}
	}
	return sb.String()
}

func main() {
	println(frequencySort("tree"))
	println(frequencySort("cccaaa"))
	println(frequencySort("Aabb"))
}
