package openapi

import (
	"context"

	"git.funxdata.com/brisk/grpc/rpc"
	"github.com/funxdata/openapi/geoip"
)

type Client struct {
	Key    string
	Secret string

	Context context.Context
}

func NewClient(key, secret string) *Client {
	return &Client{
		Key:     key,
		Secret:  secret,
		Context: context.Background(),
	}
}

func (c *Client) GeoIP() geoip.GeoIPClient {
	cc := rpc.Dial("openapi.funxdata.com")
	cli := geoip.NewGeoIPClient(cc)
	return cli
}
