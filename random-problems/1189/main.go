package main

func maxNumberOfBalloons(text string) int {
	balloonMap := make(map[rune]int)
	for _, r := range text {
		balloonMap[r]++
	}
	return min(balloonMap['n'], min(min(balloonMap['l']/2, balloonMap['o']/2), min(balloonMap['b'], balloonMap['a'])))

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	println(maxNumberOfBalloons("nlaebolko"))
	println(maxNumberOfBalloons("loonbalxballpoon"))
	println(maxNumberOfBalloons("leetcode"))
}
