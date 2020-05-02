// Code generated from simplerlang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // simplerlang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesimplerlangListener is a complete listener for a parse tree produced by simplerlangParser.
type BasesimplerlangListener struct{}

var _ simplerlangListener = &BasesimplerlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesimplerlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesimplerlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesimplerlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesimplerlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasesimplerlangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasesimplerlangListener) ExitProgram(ctx *ProgramContext) {}

// EnterShow is called when production show is entered.
func (s *BasesimplerlangListener) EnterShow(ctx *ShowContext) {}

// ExitShow is called when production show is exited.
func (s *BasesimplerlangListener) ExitShow(ctx *ShowContext) {}

// EnterLet is called when production let is entered.
func (s *BasesimplerlangListener) EnterLet(ctx *LetContext) {}

// ExitLet is called when production let is exited.
func (s *BasesimplerlangListener) ExitLet(ctx *LetContext) {}

// EnterReadint is called when production readint is entered.
func (s *BasesimplerlangListener) EnterReadint(ctx *ReadintContext) {}

// ExitReadint is called when production readint is exited.
func (s *BasesimplerlangListener) ExitReadint(ctx *ReadintContext) {}

// EnterReaddouble is called when production readdouble is entered.
func (s *BasesimplerlangListener) EnterReaddouble(ctx *ReaddoubleContext) {}

// ExitReaddouble is called when production readdouble is exited.
func (s *BasesimplerlangListener) ExitReaddouble(ctx *ReaddoubleContext) {}

// EnterSingle0 is called when production single0 is entered.
func (s *BasesimplerlangListener) EnterSingle0(ctx *Single0Context) {}

// ExitSingle0 is called when production single0 is exited.
func (s *BasesimplerlangListener) ExitSingle0(ctx *Single0Context) {}

// EnterAdd is called when production add is entered.
func (s *BasesimplerlangListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BasesimplerlangListener) ExitAdd(ctx *AddContext) {}

// EnterSub is called when production sub is entered.
func (s *BasesimplerlangListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BasesimplerlangListener) ExitSub(ctx *SubContext) {}

// EnterSingle1 is called when production single1 is entered.
func (s *BasesimplerlangListener) EnterSingle1(ctx *Single1Context) {}

// ExitSingle1 is called when production single1 is exited.
func (s *BasesimplerlangListener) ExitSingle1(ctx *Single1Context) {}

// EnterMul is called when production mul is entered.
func (s *BasesimplerlangListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production mul is exited.
func (s *BasesimplerlangListener) ExitMul(ctx *MulContext) {}

// EnterDiv is called when production div is entered.
func (s *BasesimplerlangListener) EnterDiv(ctx *DivContext) {}

// ExitDiv is called when production div is exited.
func (s *BasesimplerlangListener) ExitDiv(ctx *DivContext) {}

// EnterInt is called when production int is entered.
func (s *BasesimplerlangListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BasesimplerlangListener) ExitInt(ctx *IntContext) {}

// EnterReal is called when production real is entered.
func (s *BasesimplerlangListener) EnterReal(ctx *RealContext) {}

// ExitReal is called when production real is exited.
func (s *BasesimplerlangListener) ExitReal(ctx *RealContext) {}

// EnterToint is called when production toint is entered.
func (s *BasesimplerlangListener) EnterToint(ctx *TointContext) {}

// ExitToint is called when production toint is exited.
func (s *BasesimplerlangListener) ExitToint(ctx *TointContext) {}

// EnterToreal is called when production toreal is entered.
func (s *BasesimplerlangListener) EnterToreal(ctx *TorealContext) {}

// ExitToreal is called when production toreal is exited.
func (s *BasesimplerlangListener) ExitToreal(ctx *TorealContext) {}

// EnterPar is called when production par is entered.
func (s *BasesimplerlangListener) EnterPar(ctx *ParContext) {}

// ExitPar is called when production par is exited.
func (s *BasesimplerlangListener) ExitPar(ctx *ParContext) {}
