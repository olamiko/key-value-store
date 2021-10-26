package utils_test

import (
	"github.com/olamiko/key-value-store/utils/Store"
	"os"
	"testing"
)

func setup() {
	testStore := Store{filename: "test-storage.kv"}
}

func shutdown() {
	//delete storage file
}

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
			testStore.set(tc.key, tc.value)
			result := testStore.get(tc.key)
			if result != tc.value {
				t.Fatalf("expected value %v, but got %v", tc.value, result)
			}
		})
	}

}

func TestSaveAndLoad(t *testing.T) {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	defer os.Exit(code)

}
