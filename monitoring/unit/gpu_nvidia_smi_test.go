package monitoring

import (
	"testing"
)

func TestParsePowerValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Valid power", "350.50 W", 350.50},
		{"Power without W", "280.30", 280.30},
		{"Power with extra spaces", "  150.5 W  ", 150.5},
		{"Empty string", "", 0.0},
		{"N/A value", "N/A", 0.0},
		{"Zero power", "0 W", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parsePowerValue(tt.input)
			if err != nil && tt.input != "" && tt.input != "N/A" {
				t.Errorf("parsePowerValue(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parsePowerValue(%q) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParsePercentageValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Valid percentage", "75 %", 75.0},
		{"Percentage without %", "45", 45.0},
		{"Percentage with decimals", "60.5 %", 60.5},
		{"Empty string", "", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parsePercentageValue(tt.input)
			if err != nil && tt.input != "" {
				t.Errorf("parsePercentageValue(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parsePercentageValue(%q) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseMemoryValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"Valid memory in MiB", "24564 MiB", 24564 * 1024 * 1024},
		{"Memory without MiB", "12282", 12282 * 1024 * 1024},
		{"Empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseMemoryValue(tt.input)
			if err != nil && tt.input != "" {
				t.Errorf("parseMemoryValue(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseMemoryValue(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseTemperatureValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"Valid temperature", "65 C", 65},
		{"Temperature without C", "60", 60},
		{"Empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTemperatureValue(tt.input)
			if err != nil && tt.input != "" {
				t.Errorf("parseTemperatureValue(%q) returned error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseTemperatureValue(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
