package result

import (
	"maps"
	"net"
	"slices"
)

type Result struct {
	// matched servers per new subnet
	matched map[string][]string
	// unmatched servers
	unmatched []string
}

func New(subnets []*net.IPNet) *Result {
	out := new(Result)

	out.matched = make(map[string][]string)

	for _, el := range subnets {
		out.matched[el.String()] = make([]string, 0)
	}

	out.unmatched = make([]string, 0)

	return out
}

func (r *Result) AddMatched(subnet *net.IPNet, srvr string) {
	val := r.matched[subnet.String()]
	if val == nil {
		val = make([]string, 0)
	}

	if !slices.Contains(val, srvr) {
		val = append(val, srvr)
	}

	r.matched[subnet.String()] = val
}

func (r *Result) AddUnmatched(srvr string) {
	if slices.Contains(r.unmatched, srvr) {
		return
	}

	r.unmatched = append(r.unmatched, srvr)
}

func (r *Result) GetMatched() map[string][]string {
	out := maps.Clone(r.matched)

	return out
}

func (r *Result) GetUnmatched() []string {
	out := slices.Clone(r.unmatched)

	return out
}
