package dns

import (
	"errors"
)

type RecordCAA struct {
	Flag  uint8  `json:"flag,string"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func (r *RecordCAA) Valid() error {
	if r.Flag != 0 && r.Flag != 128 {
		return errors.New("flag:invalid flag")
	}
	switch r.Tag {
	case "issue", "issuewild", "iodef", "contactphone":
	case "contactemail":
		if !RegexMail.MatchString(r.Value) {
			return errors.New("value:invalid email, tag == contactemail")
		}
	default:
		return errors.New("tag:invalid tag")
	}
	return nil
}
