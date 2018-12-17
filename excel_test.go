package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestExcel1_1(t *testing.T) {
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	actual := fillRoster(2018, 1, 1, "./data.xlsx", r)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎２ (US)", SubDuty},
			Employee{"田中 太郎３ (US)", Off},
			Employee{"田中 太郎４ (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Duty},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
