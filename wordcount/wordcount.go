package wordCount

import (
	"fmt"
	"strings"
)

// 函数写法，并被别的包调用
func CountWords(s string) {
	words := strings.Fields(s)
	sum := make(map[string]int)
	for _, word := range words {
		sum[word]++
	}
	fmt.Println(sum)
}
