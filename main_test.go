package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMainRunRoster1_1(t *testing.T) {
	actual := runRoster(1, 1, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎 (US)	:kyuujitu:
田中 太郎２ (US)	:touban:(副)
田中 太郎３ (US)	:kyuujitu:
田中 太郎４ (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:touban:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing empty cells (= off day)
func TestMainRunRoster3_31(t *testing.T) {
	actual := runRoster(3, 31, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎 (US)	:kyuujitu:
田中 太郎２ (US)	:kyuujitu:
田中 太郎３ (US)	:kyuujitu:
田中 太郎４ (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:touban:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing PM
// TODO Need to take care PM
func TestMainRunRoster6_1(t *testing.T) {
	actual := runRoster(6, 1, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎 (US)	:idou:
田中 太郎 (US)	:kyuujitu:
田中 太郎２ (US)	:kinmu:
田中 太郎３ (US)	:kyuujitu:
田中 太郎４ (US)	:kinmu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:touban:
田中 太郎 (US)	:kyuujitu:
田中 太郎 (US)	:kyuujitu:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing 2nd sheet
func TestMainRunRoster7_15(t *testing.T) {
	actual := runRoster(7, 15, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎子 (SC)	:kyuujitu:
田中 太郎子 (SC)	:touban:
田中 太郎 (SC)	:kyuujitu:
田中 太郎子 (SC)	:kyuujitu:
田中 太郎 (Irving)	:kyuujitu:
田中 太郎子 (Irving)	:kyuujitu:
田中 太郎子 (SC)	:kyuujitu:
田中 太郎 (Irving)	:kyuujitu:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing empty name in the middle
func TestMainRunRoster10_2(t *testing.T) {
	actual := runRoster(10, 2, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎子 (SC)	:kinmu:
田中 太郎子 (SC)	:kinmu:
田中 太郎 (SC)	:kinmu:
田中太郎(Irving)	:touban:(副)
田中 太郎子 (Irving)	:kyuujitu:
田中太郎(Irving)	:touban:
田中太郎(Irving)	:junbi:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing middle month
func TestMainRunRoster11_30(t *testing.T) {
	actual := runRoster(11, 30, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎子 (SC)	:touban:
田中 太郎子 (SC)	:kyuujitu:
田中 太郎 (SC)	:kyuujitu:
田中太郎(Irving)	:kyuujitu:
田中 太郎子 (Irving)	:kyuujitu:
田中太郎(Irving)	:kinmu:
田中太郎(Irving)	:kinmu:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing business trip day
func TestMainRunRoster11_9(t *testing.T) {
	actual := runRoster(11, 9, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎子 (SC)	:kinmu:
田中 太郎子 (SC)	:touban:
田中 太郎 (SC)	:kinmu:
田中太郎(Irving)	:kyuujitu:
田中 太郎子 (Irving)	:syuttyou:
田中太郎(Irving)	:kinmu:
田中太郎(Irving)	:junbi:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}

// Testing end of year
func TestMainRunRoster12_31(t *testing.T) {
	actual := runRoster(12, 31, "./data.xlsx", Roster{Date: time.Now()})
	expected := `田中 太郎子 (SC)	:kyuujitu:
田中 太郎子 (SC)	:touban:
田中 太郎 (SC)	:kinmu:
田中太郎(Irving)	:kyuujitu:
田中 太郎子 (Irving)	:kyuujitu:
田中太郎(Irving)	:kyuujitu:
田中太郎(Irving)	:kyuujitu:
`
	if actual != expected {
		fmt.Printf("got: %v, but want: %v", actual, expected)
		t.Fail()
	}
}
