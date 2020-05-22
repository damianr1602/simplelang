package llgenerator

import (
	"strconv"
)

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
