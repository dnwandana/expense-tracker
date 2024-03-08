package utils

// PanicIfError panics if the given error is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
