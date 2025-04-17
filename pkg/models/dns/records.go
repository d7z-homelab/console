package dns

type Records struct {
	A     []*RecordA     `json:"a,omitempty"`
	AAAA  []*RecordAAAA  `json:"aaaa,omitempty"`
	TXT   []*RecordTXT   `json:"txt,omitempty"`
	CNAME []*RecordCNAME `json:"cname,omitempty"`
	NS    []*RecordNS    `json:"ns,omitempty"`
	MX    []*RecordMX    `json:"mx,omitempty"`
	SRV   []*RecordSRV   `json:"srv,omitempty"`
	CAA   []*RecordCAA   `json:"caa,omitempty"`
}

func NewRecord() *Records {
	return &Records{
		A:     make([]*RecordA, 0),
		AAAA:  make([]*RecordAAAA, 0),
		TXT:   make([]*RecordTXT, 0),
		CNAME: make([]*RecordCNAME, 0),
		NS:    make([]*RecordNS, 0),
		MX:    make([]*RecordMX, 0),
		SRV:   make([]*RecordSRV, 0),
		CAA:   make([]*RecordCAA, 0),
	}
}
