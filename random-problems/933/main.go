package main

type RecentCounter struct {
	lastTs []int
}

func Constructor() RecentCounter {
	return RecentCounter{lastTs: []int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.lastTs = append(rc.lastTs, t)
	for i := 0; i < len(rc.lastTs); i++ {
		if t-rc.lastTs[i] > 3000 {
			continue
		}
		rc.lastTs = rc.lastTs[i:]
		break
	}
	return len(rc.lastTs)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
func main() {
	rc := Constructor()
	println(rc.Ping(1))
	println(rc.Ping(100))
	println(rc.Ping(3001))
	println(rc.Ping(3002))
}
