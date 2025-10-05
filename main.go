package main

import (
	// "piss/Map"
	"fmt"
	"math/rand/v2"
	"time"
)

var RNG *rand.Rand

func init() { // v2 is pretty cool and shit bruh
	var Unix uint64 = uint64(time.Now().UnixMicro())
	RandUnix := func() uint64 { return Unix - rand.Uint64N(Unix) }
	RNG = rand.New(rand.NewPCG(RandUnix(), RandUnix()))

}

func main() {
	fmt.Println(RNG.Int64N(100))
}

// func main() {
// 	x := Map.Any{}
//
// 	for i := 0; i <= 100; i++ {
//
// 		x.Set("piss", "milk")
// 	}
//
// 	k := x.Get("piss")
// 	fmt.Println(k)
// }
