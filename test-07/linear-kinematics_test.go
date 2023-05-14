package main

import (
	"bufio"
	"testing"
	"testing/fstest"
)

func TestGenDisplaceFn(t *testing.T) {
	tests := []struct {
		name             string
		acceleration     float64
		initVelocity     float64
		initDisplacement float64
		time             float64
		expected         float64
	}{
		{"zero values", 0, 0, 0, 0, 0},
		{"with one element", 10, 2, 1, 3, 52},
		{"with one element", 10, 2, 1, 5, 136},
	}

	for _, test := range tests {
		actual := GenDisplaceFn(test.acceleration, test.initVelocity, test.initDisplacement)(test.time)
		if actual != test.expected {
			t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

func TestPromptFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected float64
	}{
		{"float number", []byte("1.2\n2.4\n"), 1.2},
		{"decimal komma", []byte("2,6\n2300\n"), 2300},
		{"invalid syntax", []byte("keene_zahl_first\n2\n"), 2},
	}

	for _, test := range tests {
		fsMap := fstest.MapFS{
			"stdin": &fstest.MapFile{Data: test.input},
		}
		in, _ := fsMap.Open("stdin")
		actual := promptFloat64(bufio.NewReader(in))
		in.Close()
		if actual != test.expected {
			t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, actual)
		}
	}
}
