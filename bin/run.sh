#!/bin/bash
antlr4="java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool"

cd parser/
$antlr4 -Dlanguage=Go simplerlang.g4
# w katalogu znajduje sie main.go
cd ../
go build
./simple_language test2.simplerlang > simple.ll
lli simple.ll


