package ehelper

type ErrorHandler interface {
	CheckError()
}

func (e Ehelper) CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
