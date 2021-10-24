package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ISHL struct {
	base.NoOperandsInstruction //int左移位
}

type ISHR struct {
	base.NoOperandsInstruction //int算术右移位
}

type IUSHR struct {
	base.NoOperandsInstruction //int逻辑右移位
}

type LSHL struct {
	base.NoOperandsInstruction //long左移位
}

type LSHR struct {
	base.NoOperandsInstruction //long算术右移位
}

type LUSHR struct {
	base.NoOperandsInstruction //long逻辑右移位
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
