package filter

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

type FilterHandler struct {
	Next plugin.Handler
}

func (h FilterHandler) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return h.Next.ServeDNS(ctx, w, r)
}

func (h FilterHandler) Name() string { return "filter" }
