package main

import (
	"fmt"
	// "math"
	"time"

	"math/rand/v2"
)

var RNG *rand.Rand

func init() {
	var Unix uint64 = uint64(time.Now().UnixMicro())
	RandUnix := func() uint64 { return Unix - rand.Uint64N(Unix) }
	RNG = rand.New(rand.NewPCG(RandUnix(), RandUnix()))

}

func main() {
	Delay := time.Duration(10 * time.Millisecond)

	var i int64
	for i = 1; i <= 100; i++ {
		fmt.Println(RNG.Int64N(i))
		time.Sleep(Delay)
	}
}
