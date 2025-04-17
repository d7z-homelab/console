package dns

import (
	"errors"
)

type RecordSOA struct {
	Ttl     uint32 `json:"ttl,string"`
	MName   string `json:"mname"`
	RName   string `json:"rname"`
	Serial  uint16 `json:"serial,string"`
	Refresh uint32 `json:"refresh,string"`
	Retry   uint32 `json:"retry,string"`
	Expire  uint32 `json:"expire,string"`
	Minimum uint32 `json:"minimum,string"`
}

func (r *RecordSOA) Valid() error {
	if !RegexHost.MatchString(r.MName) {
		return errors.New("mnane:invalid host")
	}
	if !RegexHost.MatchString(r.RName) {
		return errors.New("rname:invalid host")
	}
	if r.Retry >= r.Refresh {
		return errors.New("retry >= refresh")
	}
	if r.Expire <= r.Refresh+r.Retry {
		return errors.New("expire < retry + refresh")
	}
	return nil
}
