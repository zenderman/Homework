package main

import "fmt"

type IReader interface {
	Read() string
}

type IWriter interface {
	Write(s string)
}

type IReaderWriter interface {
	IReader
	IWriter
}

type MemoryBuffer struct {
	data string
}

func (m *MemoryBuffer) Read() string {
	return m.data
}

func (m *MemoryBuffer) Write(s string) {
	m.data = s
}

func main() {
	var b IReaderWriter = &MemoryBuffer{}
	b.Write("Hello kek")

	fmt.Println(b.Read())

	_ = b
}
