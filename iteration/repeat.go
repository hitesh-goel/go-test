package iteration

func Repeat(str string, no int) (repeated string) {
	if no == 0 {
		no = 5
	}

	for i := 0; i < no; i++ {
		repeated += str
	}
	return
}
