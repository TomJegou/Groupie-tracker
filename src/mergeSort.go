package src

import (
	"fmt"
	"strconv"
	"strings"
)

type FormatDate struct {
	Year  int
	Month int
	Day   int
}

func parseDate(date string) FormatDate {
	t := strings.Split(date, "-")
	yearInt, err := strconv.Atoi(t[2])
	if err != nil {
		fmt.Println(err)
	}
	monthInt, err := strconv.Atoi(t[1])
	if err != nil {
		fmt.Println(err)
	}
	dayInt, err := strconv.Atoi(t[0])
	if err != nil {
		fmt.Println(err)
	}
	parsedDate := FormatDate{Year: yearInt, Month: monthInt, Day: dayInt}
	return parsedDate
}

func partition(t []Artist, sortingOption string) ([]Artist, Artist, []Artist) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	sliceBefore := []Artist{}
	sliceAfter := []Artist{}
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
			formatedPivot := parseDate(pivot.FirstAlbum)
			formatedDateT := parseDate(t[i].FirstAlbum)
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

func merge(sB []Artist, p Artist, sA []Artist) []Artist {
	sB = append(sB, p)
	sB = append(sB, sA...)
	return sB
}

func quickSortControler(t []Artist, sortingOption string) []Artist {
	if len(t) < 1 {
		return t
	}
	s1, p, s2 := partition(t, sortingOption)
	a := merge(quickSortControler(s1, sortingOption), p, quickSortControler(s2, sortingOption))
	return a
}

func QuickSort(sortingOption string, asc bool) {
	z := quickSortControler(Artists, sortingOption)
	copy(Artists, z)
	if !asc {
		RunParallel(reverseSliceArtist)
	}
}
