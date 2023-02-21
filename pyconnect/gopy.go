package main

import (
	"fmt"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
	sysModule := python.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")          // path = sys.path
	python.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
	return python.PyImport_ImportModule(name)        // return __import__(name)
}

func main() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
	// import hello
	hello := ImportModule("G:/crypto/pyconnect/pypackage", "hello")
	fmt.Printf("[MODULE] repr(hello) = %s\n", GoStr(hello.Repr()))

	// // print(hello.a)
	// a := hello.GetAttrString("a")
	// fmt.Printf("[VARS] a = %#v\n", python.PyInt_AsLong(a))
}
