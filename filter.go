package filter

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("filter")

type Filter struct {
	Next plugin.Handler
}

func (h Filter) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debug("Testing filter")
	return h.Next.ServeDNS(ctx, w, r)
}

func (h Filter) Name() string { return "filter" }
