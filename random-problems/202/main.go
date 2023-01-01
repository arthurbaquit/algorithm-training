package main

func isHappy(n int) bool {
    count := 0
    for count!=50 {
        aux:=0 
        for n>0 {
            aux +=(n%10)*(n%10)
            n/=10
        }
        n = aux
        if n == 1 { return true }
        count++
    }
    return false
}

func main(){
	println(isHappy(19))
	println(isHappy(2))
}