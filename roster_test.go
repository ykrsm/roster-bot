package main

import (
	"fmt"
	"testing"
	"time"
)

var (
	emp1    = Employee{"My name1", On}
	emp2    = Employee{"My name2", On}
	emp3    = Employee{"田中 太郎", Off}
	emp4    = Employee{"田中 花子", Duty}
	roster1 = Roster{
		Date: time.Now(),
		Employees: []Employee{
			emp1,
			emp2,
			emp3,
			emp4,
		},
	}
)

func TestRosterString(t *testing.T) {
	actual := Roster{
		Date: time.Now(),
		Employees: []Employee{
			emp1,
			emp2,
			emp3,
			emp4,
		},
	}.String()
	expected := emp1.Emoji() + "\n" + emp2.Emoji() + "\n" + emp3.Emoji() + "\n" + emp4.Emoji() + "\n"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestRosterDateJp(t *testing.T) {
	roster := Roster{
		Date: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Employees: []Employee{
			emp1,
			emp2,
			emp3,
			emp4,
		},
	}
	actual := roster.DateJp()
	expected := "11月17日 (" + roster.WeekDayJp() + ")"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
func TestRosterWeekDayJp(t *testing.T) {
	actual := Roster{
		Date: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Employees: []Employee{
			emp1,
			emp2,
			emp3,
			emp4,
		},
	}.WeekDayJp()
	expected := "火"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
func TestEmployeeEmoji(t *testing.T) {
	actual := Employee{"My name", On}.Emoji()
	expected := "My name\t:kinmu:"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
func TestEmployeeString(t *testing.T) {
	actual := Employee{"My name", On}.String()
	expected := "My name\t勤務"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestStringOn(t *testing.T) {
	actual := On.String()
	expected := "勤務"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestStringSubDuty(t *testing.T) {
	actual := SubDuty.String()
	expected := "当番(副)"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestEmojiOn(t *testing.T) {
	actual := On.Emoji()
	expected := ":kinmu:"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestEmojiSubDuty(t *testing.T) {
	actual := SubDuty.Emoji()
	expected := ":touban:(副)"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestEmojiOff(t *testing.T) {
	actual := Off.Emoji()
	expected := ":kyuujitu:"
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
