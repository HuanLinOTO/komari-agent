package monitoring

import (
	"testing"
)

func TestParseAMDPower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Valid power", "150.5", 150.5},
		{"Power with W suffix", "280.30 W", 280.30},
		{"Power with extra spaces", "  125.7  ", 125.7},
		{"Empty string", "", 0.0},
		{"N/A value", "N/A", 0.0},
		{"Zero power", "0", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseAMDPower(tt.input)
			if err != nil && tt.input != "" && tt.input != "N/A" {
				t.Errorf("parseAMDPower(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseAMDPower(%q) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseAMDPercentage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Valid percentage", "75", 75.0},
		{"Percentage with %", "45 %", 45.0},
		{"Percentage with decimals", "60.5", 60.5},
		{"Empty string", "", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseAMDPercentage(tt.input)
			if err != nil && tt.input != "" {
				t.Errorf("parseAMDPercentage(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseAMDPercentage(%q) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseAMDMemoryBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"Valid bytes", "1073741824", 1073741824},
		{"Zero bytes", "0", 0},
		{"Empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseAMDMemoryBytes(tt.input)
			if err != nil && tt.input != "" {
				t.Errorf("parseAMDMemoryBytes(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseAMDMemoryBytes(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseAMDTemperature(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"Valid temperature", "65", 65},
		{"Temperature with C suffix", "60 C", 60},
		{"Empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseAMDTemperature(tt.input)
			if err != nil && tt.input != "" {
				t.Errorf("parseAMDTemperature(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseAMDTemperature(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
