package main

import (
	"fmt"

	lottowizard "github.com/wnjoon/lotto-wizard/src"
)

const GetLottoNumberUrl string = "https://www.dhlottery.co.kr/common.do?method=getLottoNumber&drwNo="

var LottoNumberMap = make(map[int]int)

// var lottoArray = make([]Lotto, 0)

func main() {

	// lottowizard.Write(1111)
	l, err := lottowizard.Read()
	if err != nil {
		panic(err)
	}
	getMostCountedNumber(l)
	getMostUnCountedNumber(l)
}

func getMostCountedNumber(m map[int]int) {
	fmt.Println("가장 많이 당첨된 번호 7개 (보너스 포함)")
	e := lottowizard.SortDescendingOrder(m)
	for i := 0; i < 7; i++ {
		fmt.Println(e[i])
	}
	fmt.Println()
}

func getMostUnCountedNumber(m map[int]int) {
	fmt.Println("가장 적게 당첨된 번호 7개 (보너스 포함)")
	e := lottowizard.SortAscendingOrder(m)
	for i := 0; i < 7; i++ {
		fmt.Println(e[i])
	}
	fmt.Println()
}
