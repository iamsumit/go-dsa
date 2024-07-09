package main

import (
	"fmt"
	"strings"
)

// https://www.naukri.com/code360/problems/983635?topList=striver-sde-sheet-problems&utm_source=striver&utm_medium=website
func wordBreak(s string, dictionary []string) []string {
	output := []string{}
	sentence := ""

	var crackDict func(str string)
	crackDict = func(str string) {
		if str == "" {
			output = append(output, strings.TrimRight(sentence, " "))
			return
		}

		for _, word := range dictionary {
			for i := 0; i <= len(str); i++ {
				if str[0:i] == word {
					sentence += word + " "
					crackDict(str[i:])
					sentence = sentence[:len(sentence)-len(word)-1]
				}
			}
		}
	}

	crackDict(s)

	return output
}

func main() {
	fmt.Println(wordBreak("godisnowherenowhere", []string{"god", "is", "now", "no", "where", "here"}))
	fmt.Println(wordBreak("godisnowhere", []string{"god", "is", "no", "here"}))
	fmt.Println(wordBreak("ilikesamsungmobile", []string{"i", "like", "sam", "sung", "samsung", "mobile", "ice", "and", "cream", "icecream", "man", "go", "mango"}))
	fmt.Println(wordBreak("ilikeicecreamandmango", []string{"i", "like", "sam", "sung", "samsung", "mobile", "ice", "and", "cream", "icecream", "man", "go", "mango"}))
}
