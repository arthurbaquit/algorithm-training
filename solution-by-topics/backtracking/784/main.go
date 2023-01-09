package main

import "fmt"

func letterCasePermutation(s string) []string {
    res := []string{}
    var backtrack func(start int)
    var curr []byte
    backtrack = func(start int) {
        if len(curr) == len(s) {
            res = append(res, string(curr))
            return
        }
            if 'A'<=s[start] && s[start] <= 'Z' || 'a'<=s[start] && s[start] <= 'z'{

    curr = append(curr, toLower(s[start]))
    backtrack(start+1)
    curr = curr[:len(curr)-1]
    curr = append(curr, toUpper(s[start]))
    backtrack(start+1)
    curr = curr[:len(curr)-1]
  } else {
    curr = append(curr, s[start])
    backtrack(start+1)
    curr = curr[:len(curr)-1]
  }
}
    backtrack(0)
    return res
}

func toLower(ch byte) byte {
  if ch >= 'A' && ch <= 'Z' {
    return 'a' + ch-'A'
  }
  return ch
}

func toUpper(ch byte) byte {
  if ch >= 'a' && ch <= 'z' {
    return 'A' + ch - 'a'
  }
  return ch
}

func main(){
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
	fmt.Println(letterCasePermutation("12345"))
	fmt.Println(letterCasePermutation("0"))
}