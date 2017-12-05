package Test

import (
	"Thoth"
	"fmt"
	"testing"
)

func assert_stack_element(t *testing.T, rpnStack *[50]float64, element uint, expected float64) {
	if rpnStack[element] != expected {
		fmt.Printf("Expected %f got %f\n", expected, rpnStack[element])
		t.Fail()
	}
}

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
	assert_stack_element(t, &stack, 0, 5)

}

func TestRPNMultiply(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 2
	stack[1] = 7
	var highest uint = 2
	Thoth.RPNMultiply(&stack, &highest)
	assert_stack_element(t, &stack, 0, 14)

}

func TestRPNSquare(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 4
	var highest uint = 1
	Thoth.RPNSquare(&stack, &highest)
	assert_stack_element(t, &stack, 0, 16)
}

func TestPRNDivide(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 4
	stack[1] = 8
	var highest uint = 2
	Thoth.RPNDivide(&stack, &highest)
	assert_stack_element(t, &stack, 0, 2)
}

func TestRPNPercentMe(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 50
	var highest uint = 1
	Thoth.RPNPercentMe(&stack, &highest)
	assert_stack_element(t, &stack, 0, 0.5)
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

	assert_stack_element(t, &stack, 0, 3)

	f.RPNAction = Thoth.RPNSquare
	highest = 1
	f.RPNAction(&stack, &highest)
	assert_stack_element(t, &stack, 0, 9)
}
