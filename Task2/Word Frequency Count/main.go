package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	left, right := 0, len(word)-1

	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	title := "Palindrome Check"
	fmt.Println("\n" + title)
	fmt.Println(strings.Repeat("-", len(title))) // Underline with dashes

	var word string
	fmt.Print("Enter the string: ")
	fmt.Scan(&word)

	if isPalindrome(word) {
		fmt.Printf("\n\"%s\" is a palindrome.\n", word)
	} else {
		fmt.Printf("\n\"%s\" is not a palindrome.\n", word)
	}
}