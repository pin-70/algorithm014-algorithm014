package Week_03

import (
	"fmt"
	"testing"
)

var combinations []string
var mp map[string]string

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}
	mp = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	combinations = []string{}
	rescur(digits, 0, "")
	return combinations
}

func rescur(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	} else {
		digit := string(digits[index])
		letters := mp[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			rescur(digits, index+1, combination+string(letters[i]))
		}
	}
}

func Test_aaaaa(t *testing.T) {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
