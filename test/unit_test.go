package Test

import (
	"testing"
	"Thoth"
)


func TestRPNAdd(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 2
	stack[1] = 7
	var highest uint = 2
	Thoth.RPNAdd(&stack, &highest)
	if stack[0] != 9 {
		t.Fail()
	}
}

func TestRPNSubtract(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 2
	stack[1] = 7
	var highest uint = 2
	Thoth.RPNSubtract(&stack, &highest)
	if stack[0] != 5 {
		t.Fail()
	}

}

func TestRPNMultiply(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 2
	stack[1] = 7
	var highest uint = 2
	Thoth.RPNMultiply(&stack, &highest)
	if stack[0] != 14 {
		t.Fail()
	}

}

func TestRPNSquare(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 4
	var highest uint = 1
	Thoth.RPNSquare(&stack, &highest)
	if stack[0] != 16 {
		t.Fail()
	}
}