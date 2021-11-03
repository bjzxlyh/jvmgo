package base

import (
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}
