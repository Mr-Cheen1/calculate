package main

import (
	"testing"
)

const (
	timeZero     = "00:00:00"
	timeSixHours = "06:00:00"
	timeOneHour  = "01:00:00"
)

func TestParseTimeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Valid time", "01:30:45", 5445},
		{"Zero time", "00:00:00", 0},
		{"Large time", "47:59:58", 172798},
		{"Invalid format", "01:30", -1},
		{"Invalid minutes", "01:60:00", -1},
		{"Invalid seconds", "01:30:60", -1},
		{"Non-numeric", "aa:bb:cc", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseTimeString(tt.input)
			if result != tt.expected {
				t.Errorf("parseTimeString(%s) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFormatSeconds(t *testing.T) {
	tests := []struct {
		name     string
		seconds  int
		expected string
	}{
		{"Zero", 0, timeZero},
		{"One hour", 3600, timeOneHour},
		{"Six hours", 21600, timeSixHours},
		{"Complex time", 5445, "01:30:45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSeconds(tt.seconds)
			if result != tt.expected {
				t.Errorf("formatSeconds(%d) = %s, expected %s", tt.seconds, result, tt.expected)
			}
		})
	}
}

func TestParseAndCalculateTime_Basic(t *testing.T) {
	app := &App{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Single time", "01:30:45", "01:30:45"},
		{"Multiple times", "01:00:00\n02:00:00\n03:00:00", timeSixHours},
		{"Empty input", "", timeZero},
		{"Invalid input", "invalid\ngarbage", timeZero},
		{"Mixed valid/invalid", "01:00:00\ninvalid\n02:00:00", "03:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := app.ParseAndCalculateTime(tt.input)
			if result.TotalFormatted != tt.expected {
				t.Errorf("ParseAndCalculateTime(%s).TotalFormatted = %s, expected %s",
					tt.input, result.TotalFormatted, tt.expected)
			}
		})
	}
}
