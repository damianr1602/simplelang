// Code generated from simplelang.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // simplelang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesimplelangListener is a complete listener for a parse tree produced by simplelangParser.
type BasesimplelangListener struct{}

var _ simplelangListener = &BasesimplelangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesimplelangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesimplelangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesimplelangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesimplelangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasesimplelangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasesimplelangListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BasesimplelangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasesimplelangListener) ExitBlock(ctx *BlockContext) {}

// EnterShow is called when production show is entered.
func (s *BasesimplelangListener) EnterShow(ctx *ShowContext) {}

// ExitShow is called when production show is exited.
func (s *BasesimplelangListener) ExitShow(ctx *ShowContext) {}

// EnterLet is called when production let is entered.
func (s *BasesimplelangListener) EnterLet(ctx *LetContext) {}

// ExitLet is called when production let is exited.
func (s *BasesimplelangListener) ExitLet(ctx *LetContext) {}

// EnterReadint is called when production readint is entered.
func (s *BasesimplelangListener) EnterReadint(ctx *ReadintContext) {}

// ExitReadint is called when production readint is exited.
func (s *BasesimplelangListener) ExitReadint(ctx *ReadintContext) {}

// EnterIntarray is called when production intarray is entered.
func (s *BasesimplelangListener) EnterIntarray(ctx *IntarrayContext) {}

// ExitIntarray is called when production intarray is exited.
func (s *BasesimplelangListener) ExitIntarray(ctx *IntarrayContext) {}

// EnterReaddouble is called when production readdouble is entered.
func (s *BasesimplelangListener) EnterReaddouble(ctx *ReaddoubleContext) {}

// ExitReaddouble is called when production readdouble is exited.
func (s *BasesimplelangListener) ExitReaddouble(ctx *ReaddoubleContext) {}

// EnterShowArrayElem is called when production showArrayElem is entered.
func (s *BasesimplelangListener) EnterShowArrayElem(ctx *ShowArrayElemContext) {}

// ExitShowArrayElem is called when production showArrayElem is exited.
func (s *BasesimplelangListener) ExitShowArrayElem(ctx *ShowArrayElemContext) {}

// EnterIf is called when production if is entered.
func (s *BasesimplelangListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BasesimplelangListener) ExitIf(ctx *IfContext) {}

// EnterRepeat is called when production repeat is entered.
func (s *BasesimplelangListener) EnterRepeat(ctx *RepeatContext) {}

// ExitRepeat is called when production repeat is exited.
func (s *BasesimplelangListener) ExitRepeat(ctx *RepeatContext) {}

// EnterSingle0 is called when production single0 is entered.
func (s *BasesimplelangListener) EnterSingle0(ctx *Single0Context) {}

// ExitSingle0 is called when production single0 is exited.
func (s *BasesimplelangListener) ExitSingle0(ctx *Single0Context) {}

// EnterAdd is called when production add is entered.
func (s *BasesimplelangListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BasesimplelangListener) ExitAdd(ctx *AddContext) {}

// EnterSub is called when production sub is entered.
func (s *BasesimplelangListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BasesimplelangListener) ExitSub(ctx *SubContext) {}

// EnterSingle1 is called when production single1 is entered.
func (s *BasesimplelangListener) EnterSingle1(ctx *Single1Context) {}

// ExitSingle1 is called when production single1 is exited.
func (s *BasesimplelangListener) ExitSingle1(ctx *Single1Context) {}

// EnterMul is called when production mul is entered.
func (s *BasesimplelangListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production mul is exited.
func (s *BasesimplelangListener) ExitMul(ctx *MulContext) {}

// EnterDiv is called when production div is entered.
func (s *BasesimplelangListener) EnterDiv(ctx *DivContext) {}

// ExitDiv is called when production div is exited.
func (s *BasesimplelangListener) ExitDiv(ctx *DivContext) {}

// EnterInt is called when production int is entered.
func (s *BasesimplelangListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BasesimplelangListener) ExitInt(ctx *IntContext) {}

// EnterId is called when production id is entered.
func (s *BasesimplelangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BasesimplelangListener) ExitId(ctx *IdContext) {}

// EnterAssignArrayElem is called when production assignArrayElem is entered.
func (s *BasesimplelangListener) EnterAssignArrayElem(ctx *AssignArrayElemContext) {}

// ExitAssignArrayElem is called when production assignArrayElem is exited.
func (s *BasesimplelangListener) ExitAssignArrayElem(ctx *AssignArrayElemContext) {}

// EnterReal is called when production real is entered.
func (s *BasesimplelangListener) EnterReal(ctx *RealContext) {}

// ExitReal is called when production real is exited.
func (s *BasesimplelangListener) ExitReal(ctx *RealContext) {}

// EnterString is called when production string is entered.
func (s *BasesimplelangListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BasesimplelangListener) ExitString(ctx *StringContext) {}

// EnterToint is called when production toint is entered.
func (s *BasesimplelangListener) EnterToint(ctx *TointContext) {}

// ExitToint is called when production toint is exited.
func (s *BasesimplelangListener) ExitToint(ctx *TointContext) {}

// EnterToreal is called when production toreal is entered.
func (s *BasesimplelangListener) EnterToreal(ctx *TorealContext) {}

// ExitToreal is called when production toreal is exited.
func (s *BasesimplelangListener) ExitToreal(ctx *TorealContext) {}

// EnterPar is called when production par is entered.
func (s *BasesimplelangListener) EnterPar(ctx *ParContext) {}

// ExitPar is called when production par is exited.
func (s *BasesimplelangListener) ExitPar(ctx *ParContext) {}

// EnterRepetitions is called when production repetitions is entered.
func (s *BasesimplelangListener) EnterRepetitions(ctx *RepetitionsContext) {}

// ExitRepetitions is called when production repetitions is exited.
func (s *BasesimplelangListener) ExitRepetitions(ctx *RepetitionsContext) {}

// EnterEqual is called when production equal is entered.
func (s *BasesimplelangListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BasesimplelangListener) ExitEqual(ctx *EqualContext) {}

// EnterArray_items is called when production array_items is entered.
func (s *BasesimplelangListener) EnterArray_items(ctx *Array_itemsContext) {}

// ExitArray_items is called when production array_items is exited.
func (s *BasesimplelangListener) ExitArray_items(ctx *Array_itemsContext) {}
