package main

func mergeAlternately(word1 string, word2 string) string {
    w1, w2 := 0, 0
    res := ""
    for w1 < len(word1) && w2 < len(word2){
        res +=(string(word1[w1])+string(word2[w2]))
        w1++
        w2++
    }
    if w1 < len(word1){res+=word1[w1:]}
    if w2 < len(word2){res+=word2[w1:]}
    return res
}

func main(){
	println(mergeAlternately("abc", "pqr"))
	println(mergeAlternately("ab", "pqrs"))
	println(mergeAlternately("abcd", "pq"))
}