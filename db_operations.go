package gedis

import (
	"time"

	"github.com/anteiro255/gedis/pkg/protocol"
)

func (c *Client) Set(key [protocol.RequestKeySize]byte, value []byte) error {
	// Sending the request to the server
	c.conn.Write(
		protocol.NewRequest(
			protocol.Set,
			key,
			value,
		).ToBytes(),
	)

	// Reading the response from the server
	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return err
	}
	return response.Header.Status
}

func (c *Client) Get(key [protocol.RequestKeySize]byte) ([]byte, error) {
	// Sending the request to the server
	c.conn.Write(
		protocol.NewRequest(
			protocol.Get,
			key,
			nil,
		).ToBytes(),
	)

	// Reading the response from the server
	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return nil, err
	}
	return response.Body, response.Header.Status
}
