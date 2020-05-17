package llactions

import (
	"fmt"
	"strconv"

	llgenerator "github.com/damianr1602/simplelang/llvm/generator"
	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
	"github.com/damianr1602/simplelang/util"
)

var variableMap map[string]util.VarType = make(map[string]util.VarType)
var arrayName map[string]string = make(map[string]string)
var llgen llgenerator.LLGenerate = llgenerator.LLGenerateInstance
var stack util.Stack

// TreeShapeListener TODO
type TreeShapeListener struct {
	*parser.BasesimplelangListener
}

// NewTreeShapeListener TODO
func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// ExitLet impl
func (tsl *TreeShapeListener) ExitLet(ctx *parser.LetContext) {
	logger.Log.Println("Statment: ", ctx.GetText())

	variable := ctx.ID().GetText()
	value := stack.Pop().(util.Value)
	variableMap[variable] = value.VarType

	if value.VarType == util.INT {
		llgen.DeclareInt(variable)
		llgen.AssignInt(variable, value.Name)
	} else if value.VarType == util.REAL {
		llgen.DeclareDouble(variable)
		llgen.AssignDouble(variable, value.Name)
	} else if value.VarType == util.STRING {
		llgen.DeclareString(variable)
		llgen.AssignString(variable, value.Name)
	} else if value.VarType == util.ARRAY {
		arrayLen := arrayName[value.Name]
		llgen.DeclareArrElemInt(variable)
		logger.Log.Println("[variable]: ", variable, "arrayName[variable]", arrayLen, "map: ", arrayName)
		logger.Log.Println("Exit ExitLet - variable: ", variable, "value: ", value, ", arrayLen: ", arrayLen)
	} else {
		logger.Log.Fatalf("Unknown variable")
	}
}

// ExitShow impl
func (tsl *TreeShapeListener) ExitShow(ctx *parser.ShowContext) {
	variable := ctx.ID().GetText()
	valueType, found := variableMap[variable]
	logger.Log.Println("variableMap[variable]: ", variableMap, "variable: ", variable)

	if found {
		if valueType == util.INT {
			llgen.PrintfInt(variable)
		} else if valueType == util.REAL {
			llgen.PrintfDouble(variable)
		} else if valueType == util.STRING {
			llgen.PrintfString(variable)
		} else if valueType == util.UNKNOWN {
			llgen.PrintfInt(variable)
		} else {
			logger.Log.Fatalf("Unknown variable")
		}
	}
	logger.Log.Println("Exit ExitShow - variable: ", variable, "valueType: ", valueType.String())
}

// ExitReadint impl
func (tsl *TreeShapeListener) ExitReadint(ctx *parser.ReadintContext) {
	variable := ctx.ID().GetText()
	variableMap[variable] = util.INT
	llgen.DeclareInt(variable)
	llgen.ScanfInt(variable)

	logger.Log.Println("Exit ExitReadint - variable: ", variable, "type: ", variableMap[variable])
}

// ExitReaddouble impl
func (tsl *TreeShapeListener) ExitReaddouble(ctx *parser.ReaddoubleContext) {
	variable := ctx.ID().GetText()
	variableMap[variable] = util.REAL
	llgen.DeclareDouble(variable)
	llgen.ScanfDouble(variable)

	logger.Log.Println("Exit ExitReadint - variable: ", variable, "type: ", variableMap[variable])
}

// ExitAdd impl
func (tsl *TreeShapeListener) ExitAdd(ctx *parser.AddContext) {
	val1 := stack.Pop().(util.Value)
	val2 := stack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			llgen.AddInt(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			llgen.AddDouble(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch")

		}
	}
	logger.Log.Println("Exit ExitAdd")
}

// ExitSub impl
func (tsl *TreeShapeListener) ExitSub(ctx *parser.SubContext) {
	val1 := stack.Pop().(util.Value)
	val2 := stack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			llgen.SubInt(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			llgen.SubDouble(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch during subtraction")

		}
	}
	logger.Log.Println("Exit ExitAdd")
}

// ExitMul impl
func (tsl *TreeShapeListener) ExitMul(ctx *parser.MulContext) {
	val1 := stack.Pop().(util.Value)
	val2 := stack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			llgen.MulInt(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			llgen.MulDouble(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch")

		}
	}
	logger.Log.Println("Exit ExitMul")
}

// ExitDiv impl
func (tsl *TreeShapeListener) ExitDiv(ctx *parser.DivContext) {
	val1 := stack.Pop().(util.Value)
	val2 := stack.Pop().(util.Value)
	if val1.VarType == val2.VarType {
		if val1.VarType == util.INT {
			llgen.DivInt(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.INT))
		} else if val1.VarType == util.REAL {
			llgen.DivDouble(val1.Name, val2.Name)
			stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.REAL))
		} else {
			logger.Log.Fatalf("Type mismatch")

		}
	}
	logger.Log.Println("Exit ExitDiv")
}

// ExitInt impl
func (tsl *TreeShapeListener) ExitInt(ctx *parser.IntContext) {
	stack.Push(util.NewValue(ctx.INT().GetText(), util.INT))
	logger.Log.Printf("ExitInt - int pushed on stack")
}

// ExitReal impl
func (tsl *TreeShapeListener) ExitReal(ctx *parser.RealContext) {
	stack.Push(util.NewValue(ctx.REAL().GetText(), util.REAL))
	logger.Log.Printf("ExitReal - real pushed on stack")
}

// ExitString impl
func (tsl *TreeShapeListener) ExitString(ctx *parser.StringContext) {
	stack.Push(util.NewValue(ctx.STRING().GetText(), util.STRING))
	logger.Log.Printf("ExitString - string pushed on stack")
}

