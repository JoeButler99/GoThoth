package Thoth

type Node struct {
	Function   *Function
	TerminalNo uint
	IsTerminal bool
}

type PopulationMember struct {
	CreationMethod string
	HasChanged     bool
	Score          float64
	Nodes          []Node
}

func (p *PopulationMember) SolveSelf(var_set *[]float64) float64 {
	var rpnStack = [50]float64{}
	highest := uint(0)
	vs := *var_set

	for x := len(p.Nodes) - 1; x > -1; x-- {
		if p.Nodes[x].IsTerminal {
			// Add to stack
			rpnStack[highest] = vs[p.Nodes[x].TerminalNo]
			highest++
		} else {
			p.Nodes[x].Function.RPNAction(&rpnStack, &highest)
			highest++
		}
	}
	return rpnStack[0]
}

func (p *PopulationMember) FillRandomNodes(maxDepth int, fitnessCases *FitnessCases, functionSet *FunctionSet) {
	if maxDepth == 0 || (p.CreationMethod == "grow" && RandPostiveIntUpTo(20)%5 == 0) {
		p.Nodes = append(p.Nodes, Node{IsTerminal: true, TerminalNo: uint(RandPostiveIntUpTo(int(fitnessCases.Terminals)))})
	} else {
		n := Node{Function: functionSet.GiveRandFunction()}
		p.Nodes = append(p.Nodes, n)
		for x := 0; x < int(n.Function.Inputs); x++ {
			p.FillRandomNodes(maxDepth-1, fitnessCases, functionSet)
		}
	}
}

// TODO - This function
//func (p *PopulationMember) DisplaySelf() {
//
//}
