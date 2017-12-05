package Thoth

type RPNAction func(rpnStack *[50]float64, highest *uint)


func RPNAdd(rpnStack *[50]float64, highest *uint) {
	*highest -= 2
	rpnStack[*highest] = rpnStack[*highest+1] + rpnStack[*highest]
}

func RPNSubtract(rpnStack *[50]float64, highest *uint) {
	*highest -= 2
	rpnStack[*highest] = rpnStack[*highest+1] - rpnStack[*highest]
}

func RPNMultiply(rpnStack *[50]float64, highest *uint) {
	*highest -= 2
	rpnStack[*highest] = rpnStack[*highest+1] * rpnStack[*highest]
}


func RPNSquare(rpnStack *[50]float64, highest *uint) {
	*highest --
	rpnStack[*highest] = rpnStack[*highest] * rpnStack[*highest]
}