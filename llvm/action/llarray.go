package llactions

import (
	// "strconv"

	"strconv"

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

// ExitIntarray impl
func (tsl *TreeShapeListener) ExitIntarray(ctx *parser.IntarrayContext) {
	variable := ctx.ID().GetText()
	tsl.LocalVariables[variable] = util.INT
	var elems = make([]string, 0, 10)
	if arrayElem := ctx.Array_items().(*parser.Array_itemsContext); arrayElem != nil {
		for _, elem := range arrayElem.AllINT() {
			elems = append(elems, elem.GetText())
		}
		tsl.Llgen.DeclareIntArray(variable, elems)
		tsl.ArrayName[variable] = strconv.Itoa(len(elems))
	}
	logger.Log.Println("ExitIntarray: ", elems)
}

// ExitShowArrayElem impl
func (tsl *TreeShapeListener) ExitShowArrayElem(ctx *parser.ShowArrayElemContext) {
	variable := ctx.ID().GetText()
	valueElem := ctx.INT().GetText()
	arrayLen := tsl.ArrayName[variable]
	valueType, found := tsl.LocalVariables[variable]
	if found {
		if valueType == util.INT {
			tsl.Llgen.PrintfArrElemInt(variable, valueElem, arrayLen)
		} else {
			logger.Log.Fatalf("Unknown variable")
		}
		logger.Log.Println("variable, valueElem, arrayLen -> ", variable, valueElem, arrayLen)
	}
}

// ExitAssignArrayElem impl
func (tsl *TreeShapeListener) ExitAssignArrayElem(ctx *parser.AssignArrayElemContext) {
	// logger.Log.Println("AssignArrayElem variable", ctx.GetText())
	tsl.CalculationsStack.Push(util.NewValue(ctx.ID().GetText(), util.ARRAY))
	logger.Log.Println("int pushed on tsl.CalculationsStack: ", ctx.ID().GetText())

	variable := ctx.ID().GetText()
	valueElem := ctx.INT().GetText()
	arrayLen := tsl.ArrayName[variable]
	valueType, found := tsl.LocalVariables[variable]
	if found {
		if valueType == util.INT {
			// AssignArrElemInt(variable string, valueElem string, arrayLen string)
			tsl.Llgen.AssignArrElemInt(variable, valueElem, arrayLen)
		} else {
			logger.Log.Fatalf("Unknown variable")
			panic("Unknown variable")
		}
		logger.Log.Println("ExitAssignArrayElem variable, valueElem, arrayLen -> ", variable, valueElem, arrayLen)
	} else {
		tsl.ShowError(ctx.GetStart().GetLine(), "variable not found"+variable)
	}
}
