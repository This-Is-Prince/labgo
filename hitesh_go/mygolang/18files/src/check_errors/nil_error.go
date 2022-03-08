package check_errors

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
