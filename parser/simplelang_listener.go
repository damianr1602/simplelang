// Code generated from simplelang.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // simplelang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// simplelangListener is a complete listener for a parse tree produced by simplelangParser.
type simplelangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterShow is called when entering the show production.
	EnterShow(c *ShowContext)

	// EnterLet is called when entering the let production.
	EnterLet(c *LetContext)

	// EnterReadint is called when entering the readint production.
	EnterReadint(c *ReadintContext)

	// EnterIntarray is called when entering the intarray production.
	EnterIntarray(c *IntarrayContext)

	// EnterReaddouble is called when entering the readdouble production.
	EnterReaddouble(c *ReaddoubleContext)

	// EnterShowArrayElem is called when entering the showArrayElem production.
	EnterShowArrayElem(c *ShowArrayElemContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterRepeat is called when entering the repeat production.
	EnterRepeat(c *RepeatContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterStructure is called when entering the structure production.
	EnterStructure(c *StructureContext)

	// EnterSBlock is called when entering the sBlock production.
	EnterSBlock(c *SBlockContext)

	// EnterStructName is called when entering the structName production.
	EnterStructName(c *StructNameContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterFBlock is called when entering the fBlock production.
	EnterFBlock(c *FBlockContext)

	// EnterResult is called when entering the result production.
	EnterResult(c *ResultContext)

	// EnterSingle0 is called when entering the single0 production.
	EnterSingle0(c *Single0Context)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterSub is called when entering the sub production.
	EnterSub(c *SubContext)

	// EnterSingle1 is called when entering the single1 production.
	EnterSingle1(c *Single1Context)

	// EnterMul is called when entering the mul production.
	EnterMul(c *MulContext)

	// EnterDiv is called when entering the div production.
	EnterDiv(c *DivContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterAssignArrayElem is called when entering the assignArrayElem production.
	EnterAssignArrayElem(c *AssignArrayElemContext)

	// EnterReal is called when entering the real production.
	EnterReal(c *RealContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterToint is called when entering the toint production.
	EnterToint(c *TointContext)

	// EnterToreal is called when entering the toreal production.
	EnterToreal(c *TorealContext)

	// EnterPar is called when entering the par production.
	EnterPar(c *ParContext)

	// EnterRepetitions is called when entering the repetitions production.
	EnterRepetitions(c *RepetitionsContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterArray_items is called when entering the array_items production.
	EnterArray_items(c *Array_itemsContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitShow is called when exiting the show production.
	ExitShow(c *ShowContext)

	// ExitLet is called when exiting the let production.
	ExitLet(c *LetContext)

	// ExitReadint is called when exiting the readint production.
	ExitReadint(c *ReadintContext)

	// ExitIntarray is called when exiting the intarray production.
	ExitIntarray(c *IntarrayContext)

	// ExitReaddouble is called when exiting the readdouble production.
	ExitReaddouble(c *ReaddoubleContext)

	// ExitShowArrayElem is called when exiting the showArrayElem production.
	ExitShowArrayElem(c *ShowArrayElemContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitRepeat is called when exiting the repeat production.
	ExitRepeat(c *RepeatContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitStructure is called when exiting the structure production.
	ExitStructure(c *StructureContext)

	// ExitSBlock is called when exiting the sBlock production.
	ExitSBlock(c *SBlockContext)

	// ExitStructName is called when exiting the structName production.
	ExitStructName(c *StructNameContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitFBlock is called when exiting the fBlock production.
	ExitFBlock(c *FBlockContext)

	// ExitResult is called when exiting the result production.
	ExitResult(c *ResultContext)

	// ExitSingle0 is called when exiting the single0 production.
	ExitSingle0(c *Single0Context)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitSub is called when exiting the sub production.
	ExitSub(c *SubContext)

	// ExitSingle1 is called when exiting the single1 production.
	ExitSingle1(c *Single1Context)

	// ExitMul is called when exiting the mul production.
	ExitMul(c *MulContext)

	// ExitDiv is called when exiting the div production.
	ExitDiv(c *DivContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitAssignArrayElem is called when exiting the assignArrayElem production.
	ExitAssignArrayElem(c *AssignArrayElemContext)

	// ExitReal is called when exiting the real production.
	ExitReal(c *RealContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitToint is called when exiting the toint production.
	ExitToint(c *TointContext)

	// ExitToreal is called when exiting the toreal production.
	ExitToreal(c *TorealContext)

	// ExitPar is called when exiting the par production.
	ExitPar(c *ParContext)

	// ExitRepetitions is called when exiting the repetitions production.
	ExitRepetitions(c *RepetitionsContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitArray_items is called when exiting the array_items production.
	ExitArray_items(c *Array_itemsContext)
}
