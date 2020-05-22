package llactions

import (
	// "strconv"

	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
	"github.com/damianr1602/simplelang/util"
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

// ExitId impl
func (tsl *TreeShapeListener) ExitId(ctx *parser.IdContext) {
	logger.Log.Println("ExitId:", ctx.GetText())
	if function == ctx.GetText() {

		tsl.Llgen.CallFunc(ctx.ID().GetText())
		tsl.CalculationsStack.Push(util.NewValue(ctx.ID().GetText(), util.FUNC))
	}

	if _, ok := ctx.GetParent().(*parser.LetContext); ok {
		logger.Log.Println("is ok? :", ok)
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.CallFunc(ctx.ID().GetText())
		tsl.CalculationsStack.Push(util.NewValue(ctx.ID().GetText(), util.FUNC))
	} else if _, ok := ctx.GetParent().(*parser.Single0Context); ok {
		logger.Log.Println("is ok? :", ok)
		logger.Log.Println("Single0Context:", ctx.GetText())

	}

	logger.Log.Println("Exit ExitId - stack: ", tsl.CalculationsStack)
}

// ExitCall impl
func (tsl *TreeShapeListener) ExitCall(ctx *parser.CallContext) {
	tsl.Llgen.CallFunc(ctx.ID().GetText())
	logger.Log.Println("Exit ExitCall - variable")
}

// EnterFBlock impl
func (tsl *TreeShapeListener) EnterFBlock(ctx *parser.FBlockContext) {
	tsl.Global = false
	logger.Log.Println("Exit EnterFBlock - ")
}

// ExitResult impl
func (tsl *TreeShapeListener) ExitResult(ctx *parser.ResultContext) {
	result := ctx.ID().GetText()
	logger.Log.Println("LocalVariables: ", tsl.LocalVariables)

	_, found := tsl.LocalVariables[result]
	if found {
		tsl.Llgen.LoadInt("%" + result)
	}
	logger.Log.Println("Result: ", result)

	tsl.Llgen.FuncStop()
	tsl.LocalVariables = make(map[string]util.VarType)
	logger.Log.Println("Exit ExitFBlock - ")
	tsl.Global = true
}

// ExitFuncName impl
func (tsl *TreeShapeListener) ExitFuncName(ctx *parser.FuncNameContext) {
	funcName := ctx.ID().GetText()
	tsl.FunctionVariables[funcName] = util.FUNC
	function = funcName
	tsl.Llgen.FuncStart(funcName)

	logger.Log.Println("Exit ExitFuncName ")
}

// SetVariable impl
func (tsl *TreeShapeListener) SetVariable(varName string, valueType util.VarType) string {
	logger.Log.Println("GlobalVariables: ", tsl.GlobalVariables)
	logger.Log.Println("LocalVariables: ", tsl.LocalVariables)
	logger.Log.Println("FunctionVariables: ", tsl.FunctionVariables)
	logger.Log.Println("Is global scope?: ", tsl.Global)
	logger.Log.Println("Varaible: ", varName, "valueType: ", valueType)

	var name string
	if tsl.Global {
		if _, found := tsl.GlobalVariables[varName]; found == false {
			tsl.GlobalVariables[varName] = valueType
			if valueType == util.INT {
				tsl.Llgen.DeclareInt(varName, true)
			} else if valueType == util.REAL {
				tsl.Llgen.DeclareDouble(varName, true)
			}
		}
		name = "@" + varName
	} else {
		if _, found := tsl.LocalVariables[varName]; found == false {
			tsl.LocalVariables[varName] = valueType
			if valueType == util.INT {
				tsl.Llgen.DeclareInt(varName, false)
			} else if valueType == util.REAL {
				tsl.Llgen.DeclareDouble(varName, false)
			}
		} else if _, found := tsl.GlobalVariables[varName]; found {
			logger.Log.Println("set variable name: ", name)
			return "@" + varName
		}
		name = "%" + varName
	}
	logger.Log.Println("set variable name: ", name)
	return name
}
