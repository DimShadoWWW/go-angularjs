package angularjs

func LogErrors() {
	val := recover()

	if val != nil {
		if err, ok := val.(error); ok {
			println("Error in code: ", err.Error())

			return
		}

		println("Panic recovered: ", val)
	}
}
