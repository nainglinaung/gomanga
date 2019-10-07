package ehelper

import "fmt"

type LogHandler interface {
	Log()
}

func (e Ehelper) Log(message string) {
	fmt.Println(message)
}
