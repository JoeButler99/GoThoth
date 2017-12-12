package Thoth

type RPNAction func(rpnStack *[50]float64, highest *int)

func RPNAdd(rpnStack *[50]float64, highest *int) {
	*highest -= 2
	rpnStack[*highest] = rpnStack[*highest+1] + rpnStack[*highest]
}

func RPNSubtract(rpnStack *[50]float64, highest *int) {
	*highest -= 2
	rpnStack[*highest] = rpnStack[*highest+1] - rpnStack[*highest]
}

func RPNMultiply(rpnStack *[50]float64, highest *int) {
	*highest -= 2
	rpnStack[*highest] = rpnStack[*highest+1] * rpnStack[*highest]
}

func RPNSquare(rpnStack *[50]float64, highest *int) {
	*highest--
	rpnStack[*highest] = rpnStack[*highest] * rpnStack[*highest]
}

func RPNDivide(rpnStack *[50]float64, highest *int) {
	*highest -= 2
	if rpnStack[*highest] == 0.0 {
		rpnStack[*highest] = rpnStack[*highest+1]
	} else {
		rpnStack[*highest] = rpnStack[*highest+1] / rpnStack[*highest]
	}

}

func RPNPercentMe(rpnStack *[50]float64, highest *int) {
	*highest--
	rpnStack[*highest] = rpnStack[*highest] / 100
}

func RPNAvgOf2(rpnStack *[50]float64, highest *int) {
	*highest -= 2
	rpnStack[*highest] = (rpnStack[*highest] + rpnStack[*highest+1]) / 2
}
