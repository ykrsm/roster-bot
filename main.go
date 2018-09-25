package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bluele/slack"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hookURL := os.Getenv("WEBHOOK_URL")
	hook := slack.NewWebHook(hookURL)

	xlsx, err := excelize.OpenFile("./data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	year, month, day := time.Now().Date()
	monthInt := int(month)

	fmt.Println(year, monthInt, day)

	sheetIndex := 0
	if monthInt < 7 {
		sheetIndex = 1
	} else {
		sheetIndex = 2
	}

	sheetName := xlsx.GetSheetName(sheetIndex)

	// Get value from cell by given worksheet name and axis.
	// cell := xlsx.GetCellValue(sheetName, "B2")
	// fmt.Println(cell)

	var monthCol int
	var monthRow int
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows(sheetName)
	for rowPos, row := range rows {
		for colPos, colCell := range row {
			if colCell == fmt.Sprintf("%d月", monthInt) {
				monthCol = colPos
				monthRow = rowPos
				fmt.Printf("r: %d, c: %d, v: %s", monthRow, monthCol, colCell)
				fmt.Println()
			}
		}
	}

	// assuming day row always start at
	// Col: monthCol + day - 1
	// Row: monthRow + 2  b.c monthRow is 0 based
	dayColAlpha := excelize.ToAlphaString(monthCol + day - 1)
	dayRow := monthRow + 2
	str := fmt.Sprintf("%s%d", dayColAlpha, dayRow)
	fmt.Printf(str)
	fmt.Println()
	dayVal := xlsx.GetCellValue(sheetName, str)

	fmt.Printf("r: %d, c: %s, v: %s\n", dayRow, dayColAlpha, dayVal)

	/*
		loop names until two empty cell in a row
		name 1
		name 2
		<empty>
		name 3
		<empty>
		<empty> <- until here

		assuming name starts from
		row: monthRow + 2
		col: monthCol - 1
	*/

	// var NAME_COL_ALPHA = excelize.ToAlphaString(monthCol - 1)
	/*
		var nameCol = monthCol - 1
			stop := false
			var textStr string
				for row := monthRow + 2; stop == false; row++ {
					// textStr += rows[row][nameCol]
					textStr += rows[nameCol][row]
					if row == 5 {
						stop = true
					}
				}
	*/
	nameColAlpha := excelize.ToAlphaString(monthCol - 1)
	stop := false
	for row := monthRow + 5; stop == false; row++ {
		strAlpha := fmt.Sprintf("%s%d", nameColAlpha, row)
		fmt.Printf(strAlpha)
		fmt.Println()
		nameVal := xlsx.GetCellValue(sheetName, strAlpha)
		fmt.Println(nameVal)

		rosterIntoRstrAlpha := fmt.Sprintf("%s%d", dayColAlpha, row)
		rosterInfoVal := xlsx.GetCellValue(sheetName, rosterIntoRstrAlpha)
		fmt.Println(rosterInfoVal)

		styleIndex := xlsx.GetCellValue(sheetName, rosterIntoRstrAlpha)
		fmt.Println(styleIndex)

		/*
			switch rosterInfoVal {
			case "D":
			case "D1":
			case "D2":

			default:
			}

		*/
		if row > monthRow+14 {
			stop = true
		}
	}

	err2 := hook.PostMessage(&slack.WebHookPostPayload{

		Text: "hello!",
		Attachments: []*slack.Attachment{
			{Text: "danger", Color: "danger"},
		},
	})
	if err2 != nil {
		panic(err)
	}
}
