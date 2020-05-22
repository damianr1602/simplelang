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
