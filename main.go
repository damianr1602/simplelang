package main

import (
	"os"

	"github.com/damianr1602/simplelang/llvm/action"
	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	logger.Log.Printf("### Start program ###")

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewsimplerlangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewsimplerlangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(llactions.NewTreeShapeListener(), tree)
	logger.Log.Printf("Antlr finished processing file")
}
