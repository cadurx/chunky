package chunky

import (
	"bytes"
	"encoding/binary"
	"io"
)

func WriteHeader(c io.Writer, count int64) error {
	var err error
	buf := &bytes.Buffer{}

	err = binary.Write(buf, be, magick)
	if err != nil {
		return err
	}
	err = binary.Write(buf, be, protocol_version)
	if err != nil {
		return err
	}
	err = binary.Write(buf, be, count)
	if err != nil {
		return err
	}
	_, err = c.Write(buf.Bytes())
	return err
}

func WriteFileHeader(c io.Writer, name string, size int64) error {
	var err error
	buf := &bytes.Buffer{}

	err = binary.Write(buf, be, int64(len(name)))
	if err != nil {
		return err
	}
	_, err = buf.Write([]byte(name))
	if err != nil {
		return err
	}
	err = binary.Write(buf, be, size)
	if err != nil {
		return err
	}
	_, err = c.Write(buf.Bytes())
	return err
}
