package main

func canConstruct(ransomNote string, magazine string) bool {
    letters := make(map[rune]int)
    for _,v := range(magazine) {
        if _, ok := letters[v]; ok{
            letters[v]++
            continue
        }
        letters[v] = 1
    }
    for _, v := range(ransomNote) {
        if _, ok := letters[v]; ok{
            if letters[v] == 0 { return false}
            letters[v]--
            continue
        }
        return false
    }
    return true
}

func main(){
	println(canConstruct("a", "b"))
	println(canConstruct("aa", "ab"))
	println(canConstruct("aa", "aab"))
}