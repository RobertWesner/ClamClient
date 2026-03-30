package client

import (
	"clamclient/packets"
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
	err    error
	ready  chan struct{}

	packetConn packets.PacketConn
	commands   *Commands
	events     *Events

	state    State
	username string
}

func (c *Client) fail(err error) {
	c.err = err
	c.cancel()
}

func (c *Client) Ready() <-chan struct{} {
	return c.ready
}

func (c *Client) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Client) Close() {
	c.cancel()
}

func (c *Client) Err() error {
	return c.err
}

func (c *Client) Events() *Events {
	return c.events
}

func Connect(ip string, username string) (c *Client) {
	c = &Client{
		username: username,
		state:    State{},
		ready:    make(chan struct{}),
		commands: NewCommands(),
		events:   NewEvents(),
	}

	conn, err := net.Dial("tcp", ip)
	if err != nil {
		c.err = fmt.Errorf("connect: %w", err)

		return
	}
	defer func() {
		<-c.ctx.Done()
		_ = conn.Close()
	}()

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		err = tcpConn.SetKeepAlive(true)
		if err != nil {
			slog.Warn("keep alive: %w", err)
		}
		err = tcpConn.SetKeepAlivePeriod(30 * time.Second)
		if err != nil {
			slog.Warn("keep alive period: %w", err)
		}
	}

	c.packetConn = packets.NewPacketConn(conn)

	keepAliveTicker := time.NewTicker(30 * time.Second)
	defer keepAliveTicker.Stop()

	c.ctx, c.cancel = context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case <-keepAliveTicker.C:
				// TODO: graceful error handling
				_ = c.packetConn.Write(packets.NewPacket0KeepAlive())
			}
		}
	}()

	go c.handlePackets()

	// TODO: move to own method :)
	err = c.packetConn.Write(packets.NewPacket1Login(14, username, 0, 0))
	if err != nil {
		c.err = fmt.Errorf("login: %w", err)

		return
	}

	return c
}
