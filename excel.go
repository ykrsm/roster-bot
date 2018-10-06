package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

const FILE_NAME = "./data.xlsx"

func makeRoster(month, day int) (res string) {

	sheetIndex := -1
	if month < 7 {
		sheetIndex = 0
	} else {
		sheetIndex = 1
	}

	monthRow, monthCol := getMonthRowCol(int(month))

	xlFile, err := xlsx.OpenFile(FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}

	// var res string
	sheet := xlFile.Sheets[sheetIndex]
	for r, row := range sheet.Rows {
		if r < monthRow+4 {
			continue
		}
		size := len(row.Cells)
		fmt.Printf("size: %s\n", size)
		if size == 0 {
			break
		}
		cell := row.Cells[monthCol-1]

		if cell.String() != "" &&
			cell.GetStyle().Fill.FgColor == "FFFFE1E1" {
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

			if workInfoStr != "" {
				res = res + name + ": \t" + workInfoStr + "\n"
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
