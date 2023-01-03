package main

func lengthOfLongestSubstring(s string) int {
    prev, curr, max := 0, 0, 0 
    occur := make(map[byte]int)

    for curr < len(s) {
        if val, ok := occur[s[curr]]; ok{
            max = Max(max, curr-prev)
            prev= Max(val+1, prev)
        }
        occur[s[curr]] = curr
        curr++
    }
    max = Max(max, curr-prev)
    return max
}

func Max(x, y int) int {
    if x > y {return x}
    return y
}

func main(){
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring("au"))
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring("tmmzuxt"))
}