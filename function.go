package Thoth

//typedef void (*rpn_action_ptr)(double rpnStack[RPN_STACK_LIMIT],int & highest);

type Function struct {
	Name      string
	Inputs    uint
	RPNAction RPNAction
}

type FunctionSet struct {
	Name                                                                       string
	SingleInputFunctionList, DoubleInputFunctionList, FunctionList             []Function
	SingleInputFunctionListSize, DoubleInputFunctionListSize, FunctionListSize int
}

func BuildFunctionSet() FunctionSet {
	fs := FunctionSet{
		Name: "Basic",

		// TODO - Invesigate if there is a dynamic way to generate
		// TODO   The later lists and ints
		FunctionList: []Function{
			{Name: "Add", Inputs: 2, RPNAction: RPNAdd},
			{Name: "Subtract", Inputs: 2, RPNAction: RPNSubtract},
			{Name: "Multiply", Inputs: 2, RPNAction: RPNMultiply},
			{Name: "Divide", Inputs: 2, RPNAction: RPNDivide},
			{Name: "Square", Inputs: 1, RPNAction: RPNSquare},
			{Name: "PercentMe", Inputs: 1, RPNAction: RPNPercentMe},
			{Name: "AVG2", Inputs: 2, RPNAction: RPNAvgOf2},
		},
	}

	for _, f := range fs.FunctionList {
		if f.Inputs == 1 {
			fs.SingleInputFunctionList = append(fs.SingleInputFunctionList, f)
		} else if f.Inputs == 2 {
			fs.DoubleInputFunctionList = append(fs.DoubleInputFunctionList, f)
		} else {
			panic("Function with wrong number of inputs")
		}
	}

	fs.FunctionListSize = len(fs.FunctionList)
	fs.DoubleInputFunctionListSize = len(fs.DoubleInputFunctionList)
	fs.SingleInputFunctionListSize = len(fs.SingleInputFunctionList)
	return fs
}

func (f *FunctionSet) GiveRandFunction() Function {
	return f.FunctionList[RandPostiveIntUpTo(f.FunctionListSize)]
}

func (f *FunctionSet) GiveRandFunctionWithSetInputSize(inputs int) Function {
	if inputs == 1 {
		return f.SingleInputFunctionList[RandPostiveIntUpTo(f.SingleInputFunctionListSize)]
	} else if inputs == 2 {
		return f.DoubleInputFunctionList[RandPostiveIntUpTo(f.DoubleInputFunctionListSize)]
	} else {
		panic("Request for function with wrong number of inputs")
	}
}

func (f *FunctionSet) GiveFunctionByName(name string) *Function {
	// Initiall created for testing.
	for _, function := range f.FunctionList {
		if function.Name == name {
			return &function
		}
	}
	panic("Could not find function")
}
