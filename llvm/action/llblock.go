package llactions

import (
	"fmt"

	generator "github.com/damianr1602/simplelang/llvm/generator"
	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
	"github.com/damianr1602/simplelang/util"
)

// var tsl.LocalVariables map[string]util.VarType = make(map[string]util.VarType)
// var globalVaricvables map[string]util.VarType = make(map[string]util.VarType)
// var tsl.ArrayName map[string]string = make(map[string]string)
// var tsl.CalculationsStack util.Stack

// TreeShapeListener TODO
type TreeShapeListener struct {
	*parser.BasesimplelangListener
	Llgen             generator.LLGenerate
	LocalVariables    map[string]util.VarType
	GlobalVariables   map[string]util.VarType
	FunctionVariables map[string]util.VarType
	ArrayName         map[string]string
	CalculationsStack util.Stack
	Global            bool
}

// NewTreeShapeListener TODO
func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// ExitLet impl
func (tsl *TreeShapeListener) ExitLet(ctx *parser.LetContext) {
	logger.Log.Println("Statment: ", ctx.GetText())
	variable := ctx.ID().GetText()

	if value, ok := tsl.CalculationsStack.Pop().(util.Value); ok {
		logger.Log.Println("value: ", value)
		if tsl.VariableOccupated(util.NewValue(variable, value.VarType)) {
			tsl.ShowError(ctx.GetStart().GetLine(), "Cannot declare "+variable+" because of previous use:")
			return
		}
		if value.VarType == util.INT {
			tsl.Llgen.AssignInt(tsl.SetVariable(variable, value.VarType), value.Name)
		} else if value.VarType == util.REAL {
			tsl.Llgen.AssignDouble(tsl.SetVariable(variable, value.VarType), value.Name)
		} else if value.VarType == util.STRING {
			tsl.Llgen.DeclareString(variable)
			tsl.Llgen.AssignString(variable, value.Name, false)
		} else if value.VarType == util.FUNC {
			tsl.Llgen.AssignInt(tsl.SetVariable(variable, util.INT), "%"+value.Name)
		} else if value.VarType == util.ARRAY {
			arrayLen := tsl.ArrayName[value.Name]
			tsl.Llgen.DeclareArrElemInt(variable)
			logger.Log.Println("[variable]: ", variable, "tsl.ArrayName[variable]", arrayLen, "map: ", tsl.ArrayName)
			logger.Log.Println("Exit ExitLet - variable: ", variable, "value: ", value, ", arrayLen: ", arrayLen)
		} else {
			logger.Log.Fatalf("Unknown variable")
		}
		if tsl.Global {
			tsl.GlobalVariables[variable] = value.VarType
		}
	}

}

// ExitShow impl
func (tsl *TreeShapeListener) ExitShow(ctx *parser.ShowContext) {
	variable := ctx.ID().GetText()
	logger.Log.Println("variable: ", variable)
	valueType, found := tsl.LocalVariables[variable]
	if found {
		variable = "%" + variable
	} else {
		valueType, found = tsl.GlobalVariables[variable]
		variable = "@" + variable
	}

	if found {
		if valueType == util.INT || valueType == util.FUNC {
			tsl.Llgen.PrintfInt(variable)
		} else if valueType == util.REAL {
			tsl.Llgen.PrintfDouble(variable)
		} else if valueType == util.STRING {
			tsl.Llgen.PrintfString(variable)
		} else if valueType == util.UNKNOWN {
			tsl.ShowError(ctx.GetStart().GetLine(), "Cannot show "+variable+" because not defined:")
			logger.Log.Fatalf("Unknown variable")
		} else {
			tsl.ShowError(ctx.GetStart().GetLine(), "Cannot show "+variable+" because not defined:")
			logger.Log.Fatalf("Unknown variable")
			return
		}
	} else {
		tsl.ShowError(ctx.GetStart().GetLine(), "Cannot show "+variable+" because not defined:")
		logger.Log.Fatalf("Unknown variable")
	}
	logger.Log.Println("Exit ExitShow - variable: ", variable, "valueType: ", valueType.String())
}

// ExitReadint impl
func (tsl *TreeShapeListener) ExitReadint(ctx *parser.ReadintContext) {
	variable := ctx.ID().GetText()
	tsl.Llgen.ScanfInt(tsl.SetVariable(variable, util.INT))

	logger.Log.Println("Exit ExitReadint - variable: ", variable)
}

// ExitReaddouble impl
func (tsl *TreeShapeListener) ExitReaddouble(ctx *parser.ReaddoubleContext) {
	variable := ctx.ID().GetText()
	tsl.Llgen.ScanfDouble(tsl.SetVariable(variable, util.REAL))

	logger.Log.Println("Exit ExitReadint - variable: ", variable, "type: ", tsl.LocalVariables[variable])
}

// EnterBlock impl
func (tsl *TreeShapeListener) EnterBlock(ctx *parser.BlockContext) {
	if _, ok := ctx.GetParent().(*parser.IfContext); ok {
		logger.Log.Println("is ok? :", ok)
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.BlockIfEnter()
	}
}

// ExitBlock impl
func (tsl *TreeShapeListener) ExitBlock(ctx *parser.BlockContext) {
	if _, ok := ctx.GetParent().(*parser.RepeatContext); ok {
		logger.Log.Println("is ok ctx? :", ok, " RepeatContext")
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.EndLoop()
	} else if _, ok := ctx.GetParent().(*parser.IfContext); ok {
		logger.Log.Println("is ok ctx? :", ok, " IfContext")
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.BlockIfExit()
	}
	logger.Log.Println("ExitBlock exit")
}

// EnterProg impl
func (tsl *TreeShapeListener) EnterProg(ctx *parser.ProgContext) {
	tsl.Global = true
	logger.Log.Println("Prog global flag: ", tsl.Global)
}

// ExitProg impl
func (tsl *TreeShapeListener) ExitProg(ctx *parser.ProgContext) {
	tsl.Llgen.StopMain()
	fmt.Println(tsl.Llgen.Generate())

	logger.Log.Println("Prog low-level code file generated")
}

// ShowError error line number
func (tsl *TreeShapeListener) ShowError(line int, msg string) {
	fmt.Println("Error, line ", line, ", "+msg)
}

// VariableOccupated status
func (tsl *TreeShapeListener) VariableOccupated(value util.Value) bool {
	logger.Log.Println("Value: ", value)
	logger.Log.Println("GlobalVariables: ", tsl.GlobalVariables)
	logger.Log.Println("LocalVariables: ", tsl.LocalVariables)
	logger.Log.Println("FunctionVariables: ", tsl.FunctionVariables)

	_, foundG := tsl.GlobalVariables[value.Name]
	_, foundL := tsl.LocalVariables[value.Name]

	if (foundG) && tsl.Global {
		logger.Log.Println("Found in @ variables")
		return true
	} else if foundL && tsl.Global == false {
		logger.Log.Println("Found in % variables")
		return true
	}
	return false
}
