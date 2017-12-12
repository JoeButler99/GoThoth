package Thoth

import (
	"fmt"
	"testing"
)

func assert_stack_element(t *testing.T, rpnStack *[50]float64, element uint, expected float64) {
	if rpnStack[element] != expected {
		fmt.Printf("Expected %f got %f\n", expected, rpnStack[element])
		t.FailNow()
	}
}

func assert_true(t *testing.T, test bool) {
	if test == false {
		t.FailNow()
	}
}

func assert_floats_equal(t *testing.T, a float64, b float64) {
	var EPSILON float64 = 0.00000001
	if (a-b) < EPSILON && (b-a) < EPSILON {

	} else {
		t.Log("Float Comparison failed")
		t.FailNow()
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

func TestRandPostiveInts(t *testing.T) {
	test_pairs := [][2]int{
		{1, 4},
		{0, 2},
		{0, 1},
		{-10, -5},
		{-100, 100},
	}
	rand_tests := 0
	for _, pair := range test_pairs {
		for i := 0; i < 50; i++ {
			result := RandPostiveIntBetween(pair[0], pair[1])
			rand_tests++
			if result < pair[0] || result > pair[1] {
				t.Logf("result %d , min %d, max %d", result, pair[0], pair[1])
				t.Fail()
			}
		}
	}
	t.Logf("Tested %d rands", rand_tests)
}

func TestFitnessCases_LoadFile(t *testing.T) {
	f := FitnessCases{}
	f.LoadFile("fitness_cases/test_sine_x")

	assert_true(t, f.ScalingEnabled == false)
	assert_true(t, f.Terminals == 1)
	assert_true(t, f.TotalCases == 63)
	assert_true(t, f.NumVars == 1)
	assert_true(t, f.NumConsts == 0)
	assert_true(t, f.ConstLower == 0)
	assert_true(t, f.ConstUpper == 0)
	assert_true(t, f.TargetScore == 0)
	assert_true(t, f.ScalingLower == 1)
	assert_true(t, f.ScalingUpper == 1)
	assert_true(t, len(f.Targets) == 63)
	assert_true(t, len(f.Cases) == 63)
	assert_true(t, len(f.Multipliers) == 63)
	assert_true(t, f.Multipliers[0] == 1)
	assert_true(t, f.Multipliers[1] == 1)
	assert_true(t, len(f.Cases[0]) == 1)
	assert_true(t, len(f.Cases[1]) == 1)
}

func TestFunctionSet_GiveRandFunction(t *testing.T) {
	// Just want to test it returns cleanly a few times
	fs := BuildFunctionSet()
	functionCounts := make(map[string]int)
	for i := 0; i < 500; i++ {
		result := fs.GiveRandFunction()
		assert_true(t, result.Inputs == 1 || result.Inputs == 2)
		if _, ok := functionCounts[result.Name]; ok {
			functionCounts[result.Name]++
		} else {
			functionCounts[result.Name] = 1
		}
	}
	assert_true(t, len(functionCounts) == fs.FunctionListSize)
	for _, val := range functionCounts {
		assert_true(t, val > 10)
	}
}

func TestFunctionSet_GiveRandFunctionWithSetInputSize(t *testing.T) {
	fs := BuildFunctionSet()
	functionCounts := make(map[string]int)
	rng := NewBoolgen()
	var inputs int
	for i := 0; i < 500; i++ {
		if rng.RandBool() {
			inputs = 1
		} else {
			inputs = 2
		}
		result := fs.GiveRandFunctionWithSetInputSize(inputs)
		assert_true(t, result.Inputs == uint(inputs))
		if _, ok := functionCounts[result.Name]; ok {
			functionCounts[result.Name]++
		} else {
			functionCounts[result.Name] = 1
		}
	}
	assert_true(t, len(functionCounts) == fs.FunctionListSize)
	for _, val := range functionCounts {
		assert_true(t, val > 10)
	}
}

func setupNodeToTestAction(t *testing.T, action string, p *PopulationMember, fs *FunctionSet) {
	t.Logf("Setup %s", action)
	p.Nodes = []Node{}
	p.Nodes = append(p.Nodes, Node{IsTerminal: false, Function: fs.GiveFunctionByName(action)})
	assert_true(t, len(p.Nodes) == 1)
	for i := 0; i < int(fs.GiveFunctionByName(action).Inputs); i++ {
		p.Nodes = append(p.Nodes, Node{TerminalNo: uint(i), IsTerminal: true})
	}
}

func TestPopulationMember_SolveSelf(t *testing.T) {
	test_vars := []float64{1, 2, 3, 4, 5}
	fs := BuildFunctionSet()
	p := PopulationMember{}

	action := "Add"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_true(t, p.SolveSelf(&test_vars) == 3)

	action = "Subtract"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_true(t, p.SolveSelf(&test_vars) == -1)

	action = "Multiply"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_true(t, p.SolveSelf(&test_vars) == 2)

	action = "Divide"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_floats_equal(t, p.SolveSelf(&test_vars), 0.5)

	action = "Square"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_true(t, p.SolveSelf(&test_vars) == 1)

	action = "PercentMe"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_true(t, p.SolveSelf(&test_vars) == 0.01)

	action = "AVG2"
	setupNodeToTestAction(t, action, &p, &fs)
	t.Logf("%s Result was %f\n", p.Nodes[0].Function.Name, p.SolveSelf(&test_vars))
	assert_floats_equal(t, p.SolveSelf(&test_vars), 1.5)

}

func TestPopulationMember_FillRandomNodes(t *testing.T) {
	fs := BuildFunctionSet()
	p := PopulationMember{}
	fc := FitnessCases{}
	fc.LoadFile("fitness_cases/test_sine_x")

	p.FillRandomNodes(3, &fc, &fs)
	assert_true(t, len(p.Nodes) >= 3)

}
