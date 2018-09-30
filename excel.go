package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

const FILE_NAME = "./data.xlsx"

func makeRoster(month, day int) string {

	sheetIndex := 0
	if month < 7 {
		sheetIndex = 1
	} else {
		sheetIndex = 2
	}

	monthRow, monthCol := getMonthRowCol(int(month))

	xlFile, err := xlsx.OpenFile(FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}

	var res string
	for s, sheet := range xlFile.Sheets {
		if s != sheetIndex-1 {
			continue
		}
		for r, row := range sheet.Rows {
			if r < monthRow+4 {
				continue
			}
			for c, cell := range row.Cells {
				if cell.String() != "" &&
					cell.GetStyle().Fill.FgColor == "FFFFE1E1" &&
					c == monthCol-1 {
					name := cell.String()
					fmt.Printf("%s\n", name)

					// get today work info
					workInfoCell := row.Cells[monthCol+day]
					workInfo := row.Cells[monthCol+day].String()
					fmt.Printf("%s %s\n", workInfo, workInfoCell.GetStyle().Fill.FgColor)

					if workInfo == "D" && workInfoCell.GetStyle().Fill.FgColor == "" {
						workInfo = "D1"
					}

					workInfoStr := makeTodayInfoStr(workInfo)

					if workInfoStr == "" {
						break
					}
					res = res + name + ": \t" + workInfoStr + "\n"
					break
				}
			}
		}
	}
	return res
}

func makeTodayInfoStr(workInfo string) (res string) {
	switch workInfo {
	case "D1":
		res = "当番"
	case "D2":
		res = "当番(副)"
	case "D":
		res = "通常勤務"
	case "R":
		res = "準備期間"
	case "I":
		res = "出張移動"
	case "T":
		res = "出張"
	}
	return
}

// month is 1 based
// January is 1
// February is 2
// and so on...
func getMonthRowCol(month int) (monthRow, monthCol int) {
	xlsx1, err := excelize.OpenFile(FILE_NAME)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	monthInt := int(month)

	sheetIndex := 0
	if monthInt < 7 {
		sheetIndex = 1
	} else {
		sheetIndex = 2
	}

	sheetName := xlsx1.GetSheetName(sheetIndex)

	// get month row and col
	rows := xlsx1.GetRows(sheetName)
	for rowPos, row := range rows {
		for colPos, colCell := range row {
			if colCell == fmt.Sprintf("%d月", monthInt) {
				monthCol = colPos
				monthRow = rowPos
				fmt.Printf("Found month:\trow: %d, col: %d, val: %s\n", monthRow, monthCol, colCell)
			}
		}
	}
	return
}
