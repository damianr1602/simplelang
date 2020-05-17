package util

// VarType is enum type
type VarType int

// enum representatoin
const (
	UNKNOWN VarType = iota
	INT
	REAL
	STRING
	ARRAY
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

func (s VarType) String() string {
	return [...]string{"UNKNOWN", "INT", "REAL", "STRING", "ARRAY"}[s]
}
