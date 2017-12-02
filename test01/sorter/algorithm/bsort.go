package algorithm

/*
Bsort xxx
*/
func Bsort(values []int) {
	fl := true

	for i := 0; i < len(values)-2; i++ {
		fl = true
		for j := 0; j < len(values)-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				fl = false
			}
		}
		if fl == true {
			break
		}
	}
}
