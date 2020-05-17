package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	llactions "github.com/damianr1602/simplelang/llvm/action"
	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
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
	antlr.ParseTreeWalkerDefault.Walk(llactions.NewTreeShapeListener(), tree)
	logger.Log.Printf("Antlr finished processing file")
}
