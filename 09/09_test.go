package main

import "testing"

func _Test(file string, knots int, expected int, t *testing.T) {
	result := run(file, knots)
	if result != expected {
		t.Errorf("expected %d but got %d", expected, result)
	}
}

func TestDemo(t *testing.T) {
	_Test("demo1.txt", 2, 13, t)
	_Test("demo2.txt", 10, 36, t)
}
