// Go 1.25.1: X:nowarf5 linux/amd64
package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Money int64 //`yaml:"Money"`
	Debt  int64 //`yaml:"Debt"`
}

type Core struct {
	User User //`yaml:"User"`
} 
// Pro-Tip: String Literals aren't required
// Lowercase the keys in the YAML file

func main() {
	x, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var CFG Core
	if err := yaml.Unmarshal(x, &CFG); err != nil {
		log.Fatal(err)
	}
	fmt.Println(CFG)
}
