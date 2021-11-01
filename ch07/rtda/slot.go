package rtda

import "jvmgo/ch07/rtda/heap"

type Slot struct {
	num int32        //存放整数
	ref *heap.Object //存放引用
}
