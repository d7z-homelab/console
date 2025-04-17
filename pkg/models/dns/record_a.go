package dns

import (
	"errors"
	"net"
)

type RecordA struct {
	Ttl uint32 `json:"ttl,string"`
	Ip  net.IP `json:"ip,string"`
}

func (r *RecordA) Valid() error {
	if r.Ip == nil || !RegexIPv4.MatchString(r.Ip.String()) {
		return errors.New("ip:invalid ip address")
	}
	return nil
}
