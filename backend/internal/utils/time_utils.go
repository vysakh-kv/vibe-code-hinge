package utils

import (
	"fmt"
	"time"
)

// TimeFormat is the standard time format used in the API
const (
	TimeFormat     = time.RFC3339
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// FormatTime formats a time.Time to a string using the standard format
func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}

// FormatDate formats a time.Time to a date string
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatDateTime formats a time.Time to a date and time string
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// ParseTime parses a string to a time.Time using the standard format
func ParseTime(s string) (time.Time, error) {
	return time.Parse(TimeFormat, s)
}

// ParseDate parses a date string to a time.Time
func ParseDate(s string) (time.Time, error) {
	return time.Parse(DateFormat, s)
}

// ParseDateTime parses a date and time string to a time.Time
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse(DateTimeFormat, s)
}

// IsValidTimeString checks if a string can be parsed as a time
func IsValidTimeString(s string) bool {
	_, err := ParseTime(s)
	return err == nil
}

// IsValidDateString checks if a string can be parsed as a date
func IsValidDateString(s string) bool {
	_, err := ParseDate(s)
	return err == nil
}

// IsValidDateTimeString checks if a string can be parsed as a date and time
func IsValidDateTimeString(s string) bool {
	_, err := ParseDateTime(s)
	return err == nil
}

// GetStartOfDay returns the start of the day for a given time
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetEndOfDay returns the end of the day for a given time
func GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// GetStartOfWeek returns the start of the week (Sunday) for a given time
func GetStartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	return GetStartOfDay(t.AddDate(0, 0, -weekday))
}

// GetEndOfWeek returns the end of the week (Saturday) for a given time
func GetEndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	return GetEndOfDay(t.AddDate(0, 0, 6-weekday))
}

// GetStartOfMonth returns the start of the month for a given time
func GetStartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfMonth returns the end of the month for a given time
func GetEndOfMonth(t time.Time) time.Time {
	return GetEndOfDay(GetStartOfMonth(t).AddDate(0, 1, -1))
}

// GetTimeAgo returns a human-readable time ago string
func GetTimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case diff < 48*time.Hour:
		return "yesterday"
	case diff < 7*24*time.Hour:
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	case diff < 30*24*time.Hour:
		weeks := int(diff.Hours() / 24 / 7)
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	case diff < 365*24*time.Hour:
		months := int(diff.Hours() / 24 / 30)
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	default:
		years := int(diff.Hours() / 24 / 365)
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

// GetHumanDuration returns a human-readable duration
func GetHumanDuration(d time.Duration) string {
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%d seconds", int(d.Seconds()))
	case d < time.Hour:
		return fmt.Sprintf("%d minutes", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%d hours", int(d.Hours()))
	case d < 7*24*time.Hour:
		return fmt.Sprintf("%d days", int(d.Hours()/24))
	case d < 30*24*time.Hour:
		return fmt.Sprintf("%d weeks", int(d.Hours()/24/7))
	case d < 365*24*time.Hour:
		return fmt.Sprintf("%d months", int(d.Hours()/24/30))
	default:
		return fmt.Sprintf("%d years", int(d.Hours()/24/365))
	}
} 