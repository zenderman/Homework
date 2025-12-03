package logger

type IStringer interface {
	String() string
}

type Iloger interface {
	Info(msg IStringer)
	Warn(msg IStringer)
	Fatal(msg IStringer)
	ReadLogs() []string
}
