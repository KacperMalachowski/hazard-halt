package main

import "testing"

func Test_containsMaliciousPage(t *testing.T) {
	page := MaliciousPage{
		RegisterPositionId: 1,
		DomainAdress:       "test.com",
		InsertDate:         "2021-01-01",
		DeleteDate:         "",
	}

	message := "Free Discord Nitro at test.com"

	if !containsMaliciousPage(message, []MaliciousPage{page}) {
		t.Error("Expected true, got false")
	}
}
