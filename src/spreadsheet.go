package lottowizard

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Write(round int) error {
	f, err := excelize.OpenFile("lotto.xlsx")
	if err != nil {
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// cell, err := f.GetCellValue("Lotto", fmt.Sprint("A", stringfy(round+1)))
	// if err != nil {
	// 	return err
	// }

	// if cell == fmt.Sprint("A", stringfy(round)) {
	// 	return errors.New("round is already inserted")
	// }

	lotto, err := getLottoNumberByRound(round)
	if err != nil {
		return err
	}

	f.SetCellValue("Lotto", fmt.Sprint("A", stringfy(round+1)), round)
	f.SetCellValue("Lotto", fmt.Sprint("B", stringfy(round+1)), lotto.DrwtNo1)
	f.SetCellValue("Lotto", fmt.Sprint("C", stringfy(round+1)), lotto.DrwtNo2)
	f.SetCellValue("Lotto", fmt.Sprint("D", stringfy(round+1)), lotto.DrwtNo3)
	f.SetCellValue("Lotto", fmt.Sprint("E", stringfy(round+1)), lotto.DrwtNo4)
	f.SetCellValue("Lotto", fmt.Sprint("F", stringfy(round+1)), lotto.DrwtNo5)
	f.SetCellValue("Lotto", fmt.Sprint("G", stringfy(round+1)), lotto.DrwtNo6)
	f.SetCellValue("Lotto", fmt.Sprint("H", stringfy(round+1)), lotto.BnusNo)

	return nil
}

func Read() (map[int]int, error) {
	f, err := excelize.OpenFile("lotto.xlsx")
	if err != nil {
		return nil, err
	}
	defer func() error {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()

	lottoMap := make(map[int]int)

	round, err := getLatestRoundNumber()
	if err != nil {
		return nil, err
	}

	// if len(rows) <= round {
	// err = Write(round)
	// if err != nil {
	// 	return nil, err
	// }
	// }
	lotto, err := getLottoNumberByRound(round)
	if err != nil {
		return nil, err
	}

	f.SetCellValue("Lotto", fmt.Sprint("A", stringfy(round+1)), round)
	f.SetCellValue("Lotto", fmt.Sprint("B", stringfy(round+1)), lotto.DrwtNo1)
	f.SetCellValue("Lotto", fmt.Sprint("C", stringfy(round+1)), lotto.DrwtNo2)
	f.SetCellValue("Lotto", fmt.Sprint("D", stringfy(round+1)), lotto.DrwtNo3)
	f.SetCellValue("Lotto", fmt.Sprint("E", stringfy(round+1)), lotto.DrwtNo4)
	f.SetCellValue("Lotto", fmt.Sprint("F", stringfy(round+1)), lotto.DrwtNo5)
	f.SetCellValue("Lotto", fmt.Sprint("G", stringfy(round+1)), lotto.DrwtNo6)
	f.SetCellValue("Lotto", fmt.Sprint("H", stringfy(round+1)), lotto.BnusNo)

	rows, err := f.GetRows("Lotto")
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(rows); i++ {
		for j := 1; j < 8; j++ {
			n, err := strconv.Atoi(rows[i][j])
			if err != nil {
				return nil, err
			}
			lottoMap[n] = lottoMap[n] + 1
		}
	}
	return lottoMap, nil
}

func stringfy(n int) string {
	return strconv.FormatInt(int64(n), 10)
}
