package gop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	tests := map[string][]string{
		"github.com/a/b/c.d@123":    {"github.com/a/b", "c.d", "123"},
		"github.com/a/b/c/d.e@123":  {"github.com/a/b", "c/d.e", "123"},
		"github.com/a/b/c/d.e@1234": {"github.com/a/b", "c/d.e", "1234"},
		"github.com/a/b/c/d.e":      {"github.com/a/b", "c/d.e", ""},
		"github.com/a/b@d":          {"github.com/a/b", "", "d"},
		"github.com/a/b":            {"github.com/a/b", "", ""},
	}
	for in, out := range tests {
		t.Run(in, func(t *testing.T) {
			a, b, c, err := ProcessRequest(in)
			require.NoError(t, err)
			require.Equal(t, out, []string{a, b, c})
		})
	}
}
