package lottowizard

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Write(r int) {
	f, err := excelize.OpenFile("lotto.xlsx")
	if err != nil {
		//
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	cell, err := f.GetCellValue("Lotto", fmt.Sprint("A", strconv.FormatInt(int64(r+1), 10)))
	if err != nil {
		// fmt.Println("No Cell is available")
		return
	}
	if cell == "" {
		fmt.Println("No cell is available")
	}
	fmt.Println(cell)

	// f.SetCellValue("Lotto", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
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
