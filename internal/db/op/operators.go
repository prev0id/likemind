package op

type Operator int

const (
	Eq Operator = iota
	Neq
	L
	G
)

func (o Operator) String() string {
	switch o {
	case Eq:
		return "="
	case Neq:
		return "<>"
	case L:
		return "<"
	case G:
		return ">"
	default:
		return "="
	}
}
