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
