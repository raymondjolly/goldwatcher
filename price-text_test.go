package main

import (
	"testing"
)

func TestApp_getPriceText(t *testing.T) {
	open, _, _ := testApp.getPriceText()
	if open.Text != "Open: $1815.1350 USD" {
		t.Error("Open price does not match: ", open.Text)
	}
}
