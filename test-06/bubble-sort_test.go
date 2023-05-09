package main

import (
	"bufio"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"with one element", []int{1}, []int{1}},
		{"with two elements", []int{2, 1}, []int{1, 2}},
		{"with three elements", []int{3, 2, 1}, []int{1, 2, 3}},
		{"with signed element", []int{3, -2, 1}, []int{-2, 1, 3}},
		{"with zero", []int{3, 0, 1}, []int{0, 1, 3}},
	}

	for _, test := range tests {
		BubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, test.input)
		}
	}
}

func TestInput(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []int
	}{
		{"empty", []byte{}, []int{}},
		{"with one element", []byte("1\n\n"), []int{1}},
		{"with two elements", []byte("2\n1\n\n"), []int{2, 1}},
		{"with three elements", []byte("3\n2\n1\n\n"), []int{3, 2, 1}},
		{"with signed element", []byte("3\n-2\n1\n\n"), []int{3, -2, 1}},
		{"with early break out", []byte("3\n2\n\n1\n\n"), []int{3, 2}},
	}

	for _, test := range tests {
		fsMap := fstest.MapFS{
			"stdin": &fstest.MapFile{Data: test.input},
		}
		in, _ := fsMap.Open("stdin")
		actual := input(bufio.NewReader(in))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, actual)
		}
	}
}
