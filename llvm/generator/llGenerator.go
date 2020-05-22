package llgenerator

type LLGenerate struct {
	HeaderText string
	MainText   string
	BufferText string
	MainReg    int
	Reg        int
	Br         int
}

//NewLLGenerate contructor for LLGenerate
func NewLLGenerate() LLGenerate {
	llInit := LLGenerate{}
	llInit.HeaderText = ""
	llInit.MainText = ""
	llInit.BufferText = ""
	llInit.MainReg = 1
	llInit.Reg = 1
	llInit.Reg = 1

	return llInit
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
