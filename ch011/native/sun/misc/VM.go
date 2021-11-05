package misc

import (
	"jvmgo/ch011/instructions/base"
	"jvmgo/ch011/native"
	"jvmgo/ch011/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) { // hack: just make VM.savedProps nonempty
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
