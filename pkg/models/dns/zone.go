package dns

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type Zone struct {
	Records map[string]*Records `json:"records"`
	SOA     *RecordSOA          `json:"soa"`
}

func NewZone(base string, dns net.IP) (*Zone, error) {
	base = strings.TrimRight(strings.ToLower(base), ".")
	if !RegexHost.MatchString(base) {
		return nil, errors.New("invalid zone name: " + base)
	}
	if dns == nil {
		return nil, errors.New("invalid zone dns")
	}
	records := make(map[string]*Records)
	record := NewRecord()
	record.NS = append(record.NS, &RecordNS{
		Ttl:  3600,
		Host: fmt.Sprintf("dns.%s", base),
	})
	if dns.To4() != nil {
		record.A = append(record.A, &RecordA{
			Ttl: 3600,
			Ip:  dns,
		})
	} else {
		record.AAAA = append(record.AAAA, &RecordAAAA{
			Ttl: 3600,
			Ip:  dns,
		})
	}

	records["@"] = record
	return &Zone{
		Records: records,
		SOA: &RecordSOA{
			Ttl:     3600,
			MName:   fmt.Sprintf("dns.%s", base),
			RName:   fmt.Sprintf("sa.%s", base),
			Serial:  1,
			Refresh: 3600,
			Retry:   1800,
			Expire:  7200,
			Minimum: 10,
		},
	}, nil
}

func (z *Zone) SetSOA(soa RecordSOA) error {
	if err := soa.Valid(); err != nil {
		return err
	}
	z.SOA = &soa
	return nil
}
