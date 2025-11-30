package utils

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Regular expressions for word splitting
var (
	// Matches uppercase letters followed by lowercase letters/numbers, or sequences of lowercase letters/numbers
	defaultRegex = regexp.MustCompile(`([A-Z][a-z0-9]*)|([a-z0-9]+)|([A-Z]+[A-Z][a-z])|([A-Z]+\d)|([A-Z]+\s)|([A-Z]+$)`)

	// Matches special characters and whitespace
	specialCharsRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
)

// convertToWords splits a string into words based on case boundaries and special characters
func convertToWords(str string) []string {
	// First, split on special characters and whitespace
	parts := specialCharsRegex.Split(str, -1)
	var words []string

	// Process each part
	for _, part := range parts {
		if part == "" {
			continue
		}

		// Split on case boundaries
		matches := defaultRegex.FindAllString(part, -1)
		words = append(words, matches...)
	}

	return words
}

func TitleCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	titleCaser := cases.Title(language.English)
	for i, word := range words {
		words[i] = titleCaser.String(word)
	}
	return strings.Join(words, " ")
}

func ConstantCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	return strings.ToUpper(strings.Join(words, "_"))
}

func PascalCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	titleCaser := cases.Title(language.English)
	for i, word := range words {
		words[i] = titleCaser.String(word)
	}
	return strings.Join(words, "")
}

func CamelCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	titleCaser := cases.Title(language.English)

	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = titleCaser.String(word)
		}
	}
	return strings.Join(words, "")
}

func UpperCase(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(strings.Join(convertToWords(str), ""))
}

func LowerCase(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToLower(strings.Join(convertToWords(str), ""))
}

func SnakeCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	return strings.ToLower(strings.Join(words, "_"))
}

func KebabCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	return strings.ToLower(strings.Join(words, "-"))
}

func FlatCase(str string) string {
	if str == "" {
		return ""
	}
	words := convertToWords(str)
	return strings.ToLower(strings.Join(words, " "))
}
