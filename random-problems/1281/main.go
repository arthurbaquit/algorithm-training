package main

func subtractProductAndSum(n int) int {
    prod, sum := 1,0
    for n>0 {
        aux := n%10
        prod *= aux
        sum += aux
        n/=10
    }
    return prod-sum
}

func main(){
	println(subtractProductAndSum(234))
	println(subtractProductAndSum(4421))
}