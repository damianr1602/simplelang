package llgenerator

import (
	"strconv"

	logger "github.com/damianr1602/simplelang/logging"
	"github.com/damianr1602/simplelang/util"
)

var brstack util.Stack
var br int = 0

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
