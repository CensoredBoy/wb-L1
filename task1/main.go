package main

import "fmt"

type Human struct  { 
	name string
	age uint
	height uint
	weight uint
}

func( h *Human) aging() uint{ 
	h.age += 1
	return h.age
}

func( h *Human) grow() uint{ 
	h.height += 1
	return h.height
}

func (h *Human) getting_fat() uint{ 
	h.weight += 1
	return h.weight
}

func (h Human) String() string { 
	return fmt.Sprintf(" Name: %v \n Age: %v \n Height: %v \n Weight: %v \n ", h.name, h.age, h.height, h.weight)
}

type Action struct { 
	Human // Встраивание типа
}


func main(){ 
	act := Action{Human{name:"Danila", age: 20, height: 186, weight: 1000}} //Создание экземпляра act типа Action с встроенным типом Human
	act.aging()
	act.grow()
	act.getting_fat()
	fmt.Println(act)

}
