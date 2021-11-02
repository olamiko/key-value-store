package loggin_test

import (
	"github.com/olamiko/key-value-store/loggin"
	"testing"
)

var commitSlice []string

func TestWriteAndReadCommitLog(t *testing.T) {

	loggin.SetCommitLog("testLog.lg")

	cases := []struct {
		name   string
		key    string
		value  string
		result string
	}{
		{"Commit case 1", "food", "apple", "1 SET food apple"},
		{"Commit case 2", "nation", "austria", "2 SET nation austria"},
		{"Commit case 3", "name", "johnny", "3 SET name johnny"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			loggin.WriteCommitLog(tc.key, tc.value)

		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			commitSlice = loggin.ReadCommitLog()
		})
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			commitLine := commitSlice[i]
			if tc.result != commitLine {
				t.Fatalf("expected value %v, but got %v", tc.result, commitLine)
			}

		})
	}

	loggin.RotateLog()

}
