package palindrome_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsPalindrome(word string) bool {
	// removing special char
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println(err)
	}
	alphaNum := reg.ReplaceAllString(word, "")

	// lowercasing
	toLower := strings.ToLower(alphaNum)
	
	// remove white space
	cleanWord := strings.ReplaceAll(toLower, " ", "")

	var revWord string
	for i := len(cleanWord) - 1; i >= 0; i-- {
		revWord += string(cleanWord[i])
	}
	return revWord == cleanWord
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct{
		name string
		value string
	}{
		{
			name: "IsPalindrome(SAIPPUAKIVIKAUPPIAS)",
			value: "SAIPPUAKIVIKAUPPIAS",
		},
		{
			name: "IsPalindrome(Aibohphobia)",
			value: "Aibohphobia",
		},
		{
			name: "IsPalindrome(Anna)",
			value: "Anna",
		},
		{
			name: "IsPalindrome(Civic)",
			value: "Civic",
		},
		{
			name: "IsPalindrome(My gym)",
			value: "My gym",
		},
		{
			name: "IsPalindrome(No lemon, no melon)",
			value: "No lemon, no melon",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPalindrome(test.value)
			assert.True(t, result)
		})
	}
}