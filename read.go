package chunky

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ReadHeader(c io.Reader) (int64, error) {
	var err error
	var m int64
	var v int64
	var n int64

	err = binary.Read(c, be, &m)
	if err != nil {
		return 0, err
	}
	if m != magick {
		return 0, fmt.Errorf("bad magick number passed: %d, expected ", m, magick)
	}
	err = binary.Read(c, be, &v)
	if v != protocol_version {
		return 0, fmt.Errorf("bad protocol version: %d, expected %d", v, protocol_version)
	}
	err = binary.Read(c, be, &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func ReadFileHeader(c io.Reader) (string, int64, error) {
	var err error
	var l int64
	var n string
	var s int64

	err = binary.Read(c, be, &l)
	if err != nil {
		return "", 0, err
	}
	nb := make([]byte, l)
	_, err = c.Read(nb)
	if err != nil {
		return "", 0, err
	}
	n = string(nb)

	err = binary.Read(c, be, &s)
	if err != nil {
		return "", 0, err
	}
	return n, s, nil
}
