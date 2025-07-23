package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get the current time broken into its components.
	// Time is represented in Zulu time.
	now := time.Now().UTC()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())

	// Pad the values with a leading zero if tanyhey are less than 10.
	if now.Month() < 10 {
		month = "0" + month
	}
	if now.Day() < 10 {
		day = "0" + day
	}
	if now.Hour() < 10 {
		hour = "0" + hour
	}
	if now.Minute() < 10 {
		minute = "0" + minute
	}
	if now.Second() < 10 {
		second = "0" + second
	}

	// Build the string
	var builder strings.Builder
	builder.WriteString("E:")
	builder.WriteString("P:")
	builder.WriteString(year)
	builder.WriteString(":")
	builder.WriteString(month)
	builder.WriteString(":")
	builder.WriteString(day)
	builder.WriteString(":")
	builder.WriteString(hour)
	builder.WriteString(":")
	builder.WriteString(minute)
	builder.WriteString(":")
	builder.WriteString(second)

	absTime := builder.String()

	fmt.Println(absTime)
}
