package filter

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("filter")

type Filter struct {
	Next      plugin.Handler
	Blacklist map[string]bool
}

func (h Filter) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Info("Testing filter")

	log.Infof("%+v", r)

	// TODO

	return h.Next.ServeDNS(ctx, w, r)
}

func (h Filter) Name() string { return "filter" }
