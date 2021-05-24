package filter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const blacklist = `
9gag.com
www.youtube.com
www.google.com.
reddit.com.
one.two.net.

simple-name
`

func TestAbs(t *testing.T) {
	f := NewFilter()
	inputReader := strings.NewReader(blacklist)
	f.ParseBlacklist(inputReader)
	t.Logf("%+v", f.Blacklist)

	assert.True(t, f.Blocks("9gag.com"))
	assert.True(t, f.Blocks("www.9gag.com"))
	assert.False(t, f.Blocks("www.9gag"))
	assert.False(t, f.Blocks("9gag"))
	assert.False(t, f.Blocks("www.wikipedia.com"))

	// Filter can't parse 'simple-name' so don't block it (uncaught entry format, permissive)
	assert.False(t, f.Blocks("simple-name"))
}
