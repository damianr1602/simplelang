package llactions

import (
	// "strconv"

	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
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

// ExitStructure impl
func (tsl *TreeShapeListener) ExitStructure(ctx *parser.StructureContext) {
	logger.Log.Println("ExitStructure: ", ctx.GetText())
	sName := ctx.StructName().GetText()
	var elemN int
	if elem, ok := ctx.SBlock().(*parser.SBlockContext); ok {
		elemN = len(elem.AllID())
	}
	logger.Log.Println("ExitStructure: ", ctx.SBlock().GetText(), ctx.StructName().GetText())

	tsl.Llgen.StructDeclare(sName, elemN)
}

// EnterStructName impl
func (tsl *TreeShapeListener) EnterStructName(ctx *parser.StructNameContext) {
	structName := ctx.ID().GetText()
	logger.Log.Println("EnterStructName: ", structName)
}
