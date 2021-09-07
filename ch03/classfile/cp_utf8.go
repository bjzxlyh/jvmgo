package classfile

type ConstantUtf8Info struct {
	str string
}

/*
readInfo（）方法先读取出[]byte，然后调用decodeMUTF8（）函数把它解码成Go字符串
*/
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
