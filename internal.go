package gedis

import (
	"io"
	"net"
	"time"

	"github.com/anteiro255/gedis/pkg/protocol"
)

func getResponse(conn net.Conn, timeout time.Duration) (*protocol.Response, error) {
	err := conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, err
	}
	defer conn.SetReadDeadline(time.Time{})

	var headerBytes [protocol.ResponseHeaderSize]byte
	_, err = io.ReadFull(conn, headerBytes[:])
	if err != nil {
		return nil, err
	}

	var r protocol.Response
	r.Header = protocol.NewResponseHeaderFromBytes(headerBytes)

	body := make([]byte, r.Header.BodySize)
	if r.Header.BodySize > 0 {
		_, err = io.ReadFull(conn, body)
		if err != nil {
			return nil, err
		}
	}

	r.Body = body

	return &r, nil
}
