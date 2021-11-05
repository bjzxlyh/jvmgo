package constants

import (
	"jvmgo/ch011/instructions/base"
	"jvmgo/ch011/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
