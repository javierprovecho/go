package test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Marshaler struct {
	X string
}

func encode(str string) string {
	buf := bytes.Buffer{}
	b64 := base64.NewEncoder(base64.StdEncoding, &buf)
	if _, err := b64.Write([]byte(str)); err != nil {
		panic(err)
	}
	if err := b64.Close(); err != nil {
		panic(err)
	}
	return buf.String()
}

func decode(str string) string {
	if len(str) == 0 {
		return ""
	}
	b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	bs := make([]byte, len(str))
	if n, err := b64.Read(bs); err != nil {
		panic(err)
	} else {
		bs = bs[:n]
	}
	return string(bs)
}

func (m Marshaler) MarshalJSON() ([]byte, error) {
	return []byte(`"MANUAL__` + encode(m.X) + `"`), nil
}

func (m *Marshaler) UnmarshalJSON(text []byte) error {
	m.X = decode(strings.TrimPrefix(strings.Trim(string(text), `"`), "MANUAL__"))
	return nil
}

var _ json.Marshaler = Marshaler{}
var _ json.Unmarshaler = &Marshaler{}

type A Marshaler

type T A
