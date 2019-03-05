package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var binaryName = "git-bug"

func fixturePath(t *testing.T, fixture string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("problems recovering caller information")
	}

	return filepath.Join(filepath.Dir(filename), fixture)
}

func loadFixture(t *testing.T, fixture string) string {
	content, err := ioutil.ReadFile(fixturePath(t, fixture))
	if err != nil {
		t.Fatal(err)
	}

	return string(content)
}

func TestCliArgs(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"version-argument", []string{"version"}, "version-argument.golden"},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			dir, err := os.Getwd()

			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.args...)

			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Fatal(err)
			}

			actual := string(output)
			expected := loadFixture(t, tt.fixture)

			actual = strings.TrimRight(actual, "\n")

			if actual != expected {
				t.Fatalf("\n%s\n%s\n%d\n%d", actual, expected, len(actual), len(expected))
			}

		})

	}

}

func TestMain(m *testing.M) {
	err := os.Chdir("..")

	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	make := exec.Command("make")
	err = make.Run()
	if err != nil {
		fmt.Printf("could not make binary file for git-bug: %v", err)
		os.Exit(1)
	}

	os.Exit(m.Run())

}
