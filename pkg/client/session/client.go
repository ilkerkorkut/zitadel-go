package session

import (
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel"
	session "github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/session/v2alpha"
)

type Client struct {
	Connection *zitadel.Connection
	session.SessionServiceClient
}

func NewClient(issuer, api string, scopes []string, options ...zitadel.Option) (*Client, error) {
	conn, err := zitadel.NewConnection(issuer, api, scopes, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Connection:           conn,
		SessionServiceClient: session.NewSessionServiceClient(conn.ClientConn),
	}, nil
}
