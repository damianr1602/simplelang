package util

// VarType is enum type
type VarType int

// enum representatoin
const (
	INT VarType = iota
	REAL
	STRING
	ARRAY
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

func (s VarType) String() string {
	return [...]string{"INT", "REAL", "STRING", "ARRAY", "UNKNOWN"}[s]
}
