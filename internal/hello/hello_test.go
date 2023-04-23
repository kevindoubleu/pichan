package hello_test

import (
	"testing"

	"github.com/kevindoubleu/pichan/internal/hello"
)

func TestBuildGreeting(t *testing.T) {
	testCases := []struct {
		desc     string
		name     string
		expected string
	}{
		{
			desc:     "returns expected string",
			name:     "there",
			expected: hello.Greeting + " there!",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := hello.BuildGreeting(tC.name)
			if tC.expected != actual {
				t.Errorf("%s != %s", tC.expected, actual)
			}
		})
	}
}
