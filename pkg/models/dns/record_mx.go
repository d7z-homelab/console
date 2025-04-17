package dns

import (
	"errors"
)

type RecordMX struct {
	Ttl        uint32 `json:"ttl,string"`
	Host       string `json:"host"`
	Preference uint16 `json:"preference,string"`
}

func (r *RecordMX) Valid() error {
	if !RegexHost.MatchString(r.Host) {
		return errors.New("host:invalid host")
	}
	return nil
}
