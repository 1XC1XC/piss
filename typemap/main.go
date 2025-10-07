package main

import (
	"fmt"
	"structs/Map"
)

func main() {
	x := Map.Any{}
	x.Set("milk", 123)

	k := x.Get("milk")
	fmt.Println(k)
}
