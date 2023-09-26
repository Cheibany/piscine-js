package main

import "fmt"

func giveMeTheEquivalent(i int, s string) byte {
	return s[i]
}

func giveMeTheIndex(c rune, s string) int {
	for i, v := range s {
		if v == c {
			return i
		}
	}
	return -1
}

func Rot14(s string) string {
	alphabet14 := "opqrstuvwxyzabcdefghijklmn"
	Alphabet14 := "OPQRSTUVWXYZABCDEFGHIJKLMN"
	str := ""
	for _, v := range s {
		if v >= 97 && v <= 122 {
			str += string(giveMeTheEquivalent(giveMeTheIndex(v, "abcdefghijklmnopqrstuvwxyz"), alphabet14))
		} else if v >= 65 && v <= 122 {
			str += string(giveMeTheEquivalent(giveMeTheIndex(v, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"), Alphabet14))
		} else {
			str += string(v)
		}
	}
	return str
}

func main() {
	fmt.Println(Rot14("hello how are you"))
}
