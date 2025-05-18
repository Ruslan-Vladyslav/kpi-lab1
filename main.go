package main

import (
	"fmt"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

func main() {
	fmt.Println(getCurrentTime())
}
