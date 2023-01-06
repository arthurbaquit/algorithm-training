package main

func isAlienSorted(words []string, order string) bool {
    ordersMap := make(map[byte]int)
    for i:=0; i<26; i++ {
        ordersMap[order[i]]=i
    }
    for i:=1;i<len(words);i++{
        flag:= false
        for j:=0;j<len(words[i-1]) && j < len(words[i]);j++ {
                if ordersMap[words[i-1][j]] > ordersMap[words[i][j]] { return false }
                if ordersMap[words[i-1][j]] == ordersMap[words[i][j]] { continue }
                flag=true
                break
                }
                if flag {continue}
                if len(words[i]) < len(words[i-1]) { return false}

    }
    return true
}

func main(){
	println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz"))
	println(isAlienSorted([]string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz"))
	println(isAlienSorted([]string{"hello","hello"}, "abcdefghijklmnopqrstuvwxyz"))
	println(isAlienSorted([]string{"apaz","app"}, "abcdefghijklmnopqrstuvwxyz"))
}