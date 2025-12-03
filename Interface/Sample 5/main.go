package main

import (
	"Interface/logger"
	"fmt"
	"time"
)

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

type LogMessage struct {
	UserId    string
	PaymentId string
	msg       string
}

func (l *LogMessage) String() string {

	l.msg = fmt.Sprintf("%s : PaymentId: %s, UserId: %s", time.Now(), l.PaymentId, l.UserId)

	return l.msg
}

func main() {

	var l logger.Iloger = &logger.LoggerApi{}

	l.Info(&LogMessage{"123", "500rub", "User paid"})
	l.Warn(&LogMessage{"555", "505rub", "User paid"})

	fmt.Println()

	var b IReaderWriter = &MemoryBuffer{}
	b.Write("Hello kek")

	fmt.Println(b.Read())

	for _, s := range l.ReadLogs() {
		fmt.Println(s)
	}

}
