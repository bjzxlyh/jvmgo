package rtda

import "jvmgo/ch06/rtda/heap"

type Frame struct {
	lower        *Frame        //实现链表
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //保存操作数栈指针
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
		method:       method,
	}
}

//getters

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
