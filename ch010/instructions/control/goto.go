package control

import (
	"jvmgo/ch010/instructions/base"
	"jvmgo/ch010/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
