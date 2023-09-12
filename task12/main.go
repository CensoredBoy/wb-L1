package main

import "fmt"

//реализация множества через хеш таблицу

func newSet(slice []string) []string {
	setMap := make(map[string]bool)
	result := []string{}
	for _, key := range slice {
		setMap[key] = true
	}
	for key, _ := range setMap {
		result = append(result, key)
	}
	return result
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := newSet(slice)
	fmt.Println(set)
}
