package dns

import (
	"errors"
	"net"
)

type RecordAAAA struct {
	Ttl uint32 `json:"ttl,string"`
	Ip  net.IP `json:"ip"`
}

func (r *RecordAAAA) Valid() error {
	if r.Ip == nil || !RegexIPv6.MatchString(r.Ip.String()) {
		return errors.New("ip:invalid ip address")
	}
	return nil
}
