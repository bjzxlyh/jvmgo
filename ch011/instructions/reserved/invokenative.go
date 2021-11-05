package reserved

import (
	"jvmgo/ch011/instructions/base"
	"jvmgo/ch011/native"
	"jvmgo/ch011/rtda"

	_ "jvmgo/ch010/native/java/lang"
	_ "jvmgo/ch010/native/sun/misc"
)

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
