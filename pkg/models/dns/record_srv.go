package dns

import (
	"errors"
)

type RecordSRV struct {
	Ttl      uint32 `json:"ttl,string"`
	Priority uint16 `json:"priority,string"`
	Weight   uint16 `json:"weight,string"`
	Port     uint16 `json:"port,string"`
	Target   string `json:"target"`
}

func (r *RecordSRV) Valid() error {
	if !RegexHost.MatchString(r.Target) {
		return errors.New("target: invalid host")
	}
	return nil
}
