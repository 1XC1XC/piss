package structs 

import (
	"fmt"
	"piss/Map"
)

func main() {
	x := Map.Any{}
	x.Set("milk", 123)

	k := x.Get("milk")
	fmt.Println(k)
}
