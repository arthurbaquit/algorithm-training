package main

func clearDigits(s string) string {
	l := 0
	for l < len(s) && l >= 0 {
		if s[l] >= 'a' && s[l] <= 'z' {
			l++
			continue
		}
		s = s[:l-1] + s[l+1:]
		l = l - 1
	}
	return s
}

type Queue struct {
	len    int
	maxLen int
	digits []byte
}

func (q *Queue) Push(x byte) {
	if q.len > q.maxLen {
		return
	}
	q.len++
	q.digits[q.len-1] = x
}

func (q *Queue) Pop() byte {
	if q.len == 0 {
		return byte(0)
	}
	q.len--
	return q.digits[q.len+1]
}

func (q *Queue) Flush() string {
	return string(q.digits[:q.len])
}

func clearDigitsQueue(s string) string {
	queue := &Queue{
		len:    0,
		maxLen: 100,
		digits: make([]byte, 100),
	}
	for _, b := range []byte(s) {
		if b >= 'a' && b <= 'z' {
			queue.Push(b)
			continue
		}
		_ = queue.Pop()
	}
	return queue.Flush()
}

func main() {
	println(clearDigits("abc"))       // abc
	println(clearDigitsQueue("abc"))  // abc
	println(clearDigits("cb34"))      // ""
	println(clearDigitsQueue("cb34")) // ""
}
