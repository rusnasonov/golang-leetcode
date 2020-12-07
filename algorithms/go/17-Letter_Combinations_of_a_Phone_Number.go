package main

import "sort"

var digitToChar = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	totalCombinations := 1
	for _, digit := range digits {
		totalCombinations *= len(digitToChar[string(digit)])
	}
	combinations := make([]string, totalCombinations)
	chunkLen := totalCombinations
	for _, digit := range digits {
		letters := digitToChar[string(digit)]
		chunkLen = chunkLen / len(letters)
		for k := 0; k < len(combinations); {
			for _, letter := range letters {
				for i := 0; i < chunkLen; i++ {
					combinations[k] += string(letter)
					k++
				}
			}
		}
	}
	sort.Strings(combinations)
	return combinations
}
