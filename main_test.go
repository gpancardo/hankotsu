package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractColumn(t *testing.T) {
	cases := []struct {
		line string
		idx  int
		want string
	}{
		{"1,Mexico City,alpha\n", 1, "Mexico City"},
		{"1,Mexico City,alpha", 2, "alpha"},
		{"1,,alpha\r\n", 1, ""},
		{"a,b\n", 5, ""},
	}
	for _, c := range cases {
		if got := extractColumn(c.line, c.idx); got != c.want {
			t.Errorf("extractColumn(%q, %d) = %q, want %q", c.line, c.idx, got, c.want)
		}
	}
}

// The last row has no trailing newline and the input lives in a subfolder:
// both used to break (row dropped, output path wrong).
func TestFilterKeepsLastRowAndWritesNextToInput(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "data")
	os.Mkdir(dir, 0o755)
	in := filepath.Join(dir, "a.csv")
	os.WriteFile(in, []byte("id,city\n1,Mexico City\n2,Monterrey\n3,Mexico City"), 0o644)

	if err := filterCSVStream(Compass{Label: "city", Words: []string{"Mexico"}}, 1, in); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(dir, "READY_a.csv"))
	if err != nil {
		t.Fatal(err)
	}
	want := "id,city\n1,Mexico City\n3,Mexico City"
	if string(got) != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}
