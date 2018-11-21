package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

// const FILE_NAME = "./data.xlsx"

func makeRoster(month, day int, fileName string, roster Roster) (res string) {

	sheetIndex := -1
	if month < 7 {
		sheetIndex = 0
	} else {
		sheetIndex = 1
	}

	monthRow, monthCol := getMonthRowCol(int(month), fileName)

	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// var res string
	sheet := xlFile.Sheets[sheetIndex]
	for r, row := range sheet.Rows {
		if r < monthRow+4 {
			continue
		}

		/*
			These if statements handle when to stop looping (going down the rows)

			TODO: come up with better conditions
			Now it bases on
			* number of Cells (number of cells in a row) is 0
			or
			* FgColor is null
		*/

		size := len(row.Cells)
		fmt.Printf("r: %d size: %d\n", r, size)

		if size == 0 {
			break
		}

		cell := row.Cells[monthCol-1]
		fmt.Printf("cell Fmt: %+v\n", cell.GetStyle().Fill)

		cellFgColor := cell.GetStyle().Fill.FgColor

		if cellFgColor == "" {
			break
		}

		name := cell.String()
		if name != "" &&
			cell.GetStyle().Fill.FgColor == "FFFFE1E1" {
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
		res = ":touban:"
	case "D2":
		res = ":touban:(副)"
	case "D":
		res = ":kinmu:"
	case "R":
		res = ":junbi:"
	case "I":
		res = ":idou:"
	case "T":
		res = ":syuttyou:"
	case "V":
		res = ":kyuujitu:"
	default:
		res = ":kyuujitu:"
	}
	return
}

// month is 1 based
// January is 1
// February is 2
// and so on...
func getMonthRowCol(month int, fileName string) (monthRow, monthCol int) {
	xlsx1, err := excelize.OpenFile(fileName)
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
