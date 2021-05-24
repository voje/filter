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
	fqdn := r.Question[0].Name

	redirectFqdn := "localhost"

	if _, ok := h.Blacklist[fqdn]; ok {
		log.Infof("Redirecting: %s => %s", redirectFqdn)
		r.Question[0].Name = redirectFqdn // redirect to localhost
	}

	return h.Next.ServeDNS(ctx, w, r)
}

func (h Filter) Name() string { return "filter" }
