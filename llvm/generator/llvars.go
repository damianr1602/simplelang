package llgenerator

import (
	"strconv"

	logger "github.com/damianr1602/simplelang/logging"
)

//LoadInt llvm
func (llgen *LLGenerate) LoadInt(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load i32, i32* " + value + "\n"
	llgen.Reg++
}

//LoadDouble llvm
func (llgen *LLGenerate) LoadDouble(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load double, double* " + value + "\n"
	llgen.Reg++
}

//ScanfInt llvm
func (llgen *LLGenerate) ScanfInt(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @__isoc99_scanf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @strsi, i32 0, i32 0), i32* " + value + ")\n"
	llgen.Reg++
}

//ScanfDouble llvm
func (llgen *LLGenerate) ScanfDouble(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @__isoc99_scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strsf, i32 0, i32 0), double* " + value + ")\n"
	llgen.Reg++
}

//PrintfInt llvm
func (llgen *LLGenerate) PrintfInt(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load i32, i32* " + value + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strpi, i32 0, i32 0), i32 %" + strconv.Itoa(llgen.Reg-1) + ")\n"
	llgen.Reg++
}

//PrintfDouble llvm
func (llgen *LLGenerate) PrintfDouble(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load double, double* " + value + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strpd, i32 0, i32 0), double %" + strconv.Itoa(llgen.Reg-1) + ")\n"
	llgen.Reg++
}

//PrintfString llvm
func (llgen *LLGenerate) PrintfString(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load i8*, i8** %" + "pointer_" + value + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @strpstr, i64 0, i64 0), i8* %" + strconv.Itoa(llgen.Reg-1) + ")\n"
	llgen.Reg++
}

//DeclareInt llvm
func (llgen *LLGenerate) DeclareInt(variable string, isGlobal bool) {
	if isGlobal {
		llgen.HeaderText += "@" + variable + " = global i32 0\n"

	} else {
		llgen.BufferText += "%" + variable + " = alloca i32\n"
	}
}

//DeclareDouble llvm
func (llgen *LLGenerate) DeclareDouble(variable string, isGlobal bool) {
	if isGlobal {
		llgen.HeaderText += "@" + variable + " = global double 0.0\n"
	} else {
		llgen.BufferText += "%" + variable + " = alloca double\n"
	}
}

//DeclareString llvm
// @Depricated
// not used
func (llgen *LLGenerate) DeclareString(variable string) {
	llgen.BufferText += "%" + variable + " = alloca i32\n"

}

//AssignInt llvm
func (llgen *LLGenerate) AssignInt(variable string, value string) {
	llgen.BufferText += "store i32 " + value + ", i32* " + variable + "\n"
}

//AssignDouble llvm
func (llgen *LLGenerate) AssignDouble(variable string, value string) {
	llgen.BufferText += "store double " + value + ", double* " + variable + "\n"
}

//AssignString llvm
func (llgen *LLGenerate) AssignString(variable string, value string, isGlobal bool) {
	v := value[1 : len(value)-1]
	sLen := len(v) + 1
	arrTmp := " [" + strconv.Itoa(sLen) + " x i8]"
	s := arrTmp + " c\"" + v + "\\00\"\n"
	logger.Log.Println("[v]: ", v)
	logger.Log.Println("[s]: ", s)
	var varGlobLoc string
	if isGlobal {
		varGlobLoc = "@" + variable
	} else {
		varGlobLoc = "%" + variable
	}

	llgen.HeaderText += "@.str." + variable + " = constant" + s
	llgen.BufferText += "store i32 0, i32* " + varGlobLoc + "\n"
	llgen.BufferText += "%" + "pointer_" + variable + " = alloca i8*\n"
	llgen.BufferText += "store i8* getelementptr inbounds (" + arrTmp + ", " + arrTmp + "* " + "@.str." + variable + ", i64 0, i64 0), i8** %" + "pointer_" + variable + " \n"

}
