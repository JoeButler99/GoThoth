package Test

import (
	"Thoth"
	"fmt"
	"testing"
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

func TestPRNDivide(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 4
	stack[1] = 8
	var highest uint = 2
	Thoth.RPNDivide(&stack, &highest)
	if stack[0] != 2 {
		fmt.Println("stack 0 is %d", stack[0])
		t.Fail()

	}
}

func TestRPNPercentMe(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 50
	var highest uint = 1
	Thoth.RPNPercentMe(&stack, &highest)
	if stack[0] != 0.5 {
		t.Fail()
	}
}

func TestFunction(t *testing.T) {
	f := Thoth.Function{
		Name:      "Fadd",
		Inputs:    2,
		RPNAction: Thoth.RPNAdd,
	}

	var stack = [50]float64{}
	var highest uint = 2
	stack[0] = 1
	stack[1] = 2

	f.RPNAction(&stack, &highest)

	if stack[0] != 3 {
		t.Fail()
	}

	f.RPNAction = Thoth.RPNSquare
	highest = 1
	f.RPNAction(&stack, &highest)
	if stack[0] != 9 {
		t.Fail()
	}
}
