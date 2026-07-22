package gedis

import (
	"net"
	"strconv"
)

type Client struct {
	Address string
	Port    int
	conn    net.Conn
}

func NewClient(address string, port int) *Client {
	return &Client{
		Address: address,
		Port:    port,
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.Address+":"+strconv.Itoa(c.Port))
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
