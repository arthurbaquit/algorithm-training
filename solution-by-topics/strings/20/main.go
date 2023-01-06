package main

func isValid(s string) bool {
    par := map[rune]rune{'(': ')', '[':']', '{':'}'}
    queue := make([]rune, 0)
    var q rune
    for _, v := range(s) {
        if v == ')' ||  v == ']' ||  v == '}' {
            if len(queue) == 0 {return false}
            q, queue = queue[0], queue[1:]
            if par[q] != v { return false}
            continue
        }
        queue = append([]rune{v}, queue...)
    }
    return len(queue)==0
}