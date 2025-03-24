package main

import "sort"

func countDays(days int, meetings [][]int) int {
	dayMap := make(map[int]int)
	prefixSum := 0
	freeDays := 0
	previousDay := days
	for _, meeting := range meetings {
		start, finish := meeting[0], meeting[1]
		previousDay = min(previousDay, start)
		dayMap[start] += 1
		dayMap[finish+1] -= 1
	}
	freeDays += previousDay - 1
	mapKeys := make([]int, len(dayMap))
	for key := range dayMap {
		mapKeys = append(mapKeys, key)
	}
	sort.Ints(mapKeys)
	for _, key := range mapKeys {
		currDay, accSumThatDay := key, dayMap[key]
		if prefixSum == 0 {
			freeDays += currDay - previousDay
		}
		prefixSum += accSumThatDay
		previousDay = currDay
	}

	freeDays += days - previousDay + 1
	return freeDays
}

func main() {
	println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})) // 2
	println(countDays(5, [][]int{{2, 4}, {1, 3}}))           // 1
}
