package main

import "testing"

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Fail()
	}
}
