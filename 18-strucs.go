package main

import "fmt"

type persona struct {
	name string
	age int
}

func main() {
	fmt.Println(persona{"bob",20})
	fmt.Println(persona{name:"alice", age:30})
	fmt.Println(&persona{"ann",40})
	s:= persona{name: "sean", age: 50}
	fmt.Println(s.name)
	sp:= &s
	fmt.Println(sp.age)
	sp.age=51
	fmt.Println(sp.age)

}