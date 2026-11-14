package main
import "strings"

func cleanInput(text string) []string {
	cleanSlice := []string{}
	cleanSlice = strings.Fields(strings.ToLower(text))
	return  cleanSlice
}
