package utils

// PanicIfError - Stop execution flow if errors
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
