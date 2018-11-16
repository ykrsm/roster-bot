package main

type WorkInfo int

const (
	On WorkInfo = iota // Working day
	Duty
	SubDuty
	Prepare
	Trip
	Moving
	Education
	Off
)

func (workInfo WorkInfo) String() string {
	workInfoStrs := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}
	return workInfoStrs[workInfo]
}
