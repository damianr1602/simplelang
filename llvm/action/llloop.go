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


// ExitEqual impl
func (tsl *TreeShapeListener) ExitEqual(ctx *parser.EqualContext) {
	id := ctx.ID().GetText()
	value := ctx.INT().GetText()
	valueType, found := tsl.LocalVariables[id]
	if found == false {
		valueType, found = tsl.GlobalVariables[id]
	}

	if found {
		tsl.Llgen.LoadInt(tsl.SetVariable(id, valueType))
		tsl.Llgen.Icmp(tsl.SetVariable(id, valueType), value)
	} else {
		logger.Log.Println("ExitEqual not found variable")
	}
	logger.Log.Println("ExitEqual exit")
}

// ExitRepetitions impl
func (tsl *TreeShapeListener) ExitRepetitions(ctx *parser.RepetitionsContext) {
	variable := ctx.GetText()
	valueType, found := tsl.LocalVariables[variable]

	logger.Log.Println("variableMap[variable]: ", tsl.LocalVariables, "variable: ", variable, "len(variableMap, ", len(tsl.LocalVariables), "is found? :", found)
	if found {
		if valueType == util.INT {
			tsl.Llgen.StartLoop(tsl.SetVariable(variable, valueType))

		}
	} else {
		repCounter, err := strconv.Atoi(variable)
		if err != nil {
			tsl.ShowError(ctx.GetStart().GetLine(), "variable not int "+variable)
		}
		valueName := "repCounter" + strconv.Itoa(repCounter+len(tsl.LocalVariables))
		tsl.Llgen.DeclareInt(valueName, false)
		tsl.Llgen.AssignInt("%"+valueName, variable)

		tsl.Llgen.StartLoop("%" + valueName)
	}

}
