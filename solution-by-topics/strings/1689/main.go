package main

func minPartitions(n string) int {
	max := -1
	for _, v := range n {
		if int(v-'0') > max {
			max = int(v - '0')
			continue
		}
	}
	return max
}

func main() {
	println(minPartitions("32"))
	println(minPartitions("82734"))
	println(minPartitions("27346209830709182346"))
}
