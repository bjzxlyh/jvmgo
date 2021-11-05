package rtda

type Stack struct {
	maxSize uint   //保存栈的容量
	size    uint   //保存栈的当前大小
	_top    *Frame //保存栈顶指针
}

//newStack()函数创建Stack结构体实例，参数表示要创建的Stack最多可以容纳多少帧
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

//push()方法把栈推入栈顶
func (self *Stack) push(frame *Frame) {
	//如果栈满了，抛出异常
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError!")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
func (self *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0, self.size)
	for frame := self._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}
func (self *Stack) isEmpty() bool {
	return self._top == nil
}
func (self *Stack) clear() {
	for !self.isEmpty() {
		self.pop()
	}
}
