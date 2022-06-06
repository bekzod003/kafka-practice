package main

import (
	"fmt"

	"github.com/bekzod003/kafka-practice/config"
)

func main() {
	cfg := config.NewConfig("localhost", "9092", "test")
	fmt.Printf("cfg: %v\n", cfg)
}
