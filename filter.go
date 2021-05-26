package filter

import (
	"bufio"
	"context"
	"io"
	"net"
	"regexp"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

var fqdnRegex = regexp.MustCompile(`(\w*\.\w*).?$`)

type Filter struct {
	Next      plugin.Handler
	Blacklist map[string]bool
}

func NewFilter() Filter {
	f := Filter{}
	f.Blacklist = make(map[string]bool)
	return f
}

func ParseFQDN(s string) (string, bool) {
	fqdnGroups := fqdnRegex.FindSubmatch([]byte(s))
	if len(fqdnGroups) < 2 {
		return "", false
	}
	return string(fqdnGroups[1]), true
}

// Check whether allows certain fqdn
func (f Filter) Blocks(fqdn string) bool {
	fqdn, formatOk := ParseFQDN(fqdn)
	if !formatOk {
		return false
	}

	_, blacklisted := f.Blacklist[fqdn]
	return blacklisted
}

func (f Filter) ParseBlacklist(r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		line = strings.Trim(line, " \n\t\r")
		if line == "" {
			continue
		}

		if fqdn, ok := ParseFQDN(line); ok {
			f.Blacklist[fqdn] = true
		}
	}

	return s.Err()
}

func (f Filter) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	fqdn := r.Question[0].Name

	if f.Blocks(fqdn) {
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

	return f.Next.ServeDNS(ctx, w, r)
}

func (f Filter) Name() string { return "filter" }
