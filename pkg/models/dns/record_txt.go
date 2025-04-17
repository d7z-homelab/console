package dns

import (
	"errors"
)

type RecordTXT struct {
	Ttl  uint32 `json:"ttl"`
	Text string `json:"text"`
}

func (r *RecordTXT) Valid() error {
	if r.Text == "" {
		return errors.New("text:empty text")
	}
	return nil
}
