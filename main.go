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

	t := time.Now()
	_, month, day := t.Date()
	fmt.Printf("Current time: %v", t)

	monthInt := int(month)

	sheetIndex := 0
	if monthInt < 7 {
		sheetIndex = 1
	} else {
		sheetIndex = 2
	}

	sheetName := xlsx1.GetSheetName(sheetIndex)

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

	var res string
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

					if workInfo == "D" && workInfoCell.GetStyle().Fill.FgColor == "" {
						workInfo = "D1"
					}

					workInfoStr := ""
					switch workInfo {
					case "D1":
						workInfoStr = "当番"
					case "D2":
						workInfoStr = "当番(副)"
					case "D":
						workInfoStr = "通常勤務"
					case "R":
						workInfoStr = "準備期間"
					case "I":
						workInfoStr = "出張移動"
					case "T":
						workInfoStr = "出張"
					}

					if workInfoStr == "" {
						break
					}
					sli = append(sli, name+" :"+workInfo)
					res = res + name + ": \t" + workInfoStr + "\n"
					break
				}
			}
		}
	}

	err2 := hook.PostMessage(&slack.WebHookPostPayload{
		Text: "おはよう！日直バざるだよ\n\n" + res,
	})
	if err2 != nil {
		panic(err)
	}
}
