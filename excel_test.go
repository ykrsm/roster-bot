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

func TestExcel3_31(t *testing.T) {
	then := time.Date(
		2018, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	actual := fillRoster(2018, 3, 31, "./data.xlsx", r)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎２ (US)", Off},
			Employee{"田中 太郎３ (US)", Off},
			Employee{"田中 太郎４ (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Duty},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestExcel6_1(t *testing.T) {
	then := time.Date(
		2018, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	actual := fillRoster(2018, 6, 1, "./data.xlsx", r)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"田中 太郎 (US)", Moving},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎２ (US)", On},
			Employee{"田中 太郎３ (US)", Off},
			Employee{"田中 太郎４ (US)", On},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Duty},
			Employee{"田中 太郎 (US)", Off},
			Employee{"田中 太郎 (US)", Off},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestExcel7_15(t *testing.T) {
	then := time.Date(
		2018, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	actual := fillRoster(2018, 7, 15, "./data.xlsx", r)
	expected := Roster{
		Date: then,
		Employees: []Employee{

			Employee{"田中 太郎子 (SC)", Off},
			Employee{"田中 太郎子 (SC)", Duty},
			Employee{"田中 太郎 (SC)", Off},
			Employee{"田中 太郎子 (SC)", Off},
			Employee{"田中 太郎 (Irving)", Off},
			Employee{"田中 太郎子 (Irving)", Off},
			Employee{"田中 太郎子 (SC)", Off},
			Employee{"田中 太郎 (Irving)", Off},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Test excel for end and beg of year
// contains 2018 2nd half and 2019 1st half

func TestExcel_2_7_15(t *testing.T) {
	then := time.Date(
		2018, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	actual := fillRoster(2018, 7, 15, "./data2.xlsx", r)
	expected := Roster{
		Date: then,
		Employees: []Employee{

			Employee{"田中 太郎子 (SC)", Off},
			Employee{"田中 太郎子 (SC)", Duty},
			Employee{"田中 太郎 (SC)", Off},
			Employee{"田中 太郎子 (SC)", Off},
			Employee{"田中 太郎 (Irving)", Off},
			Employee{"田中 太郎子 (Irving)", Off},
			Employee{"田中 太郎子 (SC)", Off},
			Employee{"田中 太郎 (Irving)", Off},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
