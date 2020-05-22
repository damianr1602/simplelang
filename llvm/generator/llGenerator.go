package llgenerator

import (
	"strconv"

	"github.com/damianr1602/simplelang/util"

	logger "github.com/damianr1602/simplelang/logging"
)

type LLGenerate struct {
	HeaderText string
	MainText   string
	BufferText string
	MainReg    int
	Reg        int
	Br         int
}

var brstack util.Stack
var br int = 0

//NewLLGenerate contructor for LLGenerate
func NewLLGenerate() LLGenerate {
	llInit := LLGenerate{}
	llInit.HeaderText = ""
	llInit.MainText = ""
	llInit.BufferText = ""
	llInit.MainReg = 1
	llInit.Reg = 1

	return llInit
}

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

//AddInt llvm
func (llgen *LLGenerate) AddInt(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = add i32 " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

//AddDouble llvm
func (llgen *LLGenerate) AddDouble(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = fadd double " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

//SubInt llvm
func (llgen *LLGenerate) SubInt(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = sub nsw i32 " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

//SubDouble llvm
func (llgen *LLGenerate) SubDouble(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = fsub double " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

//MulInt llvm
func (llgen *LLGenerate) MulInt(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = mul i32 " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

// MulDouble llvm
func (llgen *LLGenerate) MulDouble(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = fmul double " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

// DivInt llvm
func (llgen *LLGenerate) DivInt(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = sdiv i32 " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

// DivDouble llvm
func (llgen *LLGenerate) DivDouble(val1 string, val2 string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = fdiv double " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

// Sitofp ToReal llvm
func (llgen *LLGenerate) Sitofp(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = sitofp i32 " + value + " to double\n"
	llgen.Reg++
}

// Fptosi ToInt llvm
func (llgen *LLGenerate) Fptosi(value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = fptosi double " + value + " to i32\n"
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

// BlockIfEnter llvm
func (llgen *LLGenerate) BlockIfEnter() {
	br++
	llgen.BufferText += "br i1 %" + strconv.Itoa(llgen.Reg-1) + ", label %true" + strconv.Itoa(br) + ", label %false" + strconv.Itoa(br) + "\n"
	llgen.BufferText += "true" + strconv.Itoa(br) + ":\n"
	brstack.Push(br)

}

// BlockIfExit llvm
func (llgen *LLGenerate) BlockIfExit() {
	b := brstack.Pop().(int)
	logger.Log.Println("br stack popped: ", b)
	llgen.BufferText += "br label %false" + strconv.Itoa(b) + "\n"
	llgen.BufferText += "false" + strconv.Itoa(b) + ":\n"

}

// Icmp llvm
func (llgen *LLGenerate) Icmp(id string, value string) {
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = icmp eq i32 %" + strconv.Itoa(llgen.Reg-1) + ", " + value + "\n"
	llgen.Reg++

}

// StartLoop llvm
func (llgen *LLGenerate) StartLoop(repetitions string) {
	llgen.DeclareInt(strconv.Itoa(llgen.Reg), false)
	counter := llgen.Reg
	coutnerString := strconv.Itoa(counter)
	llgen.Reg++
	llgen.AssignInt("%"+coutnerString, "0")
	br++
	llgen.BufferText += "br label %cond" + strconv.Itoa(br) + "\n"
	llgen.BufferText += "cond" + strconv.Itoa(br) + ":\n"
	llgen.LoadInt("%" + coutnerString)
	llgen.AddInt("%"+strconv.Itoa(llgen.Reg-1), "1")
	llgen.AssignInt("%"+coutnerString, "%"+strconv.Itoa(llgen.Reg-1))
	llgen.LoadInt(repetitions)
	llgen.BufferText += "%" + strconv.Itoa(llgen.Reg) + " = icmp slt i32 %" + strconv.Itoa(llgen.Reg-3) + ", %" + strconv.Itoa(llgen.Reg-1) + "\n"
	llgen.Reg++
	llgen.BufferText += "br i1 %" + strconv.Itoa(llgen.Reg-1) + ", label %true" + strconv.Itoa(br) + ", label %false" + strconv.Itoa(br) + "\n"
	llgen.BufferText += "true" + strconv.Itoa(br) + ":\n"
	brstack.Push(br)
	logger.Log.Println(brstack, "stack after repeat block: ", br)

}

// EndLoop impl
func (llgen *LLGenerate) EndLoop() {
	b := brstack.Pop().(int)
	logger.Log.Println("br stack popped: ", b)
	llgen.BufferText += "br label %cond" + strconv.Itoa(b) + "\n"
	llgen.BufferText += "false" + strconv.Itoa(b) + ":\n"
}

// StopMain llvm
func (llgen *LLGenerate) StopMain() {
	llgen.MainText += llgen.BufferText
	logger.Log.Println("Stop Main")
}

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

// Generate code in ll
func (llgen *LLGenerate) Generate() string {
	text := ""
	text += llgen.HeaderText
	text += "declare i32 @printf(i8*, ...)\n"
	text += "declare i32 @__isoc99_scanf(i8*, ...)\n"
	text += "@strpi = constant [4 x i8] c\"%d\\0A\\00\"\n"
	text += "@strpd = constant [4 x i8] c\"%f\\0A\\00\"\n"
	text += "@strsi = constant [3 x i8] c\"%d\\00\"\n"
	text += "@strsf = constant [4 x i8] c\"%lf\\00\"\n"
	text += "@strpstr = constant [3 x i8] c\"%s\\00\"\n"
	text += "define i32 @main() nounwind{\n"
	text += llgen.MainText
	text += "ret i32 0 }\n"
	text += "declare void @llvm.memcpy.p0i8.p0i8.i64(i8*, i8*, i64, i1 immarg) nounwind willreturn\n"
	return text
}
