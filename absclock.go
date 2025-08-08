// Reference implementation of the Absolute Clock:
// A clock that represents a single moment in time as a unique string.
// This implementation creates a hierarchical time representation that includes
// geological time periods followed by precise human time measurements.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current absolute time
	currAbsoluteTime := getCurrentAbsoluteTime()
	fmt.Println("           Current time: ", currAbsoluteTime)

	// Convert an ISO 8601 string (UTC) to an absolute time string
	iso8601 := "1776-07-04T12:00:00Z"
	absTime, err := convertISO8601ToAbsoluteTime(iso8601)
	if err != nil {
		fmt.Printf("Error converting ISO 8601: %v\n", err)
		return
	}
	fmt.Println("Converted from ISO 8601: ", absTime)
}

const (
	eternity         = "E" // Eternity - represents the entire existence of time
	phanerozoicEon   = "4" // Phanerozoic eon - the current eon (last 541 million years)
	cenozoicEra      = "3" // Cenozoic era - the current era (last 66 million years)
	quaternaryPeriod = "3" // Quaternary period - the current period (last 2.6 million years)
	holoceneEpoch    = "2" // Holocene epoch - the current epoch (last 11,700 years)
	meghalayanAge    = "3" // Meghalayan age - the current age (last 4,200 years)
)

// getCurrentAbsoluteTime builds the absolute time string for the current moment
// by formatting it according to the Absolute Clock specification.
func getCurrentAbsoluteTime() string {
	return formatTimeToAbsolute(time.Now().UTC())
}

// convertISO8601ToAbsoluteTime converts an ISO 8601 string to an absolute time string
func convertISO8601ToAbsoluteTime(iso8601 string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, iso8601)
	if err != nil {
		return "", fmt.Errorf("failed to parse ISO 8601 time: %w", err)
	}

	return formatTimeToAbsolute(parsedTime.UTC()), nil
}

/*
// Converts an ISO 8601 string to an absolute time string
func convertISO8601ToAbsoluteTime(iso8601 string) string {
	// Parse the ISO 8601 string into a time.Time object
	parsedTime, err := time.Parse(time.RFC3339, iso8601)

	if err != nil {
		return ""
	}

	// Convert the time.Time object to an absolute time string
	// Note: truncated to the second
	var absTime = "E:4:3:3:2:3:" + parsedTime.Format("2006:01:02:15:04:05")

	return absTime
}*/

// formatTimeToAbsolute converts a time.Time to the absolute time string format
func formatTimeToAbsolute(t time.Time) string {
	// Extract time components
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	nanosecond := t.Nanosecond()

	// Calculate milliseconds and microseconds within the current second
	millisecond := nanosecond / 1000000
	microsecond := (nanosecond / 1000) % 1000000 // microseconds within the current second

	// Build the absolute time string using fmt.Sprintf for clean formatting
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%04d:%02d:%02d:%02d:%02d:%02d:%03d:%06d",
		eternity,
		phanerozoicEon,
		cenozoicEra,
		quaternaryPeriod,
		holoceneEpoch,
		meghalayanAge,
		year,
		month,
		day,
		hour,
		minute,
		second,
		millisecond,
		microsecond,
	)
}
