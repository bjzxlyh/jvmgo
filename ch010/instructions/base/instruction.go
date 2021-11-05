package base

import "jvmgo/ch010/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) //从字节码中提取操作数
	Execute(frame *rtda.Frame)            //执行指令逻辑
}

type NoOperandsInstruction struct { //表示没有操作数的指令

}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

type BranchInstruction struct { //表示跳转指令
	Offset int //存放跳转偏移量
}

//从字节码中读取一个uint16整数，再转化成int

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint //局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
