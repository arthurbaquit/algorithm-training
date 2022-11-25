package main

import "strings"

func defangIPaddrByHand(address string) string {
	var res string
	for _, v := range address {
		if v == '.' {
			res += "[.]"
			continue
		}
		res += string(v)
	}
	return res
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	println(defangIPaddr("1.1.1.1"))
	println(defangIPaddr("255.100.50.0"))
	println(defangIPaddrByHand("1.1.1.1"))
	println(defangIPaddrByHand("255.100.50.0"))

}
