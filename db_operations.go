package gedis

import (
	"encoding/binary"
	"time"

	"github.com/anteiro255/gedis/pkg/protocol"
	"github.com/anteiro255/gedis/pkg/protocol/status"
)

func (c *Client) Set(key [protocol.RequestKeySize]byte, value []byte) error {
	// Sending the request to the server
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.Set,
			key,
			value,
		).ToBytes(),
	)
	if err != nil {
		return err
	}

	// Reading the response from the server
	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return err
	}
	if response.Header.Status != status.OK {
		return response.Header.Status
	}
	return nil
}

func (c *Client) Get(key [protocol.RequestKeySize]byte) ([]byte, error) {
	// Sending the request to the server
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.Get,
			key,
			nil,
		).ToBytes(),
	)
	if err != nil {
		return nil, err
	}

	// Reading the response from the server
	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return nil, err
	}
	if response.Header.Status != status.OK {
		return nil, response.Header.Status
	}
	return response.Body, nil
}

func (c *Client) Del(key [protocol.RequestKeySize]byte) error {
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.Del,
			key,
			nil,
		).ToBytes(),
	)
	if err != nil {
		return err
	}

	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return err
	}
	if response.Header.Status != status.OK {
		return response.Header.Status
	}
	return nil
}

func (c *Client) Exist(key [protocol.RequestKeySize]byte) (bool, error) {
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.Exist,
			key,
			nil,
		).ToBytes(),
	)
	if err != nil {
		return false, err
	}

	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return false, err
	}
	if response.Header.Status == status.NoSuchKey {
		return false, nil
	}
	if response.Header.Status != status.OK {
		return false, response.Header.Status
	}
	return true, nil
}

func (c *Client) TTLSet(key [protocol.RequestKeySize]byte, seconds uint32) error {
	body := make([]byte, 4)
	binary.BigEndian.PutUint32(body, seconds)

	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.TTL_Set,
			key,
			body,
		).ToBytes(),
	)
	if err != nil {
		return err
	}

	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return err
	}
	if response.Header.Status != status.OK {
		return response.Header.Status
	}
	return nil
}

func (c *Client) TTLGet(key [protocol.RequestKeySize]byte) (uint32, error) {
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.TTL_Get,
			key,
			nil,
		).ToBytes(),
	)
	if err != nil {
		return 0, err
	}

	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return 0, err
	}
	if response.Header.Status != status.OK {
		return 0, response.Header.Status
	}
	if len(response.Body) < 4 {
		return 0, nil
	}
	return binary.BigEndian.Uint32(response.Body), nil
}

func (c *Client) TTLDel(key [protocol.RequestKeySize]byte) error {
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.TTL_Del,
			key,
			nil,
		).ToBytes(),
	)
	if err != nil {
		return err
	}

	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return err
	}
	if response.Header.Status != status.OK {
		return response.Header.Status
	}
	return nil
}

func (c *Client) TTLExist(key [protocol.RequestKeySize]byte) (bool, error) {
	_, err := c.conn.Write(
		protocol.NewRequest(
			protocol.TTL_Exist,
			key,
			nil,
		).ToBytes(),
	)
	if err != nil {
		return false, err
	}

	response, err := getResponse(c.conn, 3*time.Second)
	if err != nil {
		return false, err
	}
	if response.Header.Status == status.NoSuchKey {
		return false, nil
	}
	if response.Header.Status != status.OK {
		return false, response.Header.Status
	}
	return true, nil
}
