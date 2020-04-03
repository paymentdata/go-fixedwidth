package fixedwidth

const (
	DefaultAlignment Alignment = "default"
	Right            Alignment = "right"
	Left             Alignment = "left"
)

const (
	DefaultPadChar = ' '
)

var DefaultFormat = Format{
	Alignment: DefaultAlignment,
	PadChar:   DefaultPadChar,
}

type Format struct {
	Alignment Alignment
	PadChar   byte
}

type Alignment string

func (a Alignment) Valid() bool {
	switch a {
	case DefaultAlignment, Right, Left:
		return true
	default:
		return false
	}
}
