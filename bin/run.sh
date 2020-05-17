#!/bin/bash
antlr4="java -Xmx500M -cp "/usr/local/lib/antlr-4.8-complete.jar:$CLASSPATH" org.antlr.v4.Tool"

cd /home/damian/Desktop/studia/mgr/simplelang/parser/
$antlr4 -Dlanguage=Go simplelang.g4
# w katalogu znajduje sie main.go
cd ../
go install
cd bin/
./simplelang test.simplelang > simple.ll
lli simple.ll