// ExitToint impl
func (tsl *TreeShapeListener) ExitToint(ctx *parser.TointContext) {
	value := stack.Pop().(util.Value)
	llgen.Fptosi(value.Name)
	stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.INT))

	logger.Log.Println("ExitReal - value: ", value, " is now int")
}

// ExitToreal impl
func (tsl *TreeShapeListener) ExitToreal(ctx *parser.TorealContext) {
	value := stack.Pop().(util.Value)
	llgen.Sitofp(value.Name)
	stack.Push(util.NewValue("%"+strconv.Itoa(llgen.Reg-1), util.REAL))
	logger.Log.Println("ExitToreal - value: ", value, " is now real")
}

// ExitIntarray impl
func (tsl *TreeShapeListener) ExitIntarray(ctx *parser.IntarrayContext) {
	variable := ctx.ID().GetText()
	variableMap[variable] = util.INT
	var elems = make([]string, 0, 10)
	if arrayElem := ctx.Array_items().(*parser.Array_itemsContext); arrayElem != nil {
		for _, elem := range arrayElem.AllINT() {
			elems = append(elems, elem.GetText())
		}
		logger.Log.Println("ExitIntarray: ", elems[0])
		logger.Log.Println("ExitIntarray: ", elems[1])

		llgen.DeclareIntArray(variable, elems)
		arrayName[variable] = strconv.Itoa(len(elems))
	}
	logger.Log.Println("ExitIntarray: ", elems)
}

// ExitShowArrayElem impl
func (tsl *TreeShapeListener) ExitShowArrayElem(ctx *parser.ShowArrayElemContext) {
	variable := ctx.ID().GetText()
	valueElem := ctx.INT().GetText()
	arrayLen := arrayName[variable]
	valueType, found := variableMap[variable]
	if found {
		if valueType == util.INT {
			llgen.PrintfArrElemInt(variable, valueElem, arrayLen)
		} else {
			logger.Log.Fatalf("Unknown variable")
			panic("Unknown variable")
		}
		logger.Log.Println("variable, valueElem, arrayLen -> ", variable, valueElem, arrayLen)
	}
}

// ExitAssignArrayElem impl
func (tsl *TreeShapeListener) ExitAssignArrayElem(ctx *parser.AssignArrayElemContext) {
	logger.Log.Println("AssignArrayElem variable", ctx.GetText())
	stack.Push(util.NewValue(ctx.ID().GetText(), util.ARRAY))
	logger.Log.Println("int pushed on stack: ", ctx.ID().GetText())

	variable := ctx.ID().GetText()
	valueElem := ctx.INT().GetText()
	arrayLen := arrayName[variable]
	valueType, found := variableMap[variable]
	if found {
		if valueType == util.INT {
			// AssignArrElemInt(variable string, valueElem string, arrayLen string)
			llgen.AssignArrElemInt(variable, valueElem, arrayLen)
		} else {
			logger.Log.Fatalf("Unknown variable")
			panic("Unknown variable")
		}
		logger.Log.Println("ExitAssignArrayElem variable, valueElem, arrayLen -> ", variable, valueElem, arrayLen)
	} else {
		showError(ctx.GetStart().GetLine(), "variable not found"+variable)
	}
}

// ExitIf impl
func (tsl *TreeShapeListener) ExitIf(ctx *parser.IfContext) {
	logger.Log.Println("ExitIf exit")
}

// EnterBlockif impl
func (tsl *TreeShapeListener) EnterBlockif(ctx *parser.BlockifContext) {
	llgen.BlockIfEnter()
	logger.Log.Println("EnterBlockif exit")
}

// ExitBlockif impl
func (tsl *TreeShapeListener) ExitBlockif(ctx *parser.BlockifContext) {
	llgen.BlockIfExit()
	logger.Log.Println("ExitBlockif exit")
}

// ExitEqual impl
func (tsl *TreeShapeListener) ExitEqual(ctx *parser.EqualContext) {
	id := ctx.ID().GetText()
	value := ctx.INT().GetText()
	_, found := variableMap[id]

	if found {
		llgen.Icmp(id, value)
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

	if t, ok := ctx.GetParent().(*parser.RepeatContext); ok {
		logger.Log.Println("is ok? :", ok)
		logger.Log.Println("t=", t.GetText())
		llgen.EndLoop()
	} 
	

	//arrayElem := ctx.Array_items().(*parser.Array_itemsContext); arrayElem != nil {
	logger.Log.Println("ExitBlock exit")
}

// ExitRepetitions impl
func (tsl *TreeShapeListener) ExitRepetitions(ctx *parser.RepetitionsContext) {
	variable := ctx.GetText()
	valueType, found := variableMap[variable]
	logger.Log.Println("variableMap[variable]: ", variableMap, "variable: ", variable, "len(variableMap, ", len(variableMap), "is found? :", found)
	if found {
		if valueType == util.INT {
			llgen.StartLoop(variable)

		}
	} else {
		repCounter, err := strconv.Atoi(variable)
		if err != nil {
			showError(ctx.GetStart().GetLine(), "variable not int "+variable)
		}
		valueName := "repCounter" + strconv.Itoa(repCounter+len(variableMap))
		llgen.DeclareInt(valueName)
		llgen.AssignInt(valueName, variable)
		llgen.StartLoop(valueName)

	}

	// llgen.StartLoop(value)
	logger.Log.Println("ExitRepetitions: ", ctx.GetText())
}

// ExitProg impl
func (tsl *TreeShapeListener) ExitProg(ctx *parser.ProgContext) {
	fmt.Println(llgen.Generate())

	logger.Log.Printf("ExitProg low-level code file generated")
}

// TODO error line number
func showError(line int, msg string) {
	fmt.Println("Error, line ", line, ", "+msg)
}
