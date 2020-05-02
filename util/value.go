package util

// VarType is enum type
type VarType int

// enum representatoin
const (
	INT VarType = iota
	REAL
	UNKNOWN
)

// Value for simple lang variables
type Value struct {
	Name    string
	VarType VarType
}

// var LLGenerateInstance = NewLLGenerate()

// NewValue constuctor for Value
func NewValue(name string, varType VarType) Value {
	valueInit := Value{}
	valueInit.Name = name
	valueInit.VarType = varType

	return valueInit
}
