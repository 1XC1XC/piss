package main

import (
	"fmt"
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
	fmt.Println(RNG.Int64N(100))
}
