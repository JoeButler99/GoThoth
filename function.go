package Thoth

//typedef void (*rpn_action_ptr)(double rpnStack[RPN_STACK_LIMIT],int & highest);

type Function struct {
	Name      string
	Inputs    uint
	RPNAction RPNAction
}

type FunctionSet struct {
	Name         string
	FunctionList []Function
}

func BuildFunctionSet() FunctionSet {
	fs := FunctionSet{
		Name: "Basic",
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

	return fs
}
