package test

func Test() {

	// array = 2, 1
	a := 1
	b := 2

	if bite(a, b) {
		println("bite")
	} else {
		println("no bite")
	}
}

func bite(a int, b int) bool {
	if a < b && (a+3) >= b {
		return true
	}
	return false
}
