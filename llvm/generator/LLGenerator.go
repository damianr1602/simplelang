package llgenerator

import (
	"strconv"
)

type LLGenerate struct {
	HeaderText string
	MainText   string
	Reg        int
}

var LLGenerateInstance = NewLLGenerate()

//NewLLGenerate contructor for LLGenerate
func NewLLGenerate() LLGenerate {
	llInit := LLGenerate{}
	llInit.HeaderText = ""
	llInit.MainText = ""
	llInit.Reg = 1

	return llInit
}

//LoadInt llvm
func (llgen *LLGenerate) LoadInt(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = load i32, i32* %" + value + ")\n"
	llgen.Reg++
}

//LoadDouble llvm
func (llgen *LLGenerate) LoadDouble(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = load double, double* %" + value + ")\n"
	llgen.Reg++
}

//ScanfInt llvm
func (llgen *LLGenerate) ScanfInt(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @__isoc99_scanf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @strsi, i32 0, i32 0), i32* %" + value + ")\n"
	llgen.Reg++
}

//ScanfDouble llvm
func (llgen *LLGenerate) ScanfDouble(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @__isoc99_scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strsf, i32 0, i32 0), double* %" + value + ")\n"
	llgen.Reg++
}

//PrintfInt llvm
func (llgen *LLGenerate) PrintfInt(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = load i32, i32* %" + value + "\n"
	llgen.Reg++
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strpi, i32 0, i32 0), i32 %" + strconv.Itoa(llgen.Reg-1) + ")\n"
	llgen.Reg++
}

//PrintfDouble llvm
func (llgen *LLGenerate) PrintfDouble(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = load double, double* %" + value + "\n"
	llgen.Reg++
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @strpd, i32 0, i32 0), double %" + strconv.Itoa(llgen.Reg-1) + ")\n"
	llgen.Reg++
}

//DeclareInt llvm
func (llgen *LLGenerate) DeclareInt(value string) {
	llgen.MainText += "%" + value + " = alloca i32\n"
}

//DeclareDouble llvm
func (llgen *LLGenerate) DeclareDouble(value string) {
	llgen.MainText += "%" + value + " = alloca double\n"
}

//AssignInt llvm
func (llgen *LLGenerate) AssignInt(variable string, value string) {
	llgen.MainText += "store i32 " + value + ", i32* %" + variable + "\n"
}

//AssignDouble llvm
func (llgen *LLGenerate) AssignDouble(variable string, value string) {
	llgen.MainText += "store double " + value + ", double* %" + variable + "\n"
}

//AddInt llvm
func (llgen *LLGenerate) AddInt(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = add i32 " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

//AddDouble llvm
func (llgen *LLGenerate) AddDouble(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = fadd double " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

//SubInt llvm
func (llgen *LLGenerate) SubInt(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = sub nsw i32 " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

//SubDouble llvm
func (llgen *LLGenerate) SubDouble(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = fsub double " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

//MulInt llvm
func (llgen *LLGenerate) MulInt(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = mul i32 " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

// MulDouble llvm
func (llgen *LLGenerate) MulDouble(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = fmul double " + val1 + ", " + val2 + "\n"
	llgen.Reg++
}

// DivInt llvm
func (llgen *LLGenerate) DivInt(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = sdiv i32 " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

// DivDouble llvm
func (llgen *LLGenerate) DivDouble(val1 string, val2 string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = fdiv double " + val2 + ", " + val1 + "\n"
	llgen.Reg++
}

// Sitofp ToReal llvm
func (llgen *LLGenerate) Sitofp(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = sitofp i32 " + value + " to double\n"
	llgen.Reg++
}

// Fptosi ToInt llvm
func (llgen *LLGenerate) Fptosi(value string) {
	llgen.MainText += "%" + strconv.Itoa(llgen.Reg) + " = fptosi double " + value + " to i32\n"
	llgen.Reg++
}

// Generate code in ll
func (llgen *LLGenerate) Generate() string {
	text := ""
	text += "declare i32 @printf(i8*, ...)\n"
	text += "declare i32 @__isoc99_scanf(i8*, ...)\n"
	text += "@strpi = constant [4 x i8] c\"%d\\0A\\00\"\n"
	text += "@strpd = constant [4 x i8] c\"%f\\0A\\00\"\n"
	text += "@strsi = constant [3 x i8] c\"%d\\00\"\n"
	text += "@strsf = constant [4 x i8] c\"%lf\\00\"\n"
	text += llgen.HeaderText
	text += "define i32 @main() nounwind{\n"
	text += llgen.MainText
	text += "ret i32 0 }\n"
	return text
}
