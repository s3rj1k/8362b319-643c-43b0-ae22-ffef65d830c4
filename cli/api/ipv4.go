package api

import (
	"net"
)

func (r *Response) GetIPv4List() []net.IP {
	out := make([]net.IP, 0)

	if r.Station.To4() != nil {
		out = append(out, r.Station)
	}

	for _, el := range r.IPs {
		if el.IP.IP.To4() == nil {
			continue
		}

		if el.IP.IP.Equal(r.Station) {
			continue
		}

		out = append(out, r.Station)
	}

	return out
}
