package main

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func fillRoster(year, month, day int, fileName string, roster Roster, names []string) (res Roster) {

	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var halfYearStr string
	if month < 7 {
		halfYearStr = "上半期"
	} else {
		halfYearStr = "下半期"
	}

	sheetKey := fmt.Sprintf("%d年%s", year, halfYearStr)

	monthRow, monthCol := getMonthRowCol(month, sheetKey, xlFile)

	empSlice := []Employee{}

	numOfNames := len(names)

	sheet := xlFile.Sheet[sheetKey]
	for r, row := range sheet.Rows {
		if r > monthRow+numOfNames+3 {
			break
		}

		if r < monthRow+4 {
			continue
		}

		size := len(row.Cells)

		if size == 0 {
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

		cell := row.Cells[monthCol-1]

		name := cell.String()
		fmt.Printf("%+v\n", name)

		isNameMatch := contains(names, name)
		fmt.Printf("%+v\n", isNameMatch)

		if !isNameMatch {
			continue
		}

		// cellFgColor := cell.GetStyle().Fill.FgColor

		/*
			if cellFgColor != "FFFFE1E1" {
				continue
			}
		*/
		if name == "" {
			continue
		}

		// get today work info
		workInfoCell := row.Cells[monthCol+day]
		workInfo := row.Cells[monthCol+day].String()

		if workInfo == "D" && workInfoCell.GetStyle().Fill.FgColor == "" {
			workInfo = "D1"
		}

		workInfoObj := makeWorkInfo(workInfo)

		emp := Employee{name, workInfoObj}
		empSlice = append(empSlice, emp)
		fmt.Printf("%+v\n", emp)
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
	case "S":
		res = On
	case "DR":
		res = Remote
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
func getMonthRowCol(month int, sheetKey string, xlFile *xlsx.File) (monthRow, monthCol int) {

	fmt.Printf("sheet index %s\n", sheetKey)

	sheet := xlFile.Sheet[sheetKey]
	for rowPos, row := range sheet.Rows {
		for colPos, colCell := range row.Cells {
			if colCell.String() == fmt.Sprintf("%d月", month) {
				monthCol = colPos
				monthRow = rowPos
			}
		}
	}

	return
}

// return 0 based sheet index
func makeSheetName(year, month int) string {

	var halfYearStr string
	if month < 7 {
		halfYearStr = "上半期"
	} else {
		halfYearStr = "下半期"
	}

	return fmt.Sprintf("%d年%s", year, halfYearStr)
}

func contains(names []string, target string) bool {
	for _, name := range names {
		if strings.Contains(target, name) {
			return true
		}
	}
	return false
}
