package info

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"

	clog "github.com/coredns/coredns/plugin/pkg/log"
)

var log = clog.NewWithPlugin("info")

type info struct {
	Next plugin.Handler
}

func (i *info) Name() string { return "info" }

func (i *info) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	clog.Infof(r.String())
	return plugin.NextOrFailure(i.Name(), i.Next, ctx, w, r)
}
