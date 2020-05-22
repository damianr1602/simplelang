package llgenerator

import (
	"strconv"

	logger "github.com/damianr1602/simplelang/logging"
)

// PrintfArrElemInt llvm
func (llgen *LLGenerate) PrintfArrElemInt(value string, valueElem string, arrayLen string) {
	arrStr := " [" + arrayLen + " x i32]"
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = getelementptr inbounds" + arrStr + "," + arrStr + "* %" + value + ", i64 0, i64 " + valueElem + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load i32, i32* %" + strconv.Itoa(llgen.Reg-1) + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strpi, i32 0, i32 0), i32 %" + strconv.Itoa(llgen.Reg-1) + ")\n"
	llgen.Reg++
}

// DeclareIntArray ToInt llvm
func (llgen *LLGenerate) DeclareIntArray(variable string, values []string) {
	var s, sArrStr string
	sLen := len(values)
	sArrStr += " [" + strconv.Itoa(sLen) + " x i32]"
	s += sArrStr + " ["
	for i, v := range values {
		logger.Log.Println("[val]: ", v, ", index: ", i)

		s += "i32 " + v + ", "
	}

	s = s[:len(s)-2]

	logger.Log.Println("[s]: ", s)

	llgen.HeaderText += "@__const.main." + variable + " = private unnamed_addr constant " + s + "]\n"

	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = alloca i32\n"
	llgen.BufferText += "store i32 0" + ", i32* %" + strconv.Itoa(llgen.Reg) + "\n"
	llgen.BufferText += "%" + variable + " = alloca " + sArrStr + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = bitcast " + sArrStr + "* " + "%" + variable + " to i8*\n"
	llgen.BufferText += "call void @llvm.memcpy.p0i8.p0i8.i64(i8* %" + strconv.Itoa(llgen.Reg) + ", i8* bitcast (" + sArrStr + "* @__const.main." + variable + " to i8*), i64 " + strconv.Itoa(sLen*4) + ", i1 false)\n"
	llgen.Reg++
}

// AssignArrElemInt llvm
func (llgen *LLGenerate) AssignArrElemInt(variable string, valueElem string, arrayLen string) {
	arrStr := " [" + arrayLen + " x i32]"
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = getelementptr inbounds" + arrStr + "," + arrStr + "* %" + variable + ", i64 0, i64 " + valueElem + "\n"
	llgen.Reg++
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = load i32, i32* %" + strconv.Itoa(llgen.Reg-1) + "\n"
	llgen.Reg++
}

// DeclareArrElemInt llvm
func (llgen *LLGenerate) DeclareArrElemInt(variable string) {
	llgen.BufferText += "%" + variable + " = alloca i32\n"
	llgen.BufferText += "store i32 %" + strconv.Itoa(llgen.Reg-1) + ", i32* %" + variable + "\n"
}
