package constants

import (
	"jvmgo/ch010/instructions/base"
	"jvmgo/ch010/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
