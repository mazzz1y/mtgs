package client

import (
	"context"
	"net"

	"mtgs/antireplay"
	"mtgs/config"
	"mtgs/mtproto"
	"mtgs/wrappers"
)

// Init defines common method for initializing client connections.
type Init func(context.Context, context.CancelFunc, net.Conn, string,
	antireplay.Cache, *config.Config, [][]byte) (wrappers.Wrap, *mtproto.ConnectionOpts, error)
