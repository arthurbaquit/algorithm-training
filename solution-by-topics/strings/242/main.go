package main

func isAnagram(s string, t string) bool {
    letters := make(map[rune]int)
    for _,v := range(s) {
        if _, ok := letters[v]; ok{
            letters[v]++
            continue
        }
        letters[v] = 1
    }
    for _, v := range(t) {
        if _, ok := letters[v]; ok{
            if letters[v] == 0 { return false}
            letters[v]--
            continue
        }
        return false
    }
    for _, v := range(letters) {
        if v != 0{
             return false
        }
    }
    return true
}

func main(){
	println(isAnagram("anagram", "nagaram"))
	println(isAnagram("rat", "car"))
}