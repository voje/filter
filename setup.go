package filter

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("foo", setup) }

func setup(c *caddy.Controller) error {
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Filter{Next: next}
	})
	return nil
}
