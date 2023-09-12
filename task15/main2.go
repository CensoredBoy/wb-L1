package main

import "strings"
var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}

func someFunc() {

	//в исходной программе может быть переполнение стека из за большой строки
	//так как строки в го это байтовые слайсы только для чтения, поэтому меняя/присваивая строку
	//создается новый байтовый слайс. Так как у нас большие строки то может быть переполнение стека


	//Builder минимизирует копирование 
	var b strings.Builder
	b.Grow(100)
	b.WriteString(createHugeString(1 << 10)[:100])

	justString = b.String()


}
func main() {
	someFunc()
}
