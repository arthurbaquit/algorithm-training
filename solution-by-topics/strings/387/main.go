package main

func firstUniqChar(s string) int {
    res := make(map[rune]int)
    for i, v := range(s){
        if _, ok := res[v]; ok{
            res[v] = -1
            continue
        }
        res[v] = i
    }
    for _, v := range(s){
        if res[v] != -1 { return res[v]}
    }
    return -1
}

func main(){
	println(firstUniqChar("leetcode"))
	println(firstUniqChar("loveleetcode"))
	println(firstUniqChar("aabb"))
}