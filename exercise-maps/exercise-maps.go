package main

import (
	"strings"
	"golang.org/x/tour/wc"
)
//WordCount 统计单词
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	origin := strings.Fields(s)
	for _,value := range origin {
		result[value]++
	}
	return result
}
func main() {
	wc.Test(WordCount)
}