package utils

import (
	"fmt"
	"time"
)

// RFC3339Time is a custom type for handling RFC 3339 timestamps
type RFC3339Time time.Time

// MarshalJSON converts RFC3339Time to JSON in RFC 3339 format
func (t RFC3339Time) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format(time.RFC3339Nano)
	return []byte(fmt.Sprintf(`"%s"`, formatted)), nil
}

// UnmarshalJSON parses RFC 3339 timestamps from JSON
func (t *RFC3339Time) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02T15:04:05.999999999Z"`, string(data))
	if err != nil {
		return err
	}
	*t = RFC3339Time(parsedTime)
	return nil
}

// String returns the string representation in RFC 3339 format
func (t RFC3339Time) String() string {
	return time.Time(t).Format(time.RFC3339Nano)
}

// TimeNow returns the current time as RFC3339Time
func TimeNow() RFC3339Time {
	return RFC3339Time(time.Now().UTC())
}

// ToTime converts RFC3339Time to standard time.Time
func (t RFC3339Time) ToTime() time.Time {
	return time.Time(t)
}

// ParseRFC3339 parses an RFC 3339 timestamp string
func ParseRFC3339(input string) (RFC3339Time, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		return RFC3339Time{}, err
	}
	return RFC3339Time(parsedTime), nil
}

// ParseFromString sets the RFC3339Time from a string
func (t *RFC3339Time) ParseFromString(input string) error {
	parsedTime, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		return err
	}
	*t = RFC3339Time(parsedTime)
	return nil
}

// IsBefore checks if one RFC3339Time is before another
func (t RFC3339Time) IsBefore(other RFC3339Time) bool {
	return time.Time(t).Before(time.Time(other))
}

// IsAfter checks if one RFC3339Time is after another
func (t RFC3339Time) IsAfter(other RFC3339Time) bool {
	return time.Time(t).After(time.Time(other))
}
