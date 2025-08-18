// Reference implementation of the Absolute Clock:
// A clock that represents a single moment in time as a unique string.
// This implementation creates a hierarchical time representation that includes
// geological time periods followed by precise human time measurements.
package absclock

import (
	"fmt"
	"time"
)

const (
	eternity      = "E" // Eternity - represents the entire existence of time
	currentEon    = "4" // Phanerozoic eon - the current eon (last 541 million years)
	currentEra    = "3" // Cenozoic era - the current era (last 66 million years)
	currentPeriod = "3" // Quaternary period - the current period (last 2.6 million years)
	currentEpoch  = "2" // Holocene epoch - the current epoch (last 11,700 years)
)

type AbsClock struct {
}

// builds the absolute time string for the current moment
// by formatting it according to the Absolute Clock specification.
func (clock *AbsClock) GetCurrentAbsoluteTime() string {
	return clock.FormatTimeToAbsolute(time.Now().UTC())
}

// Finds the age for a given year.
// Expects BCE dates to be expressed as negative value (e.g. -5000 = 5000 BCE).
// Incomplete implementation; only handles years in the current epoch (Holocene).
func (clock *AbsClock) GetAgeForYear(year int) string {
	if year >= -2200 {
		return "3" // Meghalayan Age (2200 BCE to today)
	} else if year >= -6200 {
		return "2" // Northgrippian Age (6200 to 2200 BCE)
	} else if year >= -9700 {
		return "1" // Greenlandian Age (9700 to 6200 BCE)
	}
	return "." // Placeholder--outside Holocene epoch
}

// Converts an ISO 8601 string to an absolute time string
// Note: the time package doesn't allow for dates before 1 CE.
func (clock *AbsClock) ConvertISO8601ToAbsoluteTime(iso8601 string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, iso8601)
	if err != nil {
		return "", fmt.Errorf("failed to parse ISO 8601 time: %w", err)
	}

	return clock.FormatTimeToAbsolute(parsedTime.UTC()), nil
}

// Converts a time.Time to the absolute time string format
func (clock *AbsClock) FormatTimeToAbsolute(t time.Time) string {
	// Extract time components
	age := clock.GetAgeForYear(t.Year()) // Age is inferred from the year
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	nanosecond := t.Nanosecond()

	// Calculate milliseconds and microseconds within the current second
	// Note: this needs conceptual clarification
	millisecond := nanosecond / 1000000
	microsecond := (nanosecond / 1000) % 1000000 // microseconds within the current second

	// Build the absolute time string using fmt.Sprintf for clean formatting
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%04d:%02d:%02d:%02d:%02d:%02d:%03d:%06d",
		eternity,
		currentEon,
		currentEra,
		currentPeriod,
		currentEpoch,
		age,
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
