package main

import "testing"

func TestParser(t *testing.T) {
	config := Parser()
	if config.Account.Username != "17681888658" {
		t.Errorf("Parser is not work. %+v\n", config)
	}
}
