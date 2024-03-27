package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

//try 1

func IsAnnagrams(word1 string, word2 string) bool {
	dict := make(map[rune]int)
	word1Lower := strings.ToLower(word1)
	word2Lower := strings.ToLower(word2)
	for _, v := range word1Lower {
		dict[v]++
	}

	for _, v := range word2Lower {
		if count, ok := dict[v]; ok && count-1 >= 0 {
			dict[v]--
		} else {
			return false
		}
	}
	return true
}

func groupAnnagrams(wordsList []string) [][]string {
	result := [][]string{}

	for _, word := range wordsList {
		isFind := false
		for i := 0; i < len(result); i++ {
			if IsAnnagrams(word, result[i][0]) {
				result[i] = append(result[i], word)
				isFind = true
			}
		}
		if !isFind {
			result = append(result, []string{word})
		}

	}

	sortFunc := func(i, j int) bool {
		return utf8.RuneCountInString(result[i][0]) > utf8.RuneCountInString(result[j][0])
	}
	sort.Slice(result, sortFunc)
	return result
}

func main() {
	words := []string{"материк", "клоун", "метрика", "уклон", "Керамит", "норка", "колун", "рок", "карета", "ракета"}
	res := groupAnnagrams(words)
	for _, group := range res {
		fmt.Println(group)
	}
}
