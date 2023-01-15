package hello_test

import (
	"testing"

	"github.com/kevindoubleu/pichan/pkg/hello"
)

func TestBuildGreeting(t *testing.T) {
	testCases := []struct {
		desc     string
		expected string
	}{
		{
			desc:     "returns expected string",
			expected: hello.Greeting,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := hello.BuildGreeting()
			if tC.expected != actual {
				t.Errorf("%s != %s", tC.expected, actual)
			}
		})
	}
}
