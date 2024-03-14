package lottowizard

import "sort"

type Elem struct {
	Number int
	Count  int
}

type ElemTable []Elem

func (s ElemTable) Len() int {
	return len(s)
}

func (s ElemTable) Less(i, j int) bool {
	return s[i].Count < s[j].Count
}

func (s ElemTable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortDescendingOrder(m map[int]int) ElemTable {
	sorted := make(ElemTable, 0)

	for number, count := range m {
		sorted = append(sorted, Elem{number, count})
	}
	sort.Sort(sort.Reverse(sorted))

	return sorted
}

func SortAscendingOrder(m map[int]int) ElemTable {
	sorted := make(ElemTable, 0)

	for number, count := range m {
		sorted = append(sorted, Elem{number, count})
	}
	sort.Sort(sorted)

	return sorted
}
