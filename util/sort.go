package lottowizard

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
