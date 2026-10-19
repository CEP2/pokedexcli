package main

import "strings"

func cleanInput(text string) []string {
	result := text
	result = strings.Trim(text, " ")
	result = strings.ToLower(result)

	fieldsResult := strings.Fields(result)

	return fieldsResult
}
