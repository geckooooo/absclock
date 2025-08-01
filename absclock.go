// Reference implementation of the Absolute Clock:
// A clock that represents a single moment in time as a unique string.
// This implementation creates a hierarchical time representation that includes
// geological time periods followed by precise human time measurements.
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	currAbsoluteTime := getCurrentAbsoluteTime()
	fmt.Println(currAbsoluteTime)
}

// Builds the absolute time string by extracting the current time
// and formatting it according to the Absolute Clock specification.
// This function duplicates the logic in main() for demonstration purposes.
func getCurrentAbsoluteTime() string {
	// Get the current time broken into its components.
	// Time is represented in UTC to ensure consistency across time zones.
	now := time.Now().UTC()

	// Extract individual time components and convert them to strings
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	millisecond := strconv.Itoa(now.Nanosecond() / 1000000) // Convert nanoseconds to milliseconds
	nanosecond := strconv.Itoa(now.Nanosecond() / 1000)     // Get microseconds within the current millisecond

	// Pad the values with a leading zero if needed to ensure consistent formatting.
	// This ensures all time components have at least 2 digits.
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

	// Pad milliseconds and nanoseconds for consistent formatting
	if (now.Nanosecond() / 1000000) < 100 {
		millisecond = "0" + millisecond
	}

	if (now.Nanosecond() / 1000) < 100000 {
		nanosecond = "0" + nanosecond
	}

	// Build the string using a string builder for efficiency
	var builder strings.Builder

	// Add geological time hierarchy - these values represent our current position in Earth's history
	builder.WriteString("E:") // Eternity - represents the entire existence of time
	builder.WriteString("4:") // Phanerozoic eon - the current eon (last 541 million years)
	builder.WriteString("3:") // Cenozoic era - the current era (last 66 million years)
	builder.WriteString("3:") // Quaternary period - the current period (last 2.6 million years)
	builder.WriteString("2:") // Holocene epoch - the current epoch (last 11,700 years)
	builder.WriteString("3:") // Meghalayan age - the current age (last 4,200 years)

	// Add precise human time measurements
	builder.WriteString(year) // Year (e.g., 2024)
	builder.WriteString(":")
	builder.WriteString(month) // Month (01-12)
	builder.WriteString(":")
	builder.WriteString(day) // Day (01-31)
	builder.WriteString(":")
	builder.WriteString(hour) // Hour (00-23)
	builder.WriteString(":")
	builder.WriteString(minute) // Minute (00-59)
	builder.WriteString(":")
	builder.WriteString(second) // Second (00-59)
	builder.WriteString(":")
	builder.WriteString(millisecond) // Millisecond (000-999)
	builder.WriteString(":")
	builder.WriteString(nanosecond) // Microsecond within millisecond (000000-999999)

	// Return the complete absolute time string
	return builder.String()
}
