package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	llactions "github.com/damianr1602/simplelang/llvm/action"
	llgenerator "github.com/damianr1602/simplelang/llvm/generator"
	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
	"github.com/damianr1602/simplelang/util"
)

func main() {
	logger.Log.Printf("### Start program ###")

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewsimplelangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewsimplelangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()

	var localVariables map[string]util.VarType = make(map[string]util.VarType)
	var globalVaricvables map[string]util.VarType = make(map[string]util.VarType)
	var functionVariables map[string]util.VarType = make(map[string]util.VarType)
	var arrayName map[string]string = make(map[string]string)
	var stack util.Stack

	llact := llactions.NewTreeShapeListener()
	llact.Llgen = llgenerator.NewLLGenerate()
	llact.LocalVariables = localVariables
	llact.GlobalVariables = globalVaricvables
	llact.FunctionVariables = functionVariables
	llact.ArrayName = arrayName
	llact.CalculationsStack = stack
	
	antlr.ParseTreeWalkerDefault.Walk(llact, tree)
	logger.Log.Printf("Antlr finished processing file")
}
