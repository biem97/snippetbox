package main

import (
	"testing"
	"time"

	"github.com/biem97/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() functin (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 1, 29, 9, 30, 0, 15, time.UTC),
			want: "29 Jan 2023 at 09:30",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 1, 29, 9, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "29 Jan 2023 at 08:15",
		},
	}

	// Loop over the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ARRANGE AND ACT
			got := humanDate(tt.tm)

			// ASSERTION
			assert.Equal(t, got, tt.want)
		})
	}
}
