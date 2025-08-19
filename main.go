// Reference implementation of the Absolute Clock:
// A clock that represents a single moment in time as a unique string.
// This implementation creates a hierarchical time representation that includes
// geological time periods followed by precise human time measurements.
package main

import (
	"absclock/absclock"
	"fmt"
)

func main() {
	clock := absclock.AbsClock{}

	// Get the current absolute time
	fmt.Println("  Current absolute time: ", clock.GetCurrentAbsoluteTime())

	// Convert an ISO 8601 string (UTC) to an absolute time string
	iso8601 := "1776-07-04T12:00:00Z"

	absTime, err := clock.ConvertISO8601ToAbsoluteTime(iso8601)
	if err != nil {
		fmt.Printf("Error converting ISO 8601: %v\n", err)
		return
	}
	fmt.Println("Converted from ISO 8601: ", absTime)
}
