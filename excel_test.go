package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestExcel7_15(t *testing.T) {
	then := time.Date(
		2018, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	names := []string{"一田中", "二田中", "三田中", "四田中", "五田中", "六田中", "七田中", "八田中", "九田中", "十田中"}
	actual := fillRoster(2018, 7, 15, "./data.xlsx", r, names)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"一田中 太郎 (US)", Off},
			Employee{"二田中 太郎２ (US)", Duty},
			Employee{"三田中 太郎３ (US)", Off},
			Employee{"四田中 太郎４ (US)", Off},
			Employee{"五田中 太郎 (US)", Off},
			Employee{"六田中 太郎 (US)", Off},
			Employee{"七田中 太郎 (US)", Off},
			Employee{"八田中 太郎 (US)", Off},
			Employee{"九田中 太郎 (US)", Off},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

func TestExcel1_1(t *testing.T) {
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := Roster{Date: then}
	names := []string{"一田中", "二田中", "三田中", "四田中", "五田中", "六田中", "七田中", "八田中", "九田中", "十田中"}
	actual := fillRoster(2018, 1, 1, "./data.xlsx", r, names)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"一田中 太郎 (US)", Off},
			Employee{"二田中 太郎２ (US)", SubDuty},
			Employee{"三田中 太郎３ (US)", Off},
			Employee{"四田中 太郎４ (US)", Off},
			Employee{"五田中 太郎 (US)", Off},
			Employee{"六田中 太郎 (US)", Off},
			Employee{"七田中 太郎 (US)", Off},
			Employee{"八田中 太郎 (US)", Off},
			Employee{"九田中 太郎 (US)", Off},
			Employee{"十田中 太郎 (US)", Duty},
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
	names := []string{"一田中", "二田中", "三田中", "四田中", "五田中", "六田中", "七田中", "八田中", "九田中", "十田中"}
	actual := fillRoster(2018, 3, 31, "./data.xlsx", r, names)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"一田中 太郎 (US)", Off},
			Employee{"二田中 太郎２ (US)", Off},
			Employee{"三田中 太郎３ (US)", Off},
			Employee{"四田中 太郎４ (US)", Off},
			Employee{"五田中 太郎 (US)", Off},
			Employee{"六田中 太郎 (US)", Duty},
			Employee{"七田中 太郎 (US)", Off},
			Employee{"八田中 太郎 (US)", Off},
			Employee{"九田中 太郎 (US)", Off},
			Employee{"十田中 太郎 (US)", Off},
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
	names := []string{"一田中", "二田中", "三田中", "四田中", "五田中", "六田中", "七田中", "八田中", "九田中", "十田中"}
	actual := fillRoster(2018, 7, 15, "./data2.xlsx", r, names)
	expected := Roster{
		Date: then,
		Employees: []Employee{
			Employee{"一田中 太郎 (US)", Off},
			Employee{"二田中 太郎２ (US)", Duty},
			Employee{"三田中 太郎３ (US)", Off},
			Employee{"四田中 太郎４ (US)", Off},
			Employee{"五田中 太郎 (US)", Off},
			Employee{"六田中 太郎 (US)", Off},
			Employee{"七田中 太郎 (US)", Off},
			Employee{"八田中 太郎 (US)", Off},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
