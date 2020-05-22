package llgenerator

import (
	logger "github.com/damianr1602/simplelang/logging"
)

// StructDeclare llvm
func (llgen *LLGenerate) StructDeclare(structName string, n int) {
	var params string
	for i := 0; i < n; i++ {
		params += "double, "
	}
	params = params[:len(params)-2]
	llgen.HeaderText = "%struct." + structName + " = type { " + params + " }\n"
	logger.Log.Println("StructDeclare ", structName)
}
