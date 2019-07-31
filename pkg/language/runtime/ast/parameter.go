package ast

type Parameter struct {
	Label      string
	Identifier string
	Type       Type
	StartPos   *Position
	EndPos     *Position
}

func (p *Parameter) StartPosition() *Position {
	return p.StartPos
}

func (p *Parameter) EndPosition() *Position {
	return p.EndPos
}
