package filter

import (
	"context"
	"net"

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

	fqdn := r.Question[0].Name
	log.Info("fqdn: " + fqdn)

	if _, ok := h.Blacklist[fqdn]; ok {
		reply := "127.0.0.1"
		log.Infof("Redirecting: %s => %s", fqdn, reply)

		answers := []dns.RR{}

		rr := new(dns.A)
		rr.Hdr = dns.RR_Header{Name: fqdn, Rrtype: dns.TypeA, Class: dns.ClassINET}
		rr.A = net.ParseIP(reply).To4()

		answers = append(answers, rr)

		m := new(dns.Msg)
		m.SetReply(r)
		m.Authoritative = true
		m.Answer = answers

		w.WriteMsg(m)
		return dns.RcodeSuccess, nil
	}

	return h.Next.ServeDNS(ctx, w, r)
}

func (h Filter) Name() string { return "filter" }
