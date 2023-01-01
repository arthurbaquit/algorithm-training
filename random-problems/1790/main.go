package main

func areAlmostEqual(s1 string, s2 string) bool {
    l := len(s1)
    first,second, count := -1,-1,0
    for i:=0; i<l;i++{
        if s1[i] != s2[i]{
            count++
            if count>2 { return false}
            if first == -1 { first = i; continue}
            second = i
        }   
    }
    if count == 0 {return true}
     if first != -1 && second != -1 && s1[first] == s2[second] && s1[second] == s2[first] {
        return true
    }
    
    return false
}

func main(){
	println(areAlmostEqual("bank","kanb"))
	println(areAlmostEqual("attack","defend"))
	println(areAlmostEqual("kelb","kelb"))
	println(areAlmostEqual("abcd","dcba"))
}