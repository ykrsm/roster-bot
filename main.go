package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bluele/slack"
	"github.com/joho/godotenv"
	"github.com/tealeg/xlsx"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hookURL := os.Getenv("WEBHOOK_URL")
	hook := slack.NewWebHook(hookURL)

	xlsx1, err := excelize.OpenFile("./data.xlsx")
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

	sheetName := xlsx1.GetSheetName(sheetIndex)

	// Get value from cell by given worksheet name and axis.
	// cell := xlsx.GetCellValue(sheetName, "B2")
	// fmt.Println(cell)

	var monthCol int
	var monthRow int

	// get month row and col
	rows := xlsx1.GetRows(sheetName)
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

	excelFileName := "./data.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	sli := make([]string, 0)
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

					sli = append(sli, name+" :"+workInfo)

					break
				}
			}
		}
	}

	fmt.Println(sli)

	attaSli := make([]*slack.Attachment, 0)

	for _, ele := range sli {
		temp := &slack.Attachment{Text: ele}
		attaSli = append(attaSli, temp)
	}

	err2 := hook.PostMessage(&slack.WebHookPostPayload{

		Text:        "おはよう！日直バざるインフォだよ",
		Attachments: attaSli,
	})
	if err2 != nil {
		panic(err)
	}
}
