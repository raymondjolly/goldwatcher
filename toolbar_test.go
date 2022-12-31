package main

import "testing"

func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolBar()
	//toolbar includes the space bar so account for it in the length
	if len(tb.Items) != 4 {
		t.Error("Incorrect number of items in toolbar")
	}
}
