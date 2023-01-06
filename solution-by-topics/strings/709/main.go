package main

func toLowerCase(s string) string {
    res:= ""
    for i:=0; i<len(s);i++{
        if s[i] <= 'Z' && s[i] >= 'A' { res += string(s[i]+32); continue}
        res+=string(s[i])
    }
    return res
}

func main(){
	println(toLowerCase("Hello"))
	println(toLowerCase("here"))
	println(toLowerCase("LOVELY"))
	println(toLowerCase("al&phaBET"))
}