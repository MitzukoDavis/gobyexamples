package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i , " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")

	}
	t:= time.Now()
	switch {
	case t.Hour() <12:
		fmt.Println("it's before noon")
		fmt.Println(t.Hour())
		fmt.Println(t.Minute())
	default:
		fmt.Print("it's after noon")

	}


}