package main

import "time"

type Roster struct {
	Date      time.Time
	Employees []Employee
}

type Employee struct {
	RawName string
	Status  WorkInfo
	/*
		FirstName     string
		LastName     string
		Office     string
		Hopefully do some Regex and separate raw string
	*/
}

type WorkInfo int

const (
	On WorkInfo = iota // Working day
	Duty
	SubDuty
	Prepare
	Trip
	Moving
	Off
)

func (roster Roster) String() (res string) {
	for _, emp := range roster.Employees {
		res = res + emp.Emoji() + "\n"
	}
	return res
}

func (roster Roster) DateJp() string {
	// Making date string in Japnaese
	return roster.Date.Format("1月2日 (" + roster.WeekDayJp() + ")")
}

func (roster Roster) WeekDayJp() string {
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	return wdays[roster.Date.Weekday()]
}

func (employee Employee) String() string {
	return employee.RawName + "\t" + employee.Status.String()
}

func (employee Employee) Emoji() string {
	return employee.RawName + "\t" + employee.Status.Emoji()
}

func (workInfo WorkInfo) String() string {
	strings := [...]string{
		"勤務",
		"当番",
		"当番(副)",
		"準備",
		"出張",
		"移動",
		"休み"}
	return strings[workInfo]
}

func (workInfo WorkInfo) Emoji() string {
	emojis := [...]string{
		":kinmu:",
		":touban:",
		":touban:(副)",
		":junbi:",
		":syuttyou:",
		":idou:",
		":kyuujitu:"}
	return emojis[workInfo]
}
