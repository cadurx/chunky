package chunky

import (
	"encoding/binary"
	"encoding/json"
)

const magick int64 = 5460303
const protocol_version int64 = 0

var be = binary.BigEndian

type Event struct {
	Severity string `json:"severity"`
	Msg      string `json:"msg"`

	Row int `json:"row"`
	Col int `json:"col"`
}

type Log struct {
	Exception string `json:"exception"`
	Row       int    `json:"row"`
	Col       int    `json:"col"`

	Events []Event `json:"events"`
}

func ParseLog(json_bytes []byte) (*Log, error) {
	l := Log{}

	err := json.Unmarshal(json_bytes, &l)

	if err != nil {
		return nil, err
	}

	return &l, nil
}
