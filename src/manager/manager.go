package manager

import (
	"io"
	"net"
)

type Client struct {
	Conn *net.TCPConn
	IP   string
	Port string
}

func (c *Client) ReadString(text *string, buf int) error {
	buffar := make([]byte, buf)
	n, err := c.Conn.Read(buffar)
	if err != nil {
		return err
	}
	res := string(buffar[:n])
	*text = res
	return nil
}

func (c *Client) WriteString(text string) error {
	_, err := io.WriteString(c.Conn, text)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ReadBytes(bytes *[]byte, buf int) error {
	buffar := make([]byte, buf)
	n, err := c.Conn.Read(buffar)
	if err != nil {
		return err
	}
	res := buffar[:n]
	*bytes = res
	return nil
}

func (c *Client) WriteBytes(bytes []byte) error {
	_, err := c.Conn.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
