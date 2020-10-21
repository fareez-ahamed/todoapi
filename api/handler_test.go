package api

import "testing"

func TestSlashStripEnd(t *testing.T) {
	type record struct {
		input  string
		output string
	}
	data := []record{
		{"/api/todo", "/api/todo"},
		{"/api/todo/", "/api/todo"},
	}
	for _, rec := range data {
		if res := stripEndSlash(rec.input); res != rec.output {
			t.Errorf("Expected %v\nGot %v", rec.output, res)
		}
	}
}
