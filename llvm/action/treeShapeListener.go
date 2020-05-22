package llactions

import (
	"fmt"
	"strconv"

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

// ExitAdd impl
func (tsl *TreeShapeListener) ExitAdd(ctx *parser.AddContext) {
	logger.Log.Println("Stack currnet state: ", tsl.CalculationsStack)
	val1 := tsl.CalculationsStack.Pop().(util.Value)
	val2 := tsl.CalculationsStack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			tsl.Llgen.AddInt(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			tsl.Llgen.AddDouble(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch")

		}
	}
	logger.Log.Println("Exit ExitAdd")
}

// ExitSub impl
func (tsl *TreeShapeListener) ExitSub(ctx *parser.SubContext) {
	val1 := tsl.CalculationsStack.Pop().(util.Value)
	val2 := tsl.CalculationsStack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			tsl.Llgen.SubInt(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			tsl.Llgen.SubDouble(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch during subtraction")

		}
	}
	logger.Log.Println("Exit ExitAdd")
}

// ExitMul impl
func (tsl *TreeShapeListener) ExitMul(ctx *parser.MulContext) {
	val1 := tsl.CalculationsStack.Pop().(util.Value)
	val2 := tsl.CalculationsStack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			tsl.Llgen.MulInt(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			tsl.Llgen.MulDouble(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch")

		}
	}
	logger.Log.Println("Exit ExitMul")
}

// ExitDiv impl
func (tsl *TreeShapeListener) ExitDiv(ctx *parser.DivContext) {
	val1 := tsl.CalculationsStack.Pop().(util.Value)
	val2 := tsl.CalculationsStack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			tsl.Llgen.DivInt(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			tsl.Llgen.DivDouble(val1.Name, val2.Name)
			tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch")

		}
	}
	logger.Log.Println("Exit ExitDiv")
}

// ExitInt impl
func (tsl *TreeShapeListener) ExitInt(ctx *parser.IntContext) {
	tsl.CalculationsStack.Push(util.NewValue(ctx.INT().GetText(), util.INT))
	logger.Log.Println("ExitInt - int pushed on tsl.CalculationsStack", tsl.CalculationsStack)
}

// ExitReal impl
func (tsl *TreeShapeListener) ExitReal(ctx *parser.RealContext) {
	tsl.CalculationsStack.Push(util.NewValue(ctx.REAL().GetText(), util.REAL))
	logger.Log.Println("ExitReal - real pushed on tsl.CalculationsStack", tsl.CalculationsStack)
}

// ExitString impl
func (tsl *TreeShapeListener) ExitString(ctx *parser.StringContext) {
	tsl.CalculationsStack.Push(util.NewValue(ctx.STRING().GetText(), util.STRING))
	logger.Log.Println("ExitString - string pushed on tsl.CalculationsStack", tsl.CalculationsStack)
}

// ExitToint impl
func (tsl *TreeShapeListener) ExitToint(ctx *parser.TointContext) {
	value := tsl.CalculationsStack.Pop().(util.Value)
	tsl.Llgen.Fptosi(value.Name)
	tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.INT))

	logger.Log.Println("ExitReal - value: ", value, " is now int", tsl.CalculationsStack)
}

// ExitToreal impl
func (tsl *TreeShapeListener) ExitToreal(ctx *parser.TorealContext) {
	value := tsl.CalculationsStack.Pop().(util.Value)
	tsl.Llgen.Sitofp(value.Name)
	tsl.CalculationsStack.Push(util.NewValue("%"+strconv.Itoa(tsl.Llgen.Reg-1), util.REAL))
	logger.Log.Println("ExitToreal - value: ", value, " is now real", tsl.CalculationsStack)
}

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

// ExitIf impl
func (tsl *TreeShapeListener) ExitIf(ctx *parser.IfContext) {
	logger.Log.Println("ExitIf exit")
}

// EnterBlock impl
func (tsl *TreeShapeListener) EnterBlock(ctx *parser.BlockContext) {
	if _, ok := ctx.GetParent().(*parser.IfContext); ok {
		logger.Log.Println("is ok? :", ok)
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.BlockIfEnter()
	}
	logger.Log.Println("EnterBlock exit")
}

// ExitBlockif impl
// func (tsl *TreeShapeListener) ExitBlockif(ctx *parser.BlockifContext) {
// 	tsl.Llgen.BlockIfExit()
// 	logger.Log.Println("ExitBlockif exit")
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

// func (s *BasesimplerlangListener) ExitBlock(ctx *BlockContext) {}
// func (s *BasesimplerlangListener) EnterProg(ctx *ProgContext) {}
// func (s *BasesimplerlangListener) ExitProg(ctx *ProgContext) {}

// ExitBlock impl
func (tsl *TreeShapeListener) ExitBlock(ctx *parser.BlockContext) {
	logger.Log.Println("ExitBlock enter", ctx.GetText())
	// logger.Log.Println("ExitBlock parent ", ctx.GetTypedRuleContexts(reflect.TypeOf(parser.RepetitionsContext)))

	if _, ok := ctx.GetParent().(*parser.RepeatContext); ok {
		logger.Log.Println("is ok ctx? :", ok, " RepeatContext")
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.EndLoop()
	} else if _, ok := ctx.GetParent().(*parser.IfContext); ok {
		logger.Log.Println("is ok ctx? :", ok, " IfContext")
		// logger.Log.Println("t=", t.GetText())
		tsl.Llgen.BlockIfExit()
	}

	//arrayElem := ctx.Array_items().(*parser.Array_itemsContext); arrayElem != nil {
	logger.Log.Println("ExitBlock exit")
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

	// tsl.Llgen.StartLoop(value)
	// logger.Log.Println("ExitRepetitions: ", ctx.GetText())
}

// EnterProg impl
func (tsl *TreeShapeListener) EnterProg(ctx *parser.ProgContext) {
	tsl.Global = true
	logger.Log.Println("EnterProg set global: ", tsl.Global)
}

// ExitProg impl
func (tsl *TreeShapeListener) ExitProg(ctx *parser.ProgContext) {
	tsl.Llgen.StopMain()
	fmt.Println(tsl.Llgen.Generate())

	logger.Log.Println("ExitProg low-level code file generated")
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
	// _, foundF := tsl.FunctionVariables[value.Name]
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
