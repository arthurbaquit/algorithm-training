package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * This code may not be the most perfomatic in leetcode,
 * but, in my personal experience, the solutions there uses
 * links that are not used in real life, as url/1, so, I think that
 * my solution is better to practice real life problems.
 */

type Codec struct {
	hashUrl map[string]string
}

func Constructor() Codec {
	return Codec{
		hashUrl: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c Codec) encode(longUrl string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	l := len(alphabet)
	rand.Seed(time.Now().UnixNano())
	// Getting random character
	tinyurl := "http://tinyurl.com/"
	for i := 0; i <= 6; i++ {
		tinyurl += string(alphabet[rand.Intn(l)])
	}
	c.hashUrl[tinyurl] = longUrl
	return tinyurl
}

// Decodes a shortened URL to its original URL.
func (c Codec) decode(shortUrl string) string {
	return c.hashUrl[shortUrl]
}

func main() {
	obj := Constructor()
	url := obj.encode("https://leetcode.com/problems/design-tinyurl")
	ans := obj.decode(url)
	fmt.Println(ans)
}
