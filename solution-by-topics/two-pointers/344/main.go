package main

func reverseString(s []byte)  {
    init, end := 0, len(s)-1
    for init<end {
        aux:= s[init]
        s[init] = s[end]
        s[end] = aux
        end--
        init++
    }
}

func main(){
	s := []byte{'h','e','l','l','o'}
	reverseString(s)
	println(string(s))
	s = []byte{'H','a','n','n','a','h'}
	reverseString(s)
	println(string(s))
}