package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}

type FREM struct {
	base.NoOperandsInstruction
}
type IREM struct {
	base.NoOperandsInstruction
}
type LREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
