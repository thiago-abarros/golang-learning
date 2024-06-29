package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("uuid: ", uuid.New().String())
	fmt.Println("Hello, World!")
}
