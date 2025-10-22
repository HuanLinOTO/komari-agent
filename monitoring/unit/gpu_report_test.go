package monitoring

import (
	"testing"
)

func TestDetailedGPUInfoStructure(t *testing.T) {
	// Test that DetailedGPUInfo has all required fields
	info := DetailedGPUInfo{
		Name:         "NVIDIA GeForce RTX 4090",
		MemoryTotal:  25769803776,
		MemoryUsed:   12884901888,
		Utilization:  75.5,
		Temperature:  65,
		PowerUsage:   350.5,
	}

	// Verify all fields are set
	if info.Name == "" {
		t.Error("Name should not be empty")
	}
	if info.MemoryTotal == 0 {
		t.Error("MemoryTotal should not be zero")
	}
	if info.MemoryUsed == 0 {
		t.Error("MemoryUsed should not be zero")
	}
	if info.Utilization == 0 {
		t.Error("Utilization should not be zero")
	}
	if info.Temperature == 0 {
		t.Error("Temperature should not be zero")
	}
	if info.PowerUsage == 0 {
		t.Error("PowerUsage should not be zero")
	}

	t.Logf("GPU Info: %+v", info)
}

func TestNVIDIAGPUInfoStructure(t *testing.T) {
	// Test that NVIDIAGPUInfo has all required fields including PowerUsage
	info := NVIDIAGPUInfo{
		Name:        "NVIDIA GeForce RTX 4090",
		MemoryTotal: 25769803776,
		MemoryUsed:  12884901888,
		Utilization: 75.5,
		Temperature: 65,
		PowerUsage:  350.5,
	}

	if info.PowerUsage == 0 {
		t.Error("PowerUsage should not be zero")
	}

	t.Logf("NVIDIA GPU Info: %+v", info)
}

func TestAMDGPUInfoStructure(t *testing.T) {
	// Test that AMDGPUInfo has all required fields including PowerUsage
	info := AMDGPUInfo{
		Name:        "AMD Radeon RX 7900 XTX",
		MemoryTotal: 25769803776,
		MemoryUsed:  12884901888,
		Utilization: 60.0,
		Temperature: 70,
		PowerUsage:  300.0,
	}

	if info.PowerUsage == 0 {
		t.Error("PowerUsage should not be zero")
	}

	t.Logf("AMD GPU Info: %+v", info)
}
