package errors_util

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
