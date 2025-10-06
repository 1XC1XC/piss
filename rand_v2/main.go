package main

import (
	"fmt"
	"math"
	"time"

	"math/rand/v2"
)

var RNG *rand.Rand

func init() {
	var Unix uint64 = uint64(time.Now().UnixMicro())
	RandUnix := func() uint64 { return Unix - rand.Uint64N(Unix) }
	RNG = rand.New(rand.NewPCG(RandUnix(), RandUnix()))

}

func o(k, j any) time.Duration {
	x := math.Pow(k.(float64), j.(float64))
	return time.Duration(x)
}

func main() {
	var i int64
	for i = 1; i <= 100; i++ {
		Delay := time.Millisecond * o(i, float64(i))
		fmt.Println(i, Delay)
		time.Sleep(Delay)
	}
}
