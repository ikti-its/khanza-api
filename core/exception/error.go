package exception

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
