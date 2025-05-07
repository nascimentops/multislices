package multislices

//função que recebe duas slices e retorna a multiplicação de seus elementos
func MultiElements(a, b []int) []int {
	var lenght int

	if len(a) > len(b) {
		lenght = len(a)
		for len(b) < lenght {
			b = append(b, -1)
		}
	} else {
		lenght = len(b)
		for len(a) < lenght {
			a = append(a, -1)
		}
	}

	var slice []int = make([]int, lenght)

	for i := 0; i < lenght; i++ {
		if a[i] >= 0 && b[i] >= 0 {
			slice[i] = a[i] * b[i]
		}
	}

	return slice
}
