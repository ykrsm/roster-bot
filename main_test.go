package main

import (
	"fmt"
	"testing"
	"time"
)

/*
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
*/

func TestMain(t *testing.T) {
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
