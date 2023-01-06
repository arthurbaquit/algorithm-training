package main

import "strings"

func freqAlphabets(s string) string {
    l := len(s)
    curr, future := 0, 2
    var res strings.Builder
    replace := 0
    for future < l {
        flag := false
        if s[future] == '#' {flag = true}
        if flag {
            replace = 10*int(s[curr]-48)+ int(s[curr+1]-48)
            res.WriteRune(rune('a' + replace -1))
            curr = future+1
            future = curr + 2
            continue

        } 
        replace = int(s[curr]-48)
        res.WriteRune(rune('a' + replace -1))
        curr++
        future++
    }
    for i:=curr; i<len(s);i++{
        res.WriteRune(rune('a' + int(s[i])-49))
    }
    return res.String()
}

func main(){
	println(freqAlphabets("10#11#12"))
	println(freqAlphabets("1326#"))
	println(freqAlphabets("824#15#826#"))
}