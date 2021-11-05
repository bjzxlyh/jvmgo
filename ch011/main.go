package main

import (
	"fmt"
)

import _ "jvmgo/ch011/classfile"
import _ "jvmgo/ch011/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd)
	}
}
