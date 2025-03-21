package main

import "fmt"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	suppliesDict := make(map[string]bool)
	ans := make(map[string]bool)
	for _, supply := range supplies {
		suppliesDict[supply] = true
	}
	for i := len(recipes) - 1; i >= 0; i-- {
		if _, ok := ans[recipes[i]]; ok {
			continue
		}
		canBeBaked := true
		for _, ingredient := range ingredients[i] {
			if _, ok := suppliesDict[ingredient]; ok {
				continue
			}
			canBeBaked = false
			break
		}
		if canBeBaked {
			suppliesDict[recipes[i]] = true
			ans[recipes[i]] = true
			i = len(recipes)
		}
	}
	ansArr := make([]string, len(ans))
	count := 0
	for key, _ := range ans {
		ansArr[count] = key
		count++
	}
	return ansArr
}

func main() {
	recipes := []string{"bread"}
	ingredients := [][]string{{"yeast", "flour"}}
	supplies := []string{"yeast", "flour", "corn"}
	fmt.Println(findAllRecipes(recipes, ingredients, supplies)) // [bread]
}
