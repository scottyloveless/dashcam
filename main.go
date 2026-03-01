package main

import (
	"fmt"
)

type config struct {
	Port string
}

func main() {
	cfg := config{
		Port: "8080",
	}

	fmt.Printf("Port is %s", cfg.Port)
}
