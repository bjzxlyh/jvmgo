package control

import (
	"jvmgo/ch011/instructions/base"
	"jvmgo/ch011/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
