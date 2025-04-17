package dns

import (
	"errors"
)

type RecordNS struct {
	Ttl  uint32 `json:"ttl,string"`
	Host string `json:"host"`
}

func (r *RecordNS) Valid() error {
	if !RegexHost.MatchString(r.Host) {
		return errors.New("host:invalid host")
	}
	return nil
}
