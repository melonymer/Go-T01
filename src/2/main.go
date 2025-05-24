package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	input, _ := stdin.ReadString('\n')

	var k int
	fmt.Scanln(&k)

	if input == "" {
		fmt.Fprint(os.Stdout, "")
		return
	}

	words := strings.Fields(input)

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	uniqueWords := make([]string, 0, len(freq))
	for word := range freq {
		uniqueWords = append(uniqueWords, word)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		if freq[uniqueWords[i]] == freq[uniqueWords[j]] {
			return uniqueWords[i] < uniqueWords[j]
		}
		return freq[uniqueWords[i]] > freq[uniqueWords[j]]
	})

	if k > len(uniqueWords) {
		k = len(uniqueWords)
	}
	result := uniqueWords[:k]

	fmt.Println(strings.Join(result, " "))
}