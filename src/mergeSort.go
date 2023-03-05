package src

func Partition(t []int) ([]int, int, []int) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	sliceBefore := []int{}
	sliceAfter := []int{}
	for i := 0; i < len(slice); i++ {
		if slice[i] < pivot {
			sliceBefore = append(sliceBefore, slice[i])
		} else {
			sliceAfter = append(sliceAfter, slice[i])
		}
	}
	return sliceBefore, pivot, sliceAfter
}

func Merge(sB []int, p int, sA []int) []int {
	sB = append(sB, p)
	sB = append(sB, sA...)
	return sB
}


func QuickSortControler(t []int) []int {
	if len(t) < 1 {
		return t
	}
	s_1, p, s_2 := Partition(t)
	a := Merge(QuickSortControler(s_1), p, QuickSortControler(s_2))
	return a
}

func QuickSort(t []int) {
	z := QuickSortControler(t)
	copy(t, z)
}
