package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

func fillRoster(month, day int, fileName string, roster Roster) (res Roster) {

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

	empSlice := []Employee{}

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

		if size == 0 {
			break
		}

		cell := row.Cells[monthCol-1]

		cellFgColor := cell.GetStyle().Fill.FgColor

		if cellFgColor == "" {
			break
		}

		name := cell.String()
		if name != "" &&
			cell.GetStyle().Fill.FgColor == "FFFFE1E1" {

			// get today work info
			workInfoCell := row.Cells[monthCol+day]
			workInfo := row.Cells[monthCol+day].String()

			if workInfo == "D" && workInfoCell.GetStyle().Fill.FgColor == "" {
				workInfo = "D1"
			}

			workInfoObj := makeWorkInfo(workInfo)

			emp := Employee{name, workInfoObj}
			empSlice = append(empSlice, emp)
		}
	}

	roster.Employees = empSlice
	return roster

}

func makeWorkInfo(workInfo string) (res WorkInfo) {
	switch workInfo {
	case "D1":
		res = Duty
	case "D2":
		res = SubDuty
	case "D":
		res = On
	case "R":
		res = Prepare
	case "I":
		res = Moving
	case "T":
		res = Trip
	case "V":
		res = Off
	default:
		res = Off
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
			}
		}
	}
	return
}
