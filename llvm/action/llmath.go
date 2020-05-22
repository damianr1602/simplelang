package llactions

import (
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
