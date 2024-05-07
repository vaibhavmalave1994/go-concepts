package main

import (
	"fmt"
	"time"
)

func main() {
	target := []string{"target 1", "target 2", "target 3", "target 4"}
	for _, value := range target {
		go attack(value)
	}
	time.Sleep(2 * time.Second)
}

func attack(target string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Attacked on:", target)
}
