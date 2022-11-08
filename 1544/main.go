package main

func makeGood(s string) string {
	point1, point2, l := 0, 1, len(s)
	if len(s) == 1 {
		return s
	}
	for point1 <= l-2 {
		diff := s[point1] - s[point2]
		if diff == 32 || diff == 224 {
			s = s[0:point1] + s[point2+1:]
			l = l - 2
			if point1 != 0 {
				point1--
			}
			if point2 != 1 {
				point2--
			}
			continue
		}
		point1++
		point2++
	}
	return s
}

func main() {
	s := "leEeetcode"
	println(makeGood(s))
	s = "abBAcC"
	println(makeGood(s))
	s = "s"
	println(makeGood(s))
	s = "Pp"
	println(makeGood(s))
}
