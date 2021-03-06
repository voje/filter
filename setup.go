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
		return plugin.Error("filter plugin's first argument is the path to a blacklist file", c.ArgErr())
	}
	fileName := c.Val()

	file, err := os.Open(fileName)
	err = filter.ParseBlacklist(file)

	if err != nil {
		return plugin.Error("Failed reading file: "+fileName, err)
	}
	defer file.Close()

	// Set up logging
	os.Mkdir("/var/log/coredns/", 0755)

	f, err := os.OpenFile("/var/log/coredns/filter.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return plugin.Error("Error opening file: ", err)
	}
	// Don't close log file

	filter.log.SetOutput(f)

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		filter.Next = next
		return filter
	})
	return nil
}
