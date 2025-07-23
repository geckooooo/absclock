package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var builder strings.Builder
	builder.WriteString("Current time: ")
	builder.WriteString(time.Now().String())

	result := builder.String()

	fmt.Println(result)
}
