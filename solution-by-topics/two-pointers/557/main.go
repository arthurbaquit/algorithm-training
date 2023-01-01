package main

import "strings"

func reverseWords(s string) string {
    init, end := 0, 0
    var sb strings.Builder

    l := len(s)
    for end<l{
        for end<l && s[end]!=' ' {end++}
        aux:= init
        init=end-1
        for aux <= init {
            sb.WriteString(string(s[init]))
            init--
        }
        if(end<l) { sb.WriteString(" ")}
        end++
        init = end
    }
    return sb.String()
}

func main(){
	println(reverseWords("Let's take LeetCode contest"))
}