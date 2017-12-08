package Thoth

type PopulationMember struct {
	CreationMethod string
	HasChanged     bool
	Score          float64
	Nodes          []Node
}
