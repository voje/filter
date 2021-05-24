package filter

import (
	"os"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("filter", setup) }

func setup(c *caddy.Controller) error {
	filter := NewFilter()

	// Read blacklist
	c.Next() // Skip 'filter'
	hasNextToken := c.NextArg()
	if !hasNextToken {
		return plugin.Error("filter plugin's argument is the path to a blacklist file", c.ArgErr())
	}
	fileName := c.Val()

	file, err := os.Open(fileName)
	err = filter.ParseBlacklist(file)

	if err != nil {
		return plugin.Error("Failed reading file: "+fileName, err)
	}
	defer file.Close()

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		filter.Next = next
		return filter
	})
	return nil
}
