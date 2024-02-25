package main

type Logger struct {
	LastCalled map[string]int
}

func Constructor() Logger {
	return Logger{
		LastCalled: make(map[string]int),
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if value, ok := this.LastCalled[message]; ok {
		if value+10 > timestamp {
			return false
		}
	}
	this.LastCalled[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
func main() {
	logger := Constructor()
	println(logger.ShouldPrintMessage(1, "foo"))
	println(logger.ShouldPrintMessage(2, "bar"))
	println(logger.ShouldPrintMessage(3, "foo"))
	println(logger.ShouldPrintMessage(8, "bar"))
	println(logger.ShouldPrintMessage(10, "foo"))
	println(logger.ShouldPrintMessage(11, "foo"))
}
