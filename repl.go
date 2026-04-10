package main

import "strings"

func cleanInput(text string) []string {
	formatText := strings.ToLower(text)
	endresult := strings.Fields(formatText)
	return endresult
}
