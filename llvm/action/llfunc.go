package llactions

import (
	// "strconv"

	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
	"github.com/damianr1602/simplelang/util"
	// "github.com/damianr1602/simplelang/util"
)

// type TreeShapeListener struct {
// 	*parser.BasesimplelangListener
// 	Llgen generator.LLGenerate
// 	LocalVariables map[string]util.VarType
// 	GlobalVariables map[string]util.VarType
// 	ArrayName map[string]string
// 	CalculationsStack util.Stack
// 	Global bool
// }

var function, value string

// ExitCall impl
func (tsl *TreeShapeListener) ExitCall(ctx *parser.CallContext) {
	tsl.Llgen.CallFunc(ctx.ID().GetText())
	logger.Log.Println("Exit ExitReadint - variable")
}

// EnterFBlock impl
func (tsl *TreeShapeListener) EnterFBlock(ctx *parser.FBlockContext) {
	tsl.Global = false
	logger.Log.Println("Exit ExitReadint - variable")
}

// ExitFBlock impl
func (tsl *TreeShapeListener) ExitFBlock(ctx *parser.FBlockContext) {

	logger.Log.Println("Exit ExitReadint - variable")
}

// ExitFuncName impl
func (tsl *TreeShapeListener) ExitFuncName(ctx *parser.FuncNameContext) {
	funcName := ctx.ID().GetText()
	tsl.FunctionVariables[funcName] = util.UNKNOWN
	function = funcName
	tsl.Llgen.FuncStart(funcName)

	logger.Log.Println("Exit ExitFuncName ")
}

// SetVariable impl
func (tsl *TreeShapeListener) SetVariable(varName string) string {
	var name string
	if tsl.Global {
		if VarType, found := tsl.GlobalVariables[varName]; found == false {
			tsl.GlobalVariables[varName] = VarType
			tsl.Llgen.DeclareInt(varName, true)
		}
		name = "@" + varName
	} else {
		if VarType, found := tsl.LocalVariables[varName]; found == false {
			tsl.LocalVariables[varName] = VarType
			tsl.Llgen.DeclareInt(varName, false)
		}
		name = "%" + varName
	}
	logger.Log.Println("set variable name: ", name)

	return name
}
