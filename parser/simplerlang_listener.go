// Code generated from simplerlang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // simplerlang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// simplerlangListener is a complete listener for a parse tree produced by simplerlangParser.
type simplerlangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

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

	// EnterAssignArrayElem is called when entering the assignArrayElem production.
	EnterAssignArrayElem(c *AssignArrayElemContext)

	// EnterReal is called when entering the real production.
	EnterReal(c *RealContext)

	// EnterToint is called when entering the toint production.
	EnterToint(c *TointContext)

	// EnterToreal is called when entering the toreal production.
	EnterToreal(c *TorealContext)

	// EnterPar is called when entering the par production.
	EnterPar(c *ParContext)

	// EnterArray_items is called when entering the array_items production.
	EnterArray_items(c *Array_itemsContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

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

	// ExitAssignArrayElem is called when exiting the assignArrayElem production.
	ExitAssignArrayElem(c *AssignArrayElemContext)

	// ExitReal is called when exiting the real production.
	ExitReal(c *RealContext)

	// ExitToint is called when exiting the toint production.
	ExitToint(c *TointContext)

	// ExitToreal is called when exiting the toreal production.
	ExitToreal(c *TorealContext)

	// ExitPar is called when exiting the par production.
	ExitPar(c *ParContext)

	// ExitArray_items is called when exiting the array_items production.
	ExitArray_items(c *Array_itemsContext)
}
