package llgenerator

import (
	"strconv"

	logger "github.com/damianr1602/simplelang/logging"
)

// FuncStart llvm
func (llgen *LLGenerate) FuncStart(funcName string) {
	llgen.MainText += llgen.BufferText
	llgen.MainReg = llgen.Reg
	llgen.BufferText = "define i32 @" + funcName + "() nounwind {\n"
	llgen.Reg = 1
	logger.Log.Println("Start func ", funcName)
}

// FuncStop llvm
func (llgen *LLGenerate) FuncStop() {
	llgen.BufferText += "ret i32 %" + strconv.Itoa(llgen.Reg-1) + "\n"
	llgen.BufferText += "}\n"
	llgen.HeaderText += llgen.BufferText
	llgen.BufferText = ""
	llgen.Reg = llgen.MainReg
	logger.Log.Println("FuncStop func, returning %", llgen.Reg-1)
}

// CallFunc llvm
func (llgen *LLGenerate) CallFunc(funcName string) {
	llgen.BufferText += "%" + funcName + " = call i32 @" + funcName + "()\n"
	logger.Log.Println("CallFunc ", funcName)
}

// StopMain llvm
func (llgen *LLGenerate) StopMain() {
	llgen.MainText += llgen.BufferText
	logger.Log.Println("Stop Main")
}
