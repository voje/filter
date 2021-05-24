package filter

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("filter")

type FilterHandler struct {
	Next plugin.Handler
}

func (h FilterHandler) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debug("Testing filter")
	return h.Next.ServeDNS(ctx, w, r)
}

func (h FilterHandler) Name() string { return "filter" }
