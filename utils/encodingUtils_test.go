package utils_test

import (
	"github.com/olamiko/key-value-store/utils"
	"os"
	"testing"
)

var testStorage string = "test-storage.kv"

func TestSetAndGet(t *testing.T) {

	cases := []struct {
		name  string
		key   string
		value string
	}{
		{"Simple case 1", "food", "apple"},
		{"Simple case 2", "nation", "austria"},
		{"Simple case 3", "name", "johnny"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			testStore := utils.NewStore(testStorage)
			testStore.Set(tc.key, tc.value)
			result := testStore.Get(tc.key)
			if result != tc.value {
				t.Fatalf("expected value %v, but got %v", tc.value, result)
			}
		})
	}

	os.Remove(testStorage)

}

func TestSaveAndLoad(t *testing.T) {

	cases := []struct {
		name  string
		key   string
		value string
	}{
		{"Simple case 1", "food", "apple"},
		{"Simple case 2", "nation", "austria"},
		{"Simple case 3", "name", "johnny"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			testStore := utils.NewStore(testStorage)
			testStore.Set(tc.key, tc.value)
		})
		t.Run(tc.name, func(t *testing.T) {

			testStore := utils.NewStore(testStorage)
			result := testStore.Get(tc.key)
			if result != tc.value {
				t.Fatalf("expected value %v, but got %v", tc.value, result)
			}
		})

	}
	os.Remove(testStorage)
}
