package main

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
	v := make([]byte, 100, 100)
  	copy(v, []byte(createHugeString(1 << 10)[:100]))
	justString = string(v)
}

func main() {
	someFunc()
}
