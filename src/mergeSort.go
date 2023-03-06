package src

func Partition(t []Artist, sortingOption string) ([]Artist, Artist, []Artist) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	sliceBefore := []Artist{}
	sliceAfter := []Artist{}
	for i := 0; i < len(slice); i++ {
		if sortingOption == "name" {
			if t[i].Name < pivot.Name {
				sliceBefore = append(sliceBefore, slice[i])
			} else {
				sliceAfter = append(sliceAfter, slice[i])
			}
		} else if sortingOption == "creationDate" {
			if t[i].CreationDate < pivot.CreationDate {
				sliceBefore = append(sliceBefore, slice[i])
			} else {
				sliceAfter = append(sliceAfter, slice[i])
			}
		} else if sortingOption == "numberMembers" {
			if len(t[i].Members) < len(pivot.Members) {
				sliceBefore = append(sliceBefore, slice[i])
			} else {
				sliceAfter = append(sliceAfter, slice[i])
			}
		} else if sortingOption == "Firstalbumrelease" {
			SortFirstAlbum()
		}
	}
	return sliceBefore, pivot, sliceAfter
}

func Merge(sB []Artist, p Artist, sA []Artist) []Artist {
	sB = append(sB, p)
	sB = append(sB, sA...)
	return sB
}

func QuickSortControler(t []Artist, sortingOption string) []Artist {
	if len(t) < 1 {
		return t
	}
	s1, p, s2 := Partition(t, sortingOption)
	a := Merge(QuickSortControler(s1, sortingOption), p, QuickSortControler(s2, sortingOption))
	return a
}

func QuickSort(sortingOption string, asc bool) {
	z := QuickSortControler(Artists, sortingOption)
	copy(Artists, z)
	if !asc {
		RunParallel(reverseSliceArtist)
	}
}
