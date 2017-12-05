package Thoth

//typedef void (*rpn_action_ptr)(double rpnStack[RPN_STACK_LIMIT],int & highest);

type Function struct {
	Name string
	Inputs uint
	RPNAction RPNAction
}

