package filter

import (
	"bufio"
	"os"
	"strings"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("filter", setup) }

func setup(c *caddy.Controller) error {
	var filter Filter
	filter.Blacklist = make(map[string]bool)

	// Read blacklist
	c.Next() // Skip 'filter'
	hasNextToken := c.NextArg()
	if !hasNextToken {
		return plugin.Error("filter plugin's argument is the path to a blacklist file", c.ArgErr())
	}
	fileName := c.Val()

	file, err := os.Open(fileName)
	if err != nil {
		return plugin.Error("Failed reading file: "+fileName, err)
	}
	defer file.Close()

	// Save blacklist entries into 'filter'
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		line = strings.Trim(line, " \n\t\r")
		if line == "" {
			continue
		}
		filter.Blacklist[line] = true
	}

	// Check for read errors
	err = s.Err()
	if err != nil {
		return plugin.Error("Error while reading file", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		filter.Next = next
		return filter
	})
	return nil
}
