package Thoth

import (
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
	RPNAdd(&stack, &highest)
	if stack[0] != 9 {
		t.Fail()
	}
}

func TestRPNSubtract(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 2
	stack[1] = 7
	var highest uint = 2
	RPNSubtract(&stack, &highest)
	assert_stack_element(t, &stack, 0, 5)

}

func TestRPNMultiply(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 2
	stack[1] = 7
	var highest uint = 2
	RPNMultiply(&stack, &highest)
	assert_stack_element(t, &stack, 0, 14)

}

func TestRPNSquare(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 4
	var highest uint = 1
	RPNSquare(&stack, &highest)
	assert_stack_element(t, &stack, 0, 16)
}

func TestPRNDivide(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 4
	stack[1] = 8
	var highest uint = 2
	RPNDivide(&stack, &highest)
	assert_stack_element(t, &stack, 0, 2)

	stack[0] = 0
	stack[1] = 8
	highest = 2
	RPNDivide(&stack, &highest)
	assert_stack_element(t, &stack, 0, 8)

}

func TestRPNPercentMe(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 50
	var highest uint = 1
	RPNPercentMe(&stack, &highest)
	assert_stack_element(t, &stack, 0, 0.5)
}

func TestRPNAvgOf2(t *testing.T) {
	var stack = [50]float64{}
	stack[0] = 50
	stack[1] = 10
	var highest uint = 2
	RPNAvgOf2(&stack, &highest)
	assert_stack_element(t, &stack, 0, 30)

}

func TestFunction(t *testing.T) {
	f := Function{
		Name:      "Fadd",
		Inputs:    2,
		RPNAction: RPNAdd,
	}

	var stack = [50]float64{}
	var highest uint = 2
	stack[0] = 1
	stack[1] = 2
	f.RPNAction(&stack, &highest)
	assert_stack_element(t, &stack, 0, 3)

	f.RPNAction = RPNSquare
	highest = 1
	f.RPNAction(&stack, &highest)
	assert_stack_element(t, &stack, 0, 9)
}

func TestFunctionSets(t *testing.T) {
	fs := BuildFunctionSet()

	if fs.FunctionList[0].Name != "Add" || fs.FunctionList[1].Name != "Subtract" {
		t.Fail()
	}
}

func TestRandBool(t *testing.T) {
	ts, fs := 0, 0
	r := NewBoolgen()

	for i := 0; i < 100; i++ {
		if r.RandBool() {
			ts++
		} else {
			fs++
		}
	}
	if ts < 10 || fs < 10 {
		fmt.Println("Random Boolean generation failed")
		t.Fail()
	} else {
		t.Logf("Generated %d true and %d false", ts, fs)
	}
}

func TestRandFloats(t *testing.T) {
	test_pairs := [][2]float64{
		{-1, 1},
		{0, 3},
		{-10, -5},
		{4, 5},
		{1000, 2000},
		{0, 0.1},
		{-0.1, 0},
	}
	rand_tests := 0
	for _, pair := range test_pairs {
		for i := 0; i < 150; i++ {
			result := RandFloatBetween(pair[0], pair[1])
			rand_tests++
			if result < pair[0] || result > pair[1] {
				t.Logf("result %f , min %f, max %f", result, pair[0], pair[1])
				t.Fail()
			}
		}
	}
	t.Logf("Tested %d rands", rand_tests)
}
