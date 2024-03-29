package tools

import (
	gds "absolut-music/src/globalDataStructures"
	"absolut-music/src/structures"
	"sync"
)

/*Reverse the Artists slice*/
func ReverseSliceArtist(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(gds.Artists)/2; i++ {
		gds.Artists[i], gds.Artists[len(gds.Artists)-1-i] = gds.Artists[len(gds.Artists)-1-i], gds.Artists[i]
	}
}

/*
Take the artist in the middle and use it as a pivot
loop through the table t and check if the artist is under the pivot
using the sortingOption as a condition. If the artist is under the pivot,
it's appended to the sliceBefore, if not, it's appended to the slice after
the sliceBefore, sliceAfter and the pivot are returned at the end
*/
func partition(t []structures.Artist, sortingOption string) ([]structures.Artist, structures.Artist, []structures.Artist) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	sliceBefore := []structures.Artist{}
	sliceAfter := []structures.Artist{}
	for i := 0; i < len(slice); i++ {
		switch sortingOption {
		case "name":
			if t[i].Name < pivot.Name {
				sliceBefore = append(sliceBefore, t[i])
			} else {
				sliceAfter = append(sliceAfter, t[i])
			}
		case "creationDate":
			if t[i].CreationDate < pivot.CreationDate {
				sliceBefore = append(sliceBefore, t[i])
			} else {
				sliceAfter = append(sliceAfter, t[i])
			}
		case "numberMembers":
			if len(t[i].Members) < len(pivot.Members) {
				sliceBefore = append(sliceBefore, t[i])
			} else {
				sliceAfter = append(sliceAfter, t[i])
			}
		case "Firstalbumrelease":
			formatedPivot := ParseDate(pivot.FirstAlbum)
			formatedDateT := ParseDate(t[i].FirstAlbum)
			if formatedDateT.Year < formatedPivot.Year {
				sliceBefore = append(sliceBefore, t[i])
			} else if formatedDateT.Year == formatedPivot.Year {
				if formatedDateT.Month < formatedPivot.Month {
					sliceBefore = append(sliceBefore, t[i])
				} else if formatedDateT.Month == formatedPivot.Month {
					if formatedDateT.Day < formatedPivot.Day {
						sliceBefore = append(sliceBefore, t[i])
					} else {
						sliceAfter = append(sliceAfter, t[i])
					}
				} else {
					sliceAfter = append(sliceAfter, t[i])
				}
			} else {
				sliceAfter = append(sliceAfter, t[i])
			}
		}

	}
	return sliceBefore, pivot, sliceAfter
}

/*Merge the sliceBefore to the pivot to the sliceAfter*/
func merge(sB []structures.Artist, p structures.Artist, sA []structures.Artist) []structures.Artist {
	sB = append(sB, p)
	sB = append(sB, sA...)
	return sB
}

/*
Handle the recursiv calls for the sort using the
Divide and rule policy
*/
func quickSortControler(t []structures.Artist, sortingOption string) []structures.Artist {
	if len(t) < 1 {
		return t
	}
	s1, p, s2 := partition(t, sortingOption)
	a := merge(quickSortControler(s1, sortingOption), p, quickSortControler(s2, sortingOption))
	return a
}

/*
Call the quickSortControler function and copy
the result to the slice Artist wich will be overwritten by the result
*/
func QuickSort(sortingOption string, asc bool) {
	copy(gds.Artists, quickSortControler(gds.Artists, sortingOption))
	if !asc {
		RunParallel(ReverseSliceArtist)
	}
}
