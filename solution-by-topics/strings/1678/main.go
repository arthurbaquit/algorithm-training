package main

func interpret(command string) string {
    res := ""
    for i:=0; i<len(command); i++{
        if command[i] == 'G' {res+="G"; continue}
        if command[i] == ')' && command[i-1] == '(' {res += "o"; continue}
        if command[i] == ')' && command[i-1] == 'l' {res += "al"; continue}
    }
    return res
}

func main(){
	println(interpret("G()(al)"))
	println(interpret("G()()()()(al)"))
	println(interpret("(al)G(al)()()G"))
}