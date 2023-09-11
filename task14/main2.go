package main

import "fmt"

func getType(data interface{}) string{
	switch data.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	case rune:
		return "rune"
	case bool:
		return "bool"
	case chan string:
		return "chan string"
	default:
		return "I don't know this type"
	}
}

func main() {
	var data []interface{}
	data = append(data, "strrr")
	data = append(data, 1337)
	data = append(data, 'd')
	data = append(data, true)
	data = append(data, 88.14)
	data = append(data, make(chan string))

	for _, val := range data {
		fmt.Println(getType(val))
	}

}
