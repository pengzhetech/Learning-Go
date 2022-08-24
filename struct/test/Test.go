package test

type TestsInterface interface {
	Hello(s string) string
	To(ss int) int
}

var app TestsStruct

type TestsStruct struct{}

func (t *TestsStruct) Hello(s string) string {
	return s + "----------->"
}

func (t *TestsStruct) To(ss int) int {
	return ss
}

func Instance() TestsInterface {
	return &app
}
